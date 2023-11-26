package database

import (
	"ba-digital/backend/structs"
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitializeDatabase() bool {
	openDatabase, err := sql.Open("sqlite3", "sqlite.db")
	if err != nil {
		return false
	}

	db = openDatabase
	return true
}

func CreateTables() bool {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users(
			id INT PRIMARY KEY,
			name TEXT,
			class TEXT,
			score INT
		) STRICT;
	`)

	if err != nil {
		return false
	}
	return true
}

func GetLeaderboard() []structs.User {
	rows, err := db.Query("SELECT name, class, score FROM users ORDER BY score DESC")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	output := []structs.User{}
	var (
		name  string
		class string
		score int
	)
	for rows.Next() {
		err := rows.Scan(&name, &class, &score)
		if err != nil {
			log.Fatal(err)
		}
		user := structs.User{Name: name, Class: class, Score: score}
		output = append(output, user)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return output
}

func SubmitResult(id int, name string, class string, score int) {
	query := "INSERT INTO users(id, name, class, score) VALUES (?, ?, ?, ?)"
	_, err := db.Exec(query, id, name, class, score)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateUser(name string, class string, score int) error {
	query := "UPDATE users SET class=?, score=? WHERE name=?"
	_, err := db.Exec(query, class, score, name)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
