package database

import (
	"database/sql"
	"fmt"
	"jungle-rush/backend/structs"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitializeDatabase() {
	var (
		dbUser = mustGetenv("DB_USER")
		dbPwd  = mustGetenv("DB_PASS")
		host   = mustGetenv("DB_HOST")
		dbName = mustGetenv("DB_NAME")
	)

	connectString := fmt.Sprintf("user=%s password=%s database=%s host=%s", dbUser, dbPwd, dbName, host)
	openDatabase, err := sql.Open("postgres", connectString)
	if err != nil {
		log.Fatalln("Failed to initialise database:", err)
	}
	db = openDatabase
}

func mustGetenv(envName string) string {
	v := os.Getenv(envName)
	if v == "" {
		log.Fatalf("%s environment variable not set\n", envName)
	}
	return v
}

func CreateTables() {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users(
			id INT PRIMARY KEY,
			name TEXT,
			class TEXT,
			score TIME
		);
	`)

	if err != nil {
		log.Fatalln("Failed to create tables:", err)
	}
}

func GetLeaderboard() []structs.User {
	rows, err := db.Query("SELECT name, class, score FROM users ORDER BY score DESC")
	if err != nil {
		log.Println("Error fetching leaderboard from database:", err)
		return nil
	}
	defer rows.Close()

	output := []structs.User{}
	var (
		name  string
		class string
		score time.Time
	)
	for rows.Next() {
		err := rows.Scan(&name, &class, &score)
		if err != nil {
			log.Println("Error scanning leaderboard row:", err)
			return nil
		}
		user := structs.User{Name: name, Class: class, Score: score.Format("15:04")}
		output = append(output, user)
	}
	err = rows.Err()
	if err != nil {
		log.Println("Error processing leaderboard rows:", err)
		return nil
	}
	return output
}

func SubmitResult(id int, name string, class string, score string) error {
	parsedScore, err := time.Parse("15:04", score)
	if err != nil {
		log.Printf("Error parsing score (score=%s): %s\n", score, err)
		return err
	}

	query := "INSERT INTO users(id, name, class, score) VALUES ($1, $2, $3, $4)"
	_, err = db.Exec(query, id, name, class, parsedScore)
	if err != nil {
		log.Printf("Error inserting new user (name=%s class=%s): %s\n", name, class, err)
		return err
	}
	return nil
}

func UpdateUser(id int, score string) error {
	parsedScore, err := time.Parse("15:04", score)
	if err != nil {
		log.Printf("Error parsing score (score=%s): %s\n", score, err)
		return err
	}

	query := "UPDATE users SET score=$1 WHERE id=$2"
	_, err = db.Exec(query, parsedScore, id)
	if err != nil {
		log.Printf("Error updating user (id=%d): %s\n", id, err)
		return err
	}
	return nil
}
