package main

import (
	"ddd/db"
	"fmt"
)

var (
	schema = `CREATE TABLE events (
		agg_id varchar(36) NOT NULL,
		data BLOB NOT NULL,
		created_on DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		version INT NOT NULL,
		type varchar(20) NOT NULL,
		CONSTRAINT fk_agg_id
		FOREIGN KEY (agg_id) REFERENCES aggregates(agg_id)) ENGINE = InnodB`
	aggregateSchema = `CREATE TABLE aggregates (agg_id VARCHAR(36) NOT NULL ,
		 version INT NOT NULL, UNIQUE(agg_id) ) ENGINE = InnoDB`
	// testschema = `INSERT INTO events(agg_id, type, version, created_on, data) VALUES (:aggId, :type, :version, :created_on, :data)`

)

func main() {
	conn, err := db.NewMysqlConnection()
	// fmt.Println(err)
	// arg := map[string]interface{}{
	// 	"aggId":      "Something",
	// 	"type":       "Somethingelse hehehe",
	// 	"created_on": time.Now(),
	// 	"data":       []byte("Hello world"),
	// 	"version":    1,
	// }
	// query, args, err := sqlx.Named(testschema, arg)
	// query, args, err = sqlx.In(query, args...)
	// query = conn.Rebind(query)
	// tx, err := conn.Begin()
	// fmt.Println(err)
	// _, err = tx.Exec(query, args...)
	// fmt.Println(err)
	// err = tx.Commit()
	// fmt.Println(err)
	if err != nil {
		panic(err)
	}
	tx, err := conn.Begin()
	if err != nil {
		panic(err)
	}

	_, err = tx.Exec(aggregateSchema)

	if err != nil {
		fmt.Printf("Error : %v", err)
		return
	}
	_, err = tx.Exec(schema)
	if err != nil {
		panic(err)
	}
	err = tx.Commit()
	if err != nil {
		panic(err)
	}
	fmt.Println("Sucess!!!")
}
