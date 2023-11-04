package database

import (
	"ba-digital/backend/structs"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var databaseConnection *sql.DB

func InitializeDatabase() bool {
	connection, databaseConnectionError := sql.Open("sqlite3", "database.db")

	if databaseConnectionError != nil {
		return false
	} else {
		databaseConnection = connection
		return true
	}
}

func CreateTables() bool {
	_, err := databaseConnection.Exec(`
		CREATE TABLE IF NOT EXISTS accounts (
			username VARCHAR(40) PRIMARY KEY,
			rank INT,
			password_hash VARCHAR(255),
			password_salt VARCHAR(25)
		);
			

		CREATE TABLE IF NOT EXISTS sessions (
			username VARCHAR(40),
			session_token VARCHAR(255),
			FOREIGN KEY (username) REFERENCES accounts(username)
		);

	`)

	if err != nil {
		return false
	}

	return true
}

func GetAccountDataFromSession(sessionToken string) structs.UserResponse {
	var userData structs.UserResponse

	query := "SELECT username FROM sessions WHERE session_token = ?"
	row := databaseConnection.QueryRow(query, sessionToken)

	var username sql.NullString
	if err := row.Scan(&username); err != nil {
		if err == sql.ErrNoRows {
			return userData
		}
		log.Println("Error while retrieving username:", err)
		return userData
	}

	if !username.Valid {
		return userData
	}

	row = databaseConnection.QueryRow("SELECT username, ranking FROM accounts WHERE username = ?", username)

	err := row.Scan(
		&userData.Username,
		&userData.Ranking,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return userData
		}
		log.Println("Error while scanning data:", err)
		return userData
	}

	row = databaseConnection.QueryRow("SELECT profile_picture, description, location, skills, interests, spoken_languages FROM profile_data WHERE username = ?", username)

	if err != nil {
		if err == sql.ErrNoRows {
			return userData
		}
		log.Println("Error while scanning data:", err)
		return userData
	}

	return userData
}
