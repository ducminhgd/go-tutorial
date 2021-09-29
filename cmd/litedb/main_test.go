package main

import (
	"database/sql"
	"log"
	"testing"
)

func BenchmarkWriteSQLite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WriteSQLite(SQLITE_FILE_URL)
	}
}

func BenchmarkReadSQLite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReadSQLite(SQLITE_FILE_URL)
	}
}

func BenchmarkWriteShareMemSQLite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WriteSQLite(SQLITE_FILE_MEM_URL)
	}
}

func BenchmarkReadShareMemSQLite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReadSQLite(SQLITE_FILE_MEM_URL)
	}
}

func BenchmarkWriteMemSQLite(b *testing.B) {
	db, err := sql.Open("sqlite3", SQLITE_MEM_URL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createTbStmt := `
	CREATE TABLE foo (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT);
	DELETE FROM foo;
	`
	_, err = db.Exec(createTbStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, createTbStmt)
		return
	}
	for i := 0; i < b.N; i++ {
		sqlStmt := "INSERT INTO foo(name) VALUES('Minh')"
		_, err = db.Exec(sqlStmt)

		if err != nil {
			log.Printf("%q: %s\n", err, sqlStmt)
			return
		}
	}
}

// This benchmark is not perfect yet, since it has to create table before run test
func BenchmarkReadMemSQLite(b *testing.B) {
	db, err := sql.Open("sqlite3", SQLITE_MEM_URL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	for i := 0; i < b.N; i++ {
		sqlStmt := `
		CREATE TABLE foo (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT);
		DELETE FROM foo;
		INSERT INTO foo(name) VALUES('Minh');
		`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			log.Printf("%q: %s\n", err, sqlStmt)
			return
		}

		query := "SELECT * FROM foo WHERE id=1"
		rows, err := db.Query(query)
		if err != nil {
			log.Printf("%q: %s\n", err, query)
			return
		}
		defer rows.Close()
	}
}
