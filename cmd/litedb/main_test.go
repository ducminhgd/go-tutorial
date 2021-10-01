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

// func BenchmarkWriteSQLiteGoroutine(b *testing.B) {
// 	db, err := sql.Open("sqlite3", SQLITE_FILE_URL)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	sqlStmt := "INSERT INTO foo(name) VALUES('Minh')"
// 	for i := 0; i < b.N; i++ {
// 		go db.Exec(sqlStmt)
// 	}
// }

// func BenchmarkReadSQLiteGoroutine(b *testing.B) {
// 	db, err := sql.Open("sqlite3", SQLITE_FILE_URL)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	sqlStmt := "SELECT * FROM foo WHERE id=1"
// 	for i := 0; i < b.N; i++ {
// 		go db.Query(sqlStmt)
// 	}
// }

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

func BenchmarkWriteShareMemSQLiteGoroutine(b *testing.B) {
	db, err := sql.Open("sqlite3", SQLITE_FILE_MEM_URL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := "INSERT INTO foo(name) VALUES('Minh')"
	for i := 0; i < b.N; i++ {
		go db.Exec(sqlStmt)
	}
}

func BenchmarkReadShareMemSQLiteGoroutine(b *testing.B) {
	db, err := sql.Open("sqlite3", SQLITE_FILE_MEM_URL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := "SELECT * FROM foo WHERE id=1"
	for i := 0; i < b.N; i++ {
		go db.Query(sqlStmt)
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
	sqlStmt := "INSERT INTO foo(name) VALUES('Minh')"
	for i := 0; i < b.N; i++ {
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

// func BenchmarkWriteMemSQLiteGoroutine(b *testing.B) {
// 	db, err := sql.Open("sqlite3", SQLITE_MEM_URL)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	createTbStmt := `
// 	CREATE TABLE foo (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT);
// 	DELETE FROM foo;
// 	`
// 	_, err = db.Exec(createTbStmt)
// 	if err != nil {
// 		log.Printf("%q: %s\n", err, createTbStmt)
// 		return
// 	}
// 	sqlStmt := "INSERT INTO foo(name) VALUES('Minh')"
// 	for i := 0; i < b.N; i++ {
// 		go db.Exec(sqlStmt)
// 	}
// }

// // This benchmark is not perfect yet, since it has to create table before run test
// func BenchmarkReadMemSQLiteGoroutine(b *testing.B) {
// 	db, err := sql.Open("sqlite3", SQLITE_MEM_URL)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	sqlStmt := `
// 	CREATE TABLE foo (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT);
// 	DELETE FROM foo;
// 	INSERT INTO foo(name) VALUES('Minh');
// 	`
// 	_, err = db.Exec(sqlStmt)
// 	if err != nil {
// 		log.Printf("%q: %s\n", err, sqlStmt)
// 		return
// 	}
// 	query := "SELECT * FROM foo WHERE id=1"
// 	for i := 0; i < b.N; i++ {
// 		go db.Query(query)
// 	}

// }

func BenchmarkWriteRQLite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WriteRQLite(RQLITE_URL)
	}
}
func BenchmarkReadRQLite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReadRQLite(RQLITE_URL)
	}
}

// func BenchmarkWriteRQLiteGoroutine(b *testing.B) {
// 	conn, err := gorqlite.Open(RQLITE_URL)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	conn.SetConsistencyLevel("strong")
// 	for i := 0; i < b.N; i++ {
// 		go conn.WriteOne("INSERT INTO foo (name) values ('bar')")
// 	}
// }
// func BenchmarkReadRQLiteGoroutine(b *testing.B) {
// 	conn, err := gorqlite.Open(RQLITE_URL)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	conn.SetConsistencyLevel("strong")
// 	for i := 0; i < b.N; i++ {
// 		go conn.QueryOne("SELECT name FROM foo WHERE id=1")
// 	}
// }
