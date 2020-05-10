package mysql

import (
	"cqrses/es"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/spf13/viper"
)

var (
	config = map[string]string{
		"host":     "",
		"password": "",
		"database": "",
		"port":     "",
		"username": "",
	}
)

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	viper.AddConfigPath(home)
	viper.SetConfigType("yaml")
	viper.SetConfigName("cqrs")
	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
}
func newMysqlConn() (*sqlx.DB, error) {
	connString := config["username"] +
		":" + config["password"] +
		"@" + "(" + config["host"] +
		":" + config["port"] +
		")/" + config["database"]
	return sqlx.Connect("mysql", connString)
}
func CreateDb() error {
	conn, err := newMysqlConn()

	if err != nil {
		return err
	}

	tx, err := conn.Begin()

	if err != nil {
		return err
	}

	_, err = tx.Exec(aggregateSchema)

	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = tx.Exec(eventSchema)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

type testAgg struct {
	*es.AggregateBase
	Value int
}

func (a *testAgg) Apply(ev es.Event) error {
	testEv, ok := ev.(*TestEv)
	if !ok {
		return errors.New("Event apply handler not found")
	}
	a.Value = a.Value + testEv.Add
	return nil
}

type TestEv struct {
	Add int `json:"add" db:"add"`
}

func (t *TestEv) Type() string {
	return "TestEv"
}

func (t *TestEv) Scan(val interface{}) error {
	switch v := val.(type) {
	case []byte:
		return json.Unmarshal(v, t)
	case string:
		return json.Unmarshal([]byte(v), t)
	default:
		return errors.New(fmt.Sprintf("Unsupported type: %T", val))
	}
}

func (t *TestEv) Value() (driver.Value, error) {
	return json.Marshal(t)
}

func createTestAgg(id string) *testAgg {
	agg := testAgg{AggregateBase: es.NewAggregateBase(id)}
	return &agg
}

func setupTestRepo() *MysqlDomainRepo {
	CreateDb()
	conn, err := newMysqlConn()
	if err != nil {
		panic(err)
	}

	aggFactory := es.NewDelegateAggregateFactory()
	aggFactory.RegisterDelegate("testAgg", func(id string) es.AggregateRoot {
		return createTestAgg(id)
	})

	evFactory := es.NewDelegateEventFactory()
	evFactory.RegisterDelegate(&TestEv{})

	repo := MysqlDomainRepo{aggFactory, evFactory, conn}
	return &repo
}

func TestSaveAggregate(t *testing.T) {
	Convey("Given a repo", t, func() {
		repo := setupTestRepo()
		Convey("When create a cqrs aggregate", func() {
			agg := createTestAgg("someId")
			es.AddEvent(agg, &TestEv{1})
			es.AddEvent(agg, &TestEv{2})
			Convey("When save this agg via repo", func() {
				err := repo.Save(agg)
				defer repo.delete(agg.GetID())
				Convey("There should be no error", func() {
					So(err, ShouldBeNil)
				})
				Convey("When loading this aggregate back", func() {
					savedAgg, err := repo.Load("testAgg", "someId")
					Convey("There should be no error", func() {
						So(err, ShouldBeNil)
						Convey("Aggregate should not be nil", func() {
							So(savedAgg, ShouldNotBeNil)
							Convey("Aggregate should has same type as before", func() {
								testAgg, ok := savedAgg.(*testAgg)
								So(ok, ShouldEqual, true)

								Convey("Aggregate should have value of 3", func() {
									So(testAgg.Value, ShouldEqual, 3)
								})
							})
							Convey("When delete this aggregate", func() {
								err := repo.delete("someId")
								Convey("There should be no error", func() {
									So(err, ShouldBeNil)
								})
							})

						})

					})

					// Convey("Aggregate load from history should have id of someId", func() {
					// 	expectedId := "someId"
					// 	So(savedAgg.GetID(), ShouldEqual, expectedId)
					// })
				})
			})
		})
	})
}
