package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/SAP/go-hdb/driver" // Import the driver (side effect registers it)
)

func main1() {
	const numRow = 1000 // Number of rows to be inserted into table.

	// Replace with your actual SAP HANA connection details
	dsn := "hdb://AAABBB:Mon30daq!!@localhost:30215"

	db, err := sql.Open("hdb", dsn)
	if err != nil {
		log.Fatal("Error opening database:", err)
	}
	defer db.Close()

	tableName := driver.RandomIdentifier("table_")
	fmt.Println(tableName)
	// Create table.
	if _, err := db.Exec(fmt.Sprintf("create table %s (i integer, f double)", tableName)); err != nil {
		log.Fatal(err)
	}

	// Prepare statement.
	stmt, err := db.PrepareContext(context.Background(), fmt.Sprintf("insert into %s values (?, ?)", tableName))
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// Bulk insert via 'extended' argument list.
	args := make([]any, numRow*2)
	for i := 0; i < numRow; i++ { // Corrected the loop range
		args[i*2], args[i*2+1] = i, float64(i)
	}
	if _, err := stmt.Exec(args...); err != nil {
		log.Fatal(err)
	}

	// Bulk insert via function.
	i := 0
	if _, err := stmt.Exec(func(args []any) error {
		if i >= numRow {
			return driver.ErrEndOfRows
		}
		args[0], args[1] = i, float64(i)
		i++
		return nil
	}); err != nil {
		log.Fatal(err)
	}

	// Select number of inserted rows.
	var count int
	if err := db.QueryRow(fmt.Sprintf("select count(*) from %s", tableName)).Scan(&count); err != nil {
		log.Fatal(err)
	}
	fmt.Print(count)

	// Drop table.
	// if _, err := db.Exec(fmt.Sprintf("drop table %s", tableName)); err != nil {
	// 	log.Fatal(err)
	// }
}
