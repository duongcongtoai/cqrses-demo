package mysql

import (
	"database/sql"
	"encoding/json"
	"es"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type MysqlDomainRepo struct {
	aggregateFactory es.AggregateFactory
	eventFactory     es.EventFactory
	db               *sqlx.DB
}

type EventContainer struct {
	AggID     string    `json:"agg_id" db:"agg_id"`
	EventType string    `json:"type" db:"type"`
	Data      []byte    `json:"data" db:"data"`
	Version   int       `json:"version" db:"version"`
	CreatedOn time.Time `json:"created_on" db:"created_on"`
}

func (r *MysqlDomainRepo) Load(aggType, id string) (es.AggregateRoot, error) {
	if r.aggregateFactory == nil {
		return nil, fmt.Errorf("The domain repository has no aggregate factory")
	}

	aggregate := r.aggregateFactory.GetAggregate(aggType, id)
	if aggregate == nil {
		return nil, fmt.Errorf("The aggregate typed %s has not been registered", aggType)
	}
	schema := "SELECT  data, type FROM events WHERE agg_id = ? ORDER BY version ASC"
	rows, err := r.db.Queryx(schema, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var c = struct {
			Type string
			Data []byte
		}{}
		// var c es.EventMessage
		err := rows.StructScan(&c)
		if err != nil {
			return nil, err
		}
		event := r.eventFactory.GetEvent(string(c.Type))
		err = json.Unmarshal(c.Data, &event)
		if err != nil {
			return nil, err
		}
		err = aggregate.Apply(event)
		if err != nil {
			return nil, err
		}

		// eventMsg := es.NewEventMessage(c.AggID, event, c.Version, time.Now())
		// aggregate.Hydrate(event)
		// aggregate.IncrementVersion()
	}
	rows.Close()
	return aggregate, nil

}

func (r *MysqlDomainRepo) Save(agg es.AggregateRoot) error {
	expectedVersion := agg.PersistedVersion()
	resultEvents := agg.GetChanges()
	insertSchema := "INSERT INTO events(agg_id, data, type, version, created_on) VALUE (:agg_id, :data, :type, :version, :created_on)"
	if len(resultEvents) > 0 {
		tx, err := r.db.Begin()
		if err != nil {
			tx.Rollback()
			return err
		}
		queryLatestVersion := "SELECT version FROM aggregates WHERE agg_id = ?"

		row := tx.QueryRow(queryLatestVersion, agg.ID())
		var currVer int
		err = row.Scan(&currVer)
		if err != nil {
			switch err {
			case sql.ErrNoRows:
				currVer = 0
				_, err = tx.Exec("INSERT INTO aggregates(agg_id, version) VALUE (?, ?)", agg.ID(), 0)
				if err != nil {
					tx.Rollback()
					return err
				}
			default:
				tx.Rollback()
				return err
			}
		}
		if currVer != expectedVersion {
			tx.Rollback()
			return es.ErrConcurrency{
				Expected: expectedVersion,
				Given:    currVer,
			}
		}
		updateVersion := "UPDATE aggregates SET version = ? WHERE agg_id = ?"
		_, err = tx.Exec(updateVersion, expectedVersion+len(resultEvents), agg.ID())
		if err != nil {
			tx.Rollback()
			return err
		}

		for _, event := range resultEvents {
			// savedData, err := json.Marshal(v.Content())
			// if err != nil {
			// 	tx.Rollback()
			// 	return err
			// }
			// container := EventContainer{
			// 	AggID:     v.AggregateID(),
			// 	EventType: v.Type(),
			// 	Data:      savedData,
			// 	Version:   *expectedVersion + k + 1,
			// 	CreatedOn: time.Now(),
			// }
			query, args, err := sqlx.Named(insertSchema, event)
			query, args, err = sqlx.In(query, args...)
			if err != nil {
				tx.Rollback()
				return err
			}
			query = r.db.Rebind(query)
			_, err = tx.Exec(query, args...)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
		tx.Commit()
	}
	return nil

}
