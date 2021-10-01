package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rqlite/gorqlite"
)

var SQLITE_FILE_URL string = "db.sqlite"
var SQLITE_FILE_MEM_URL string = "memdb.sqlite?cache=shared&mode=memory"
var SQLITE_MEM_URL string = ":memory:"

var RQLITE_URL string = "http://localhost:4001"

func CreateSQLiteTable(dbUrl string) {
	os.Remove(dbUrl)

	db, err := sql.Open("sqlite3", dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
	CREATE TABLE foo (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT);
	DELETE FROM foo;
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

func WriteSQLite(dbUrl string) {
	db, err := sql.Open("sqlite3", dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := "INSERT INTO foo(name) VALUES('Minh')"
	_, err = db.Exec(sqlStmt)

	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

func ReadSQLite(dbUrl string) {
	db, err := sql.Open("sqlite3", dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := "SELECT * FROM foo WHERE id=1"
	rows, err := db.Query(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
	defer rows.Close()
}

func CreateRQLIte(dbUrl string) {
	conn, err := gorqlite.Open(dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	conn.SetConsistencyLevel("strong")
	sqlStmt := `
	CREATE TABLE foo (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT);
	DELETE FROM foo;
	`
	conn.WriteOne(sqlStmt)
}

func WriteRQLite(dbUrl string) {
	conn, err := gorqlite.Open(dbUrl)
	if err != nil {
		log.Fatal(err)
	}

	conn.SetConsistencyLevel("strong")
	conn.WriteOne("INSERT INTO foo (name) values ('bar')")
}

func ReadRQLite(dbUrl string) {
	conn, err := gorqlite.Open(dbUrl)
	if err != nil {
		log.Fatal(err)
	}

	conn.SetConsistencyLevel("strong")
	conn.QueryOne("SELECT name FROM foo WHERE id=1")
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Println("Create database: " + SQLITE_FILE_URL)
	CreateSQLiteTable(SQLITE_FILE_URL)
	log.Println("Create database: " + SQLITE_FILE_MEM_URL)
	CreateSQLiteTable(SQLITE_FILE_MEM_URL)
	log.Println("Create database: " + RQLITE_URL)
	CreateRQLIte(RQLITE_URL)
}
