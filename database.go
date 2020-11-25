package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

// Create the database from string path

func createDb(db string) {
	database, _ :=
		sql.Open("sqlite3", (db))

	defer database.Close()

	statement, _ :=
		database.Prepare(
			"CREATE TABLE IF NOT EXISTS notes (id INTEGER PRIMARY KEY, content TEXT)")
	statement.Exec()

	statement, _ =
		database.Prepare(
			"CREATE TABLE IF NOT EXISTS events (id INTEGER PRIMARY KEY, title TEXT, long_desc TEXT, date TEXT)")
	statement.Exec()
}

//
//      SECTION:
//      Functions for handling Notes:
//

// Create a new note, two strings, the note content and the database path (constant in cally.go)

func createNote(text string, db string) {
	database, _ :=
		sql.Open("sqlite3", (db))

	defer database.Close()

	statement, _ :=
		database.Prepare(
			"INSERT INTO notes (content) VALUES (?)")
	statement.Exec(text)
}

// Delete a note based on id

func deleteNote(delId int, db string) {
	database, _ :=
		sql.Open("sqlite3", (db))

	defer database.Close()

	statement, _ :=
		database.Prepare(
			"DELETE FROM notes WHERE id = ?")
	statement.Exec(delId)
}

// Print all the notes in the databse

func readNotes(db string) {
	database, _ :=
		sql.Open("sqlite3", (db))

	defer database.Close()

	rows, _ :=
		database.Query(
			"SELECT id, content FROM notes")

	var id int
	var content string

	for rows.Next() {
		rows.Scan(&id, &content)
		fmt.Println(
			"\t" + strconv.Itoa(id) + ": " + content)
	}
}

//
//      SECTION:
//      Functions for dealing with events
//

// Create a new event

func createEvent(tit string, longDesc string, dt string, db string) {
    database, _ :=
        sql.Open("sqlite3", (db))
        
        defer database.Close()
        
        statement, _ :=
            database.Prepare(
                "INSERT INTO events (title, long_desc, date) VALUES (?,?,?)")
         
        statement.Exec(tit, longDesc, dt)
        fmt.Println("Created event:", tit)
}

// Delete an event, takes title field

func deleteEvent(tit string, db string) {
    database, _ :=
        sql.Open("sqlite3", (db))
        
    defer database.Close()
    
    statement, _ :=
        database.Prepare(
            "DELETE FROM events WHERE title = ?")
            
    statement.Exec(tit)
    fmt.Println("Deleted event:", tit)
}

// Print out all events in database

func printEvents(db string) {
	database, _ :=
		sql.Open("sqlite3", (db))

		defer database.Close()

		rows, _ :=
			database.Query(
				"SELECT title, long_desc, date FROM events")

		var (
			title,
			longDescription,
			date string)

		for rows.Next() {
			rows.Scan(&title, &longDescription, &date)
			fmt.Println(
				"\t", date, ":", title, "\n\t", longDescription, "\n")
		}
}
