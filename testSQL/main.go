package main

import (
	"database/sql"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	var n_tables int

	println(sql.Drivers())

	// URL connection string formats
	//    sqlserver://sa:mypass@localhost?database=master&connection+timeout=30         // username=sa, password=mypass.
	//    sqlserver://sa:my%7Bpass@somehost?connection+timeout=30                       // password is "my{pass"
	// note: pwd is "myP@55w0rd"
	connectString := "sqlserver://sa:P%40ssw0rd@172.16.20.12:1433?database=JEC_CORP&connection+timeout=30&encrypt=disable"
	println("Connection string=", connectString)

	println("open connection")
	db, err := sql.Open("mssql", connectString)
	defer db.Close()
	println("Open Error:", err)
	if err != nil {
		log.Fatal(err)
	}

	println("count records in TS_TABLES & scan")
	err = db.QueryRow("Select count(*) from appuser").Scan(&n_tables)
	if err != nil {
		log.Fatal(err)
	}
	println("count of users", n_tables)

	println("closing connection")
	db.Close()
}
