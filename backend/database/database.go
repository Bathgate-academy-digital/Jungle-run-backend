package database

import (
	"database/sql"
	"jungle-rush/backend/structs"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitializeDatabase() bool {
	connectString := "host=localhost port=5432 user=postgres password=<password> dbname=leaderboard sslmode=disable"
	openDatabase, err := sql.Open("postgres", connectString)
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
		);
	`)

	if err != nil {
		log.Fatal(err)
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
	query := "INSERT INTO users(id, name, class, score) VALUES ($1, $2, $3, $4)"
	_, err := db.Exec(query, id, name, class, score)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateUser(id int, score int) error {
	query := "UPDATE users SET score=$1 WHERE id=$2"
	_, err := db.Exec(query, score, id)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
