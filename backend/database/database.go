package database

import (
	"database/sql"
	"fmt"
	"jungle-rush/backend/structs"
	"log"
	"os"
	timeModule "time"

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
			time TIME
		);
	`)

	if err != nil {
		log.Fatalln("Failed to create tables:", err)
	}
}

func GetLeaderboard() []structs.User {
	rows, err := db.Query("SELECT name, class, time FROM users WHERE time IS NOT NULL ORDER BY time ASC")
	if err != nil {
		log.Println("Error fetching leaderboard from database:", err)
		return nil
	}
	defer rows.Close()

	output := []structs.User{}
	var (
		name  string
		class string
		time  timeModule.Time
	)
	for rows.Next() {
		err := rows.Scan(&name, &class, &time)
		if err != nil {
			log.Println("Error scanning leaderboard row:", err)
			return nil
		}
		user := structs.User{Name: name, Class: class, Time: time.Format("15:04:05")}
		output = append(output, user)
	}
	err = rows.Err()
	if err != nil {
		log.Println("Error processing leaderboard rows:", err)
		return nil
	}
	return output
}

func SubmitResult(id int, name string, class string) error {
	query := "INSERT INTO users(id, name, class) VALUES ($1, $2, $3)"
	_, err := db.Exec(query, id, name, class)
	if err != nil {
		log.Printf("Error inserting new user (name=%s class=%s): %s\n", name, class, err)
		return err
	}
	return nil
}

func UpdateUser(id int, time string) error {
	parsedTime, err := timeModule.Parse("15:04:05", time)
	if err != nil {
		log.Printf("Error parsing time (time=%s): %s\n", time, err)
		return err
	}

	query := "UPDATE users SET time=$1 WHERE id=$2"
	_, err = db.Exec(query, parsedTime, id)
	if err != nil {
		log.Printf("Error updating user (id=%d): %s\n", id, err)
		return err
	}
	return nil
}

func DeleteUser(id int) error {
	query := "DELETE FROM users WHERE id = $1"
	_, err := db.Exec(query, id)
	if err != nil {
		log.Printf("Error deleting user (id=%d): %s\n", id, err)
		return err
	}
	return nil
}

func GetUserId(name string, class string) int {
	query := "SELECT id FROM users WHERE name = $1 AND class = $2"
	var id int
	err := db.QueryRow(query, name, class).Scan(&id)
	if err != nil {
		log.Printf("Error getting user id (name=%s class=%s): %s\n", name, class, err)
		return -1
	}
	return id
}
