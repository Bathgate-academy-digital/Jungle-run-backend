package endpoints

import (
	"bufio"
	"jungle-rush/backend/database"
	ReturnModule "jungle-rush/backend/modules/return_module"
	"log"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ReturnModule.MethodNotAllowed(w)
		return
	}

	r.ParseForm()
	name := r.Form.Get("name")
	class := r.Form.Get("class")
	if !validParams(w, r, name, class) {
		return
	}

	// Generate random ID up to limit of PostgreSQL int https://www.postgresql.org/docs/15/datatype-numeric.html
	randomId := rand.Intn(2147483647)
	err := database.SubmitResult(randomId, name, class)
	if err != nil {
		ReturnModule.InternalServerError(w, "Failed to create user")
		return
	}
	ReturnModule.ID(w, randomId)
}

func validParams(w http.ResponseWriter, r *http.Request, name string, class string) bool {
	if name == "" || class == "" {
		ReturnModule.BadRequest(w, "Name and class cannot be empty")
		return false
	}
	if containsBadWords(name) {
		ReturnModule.BadRequest(w, "Name contains inappropriate words")
		return false
	}
	if len(class) != 2 {
		ReturnModule.BadRequest(w, "Class must be 2 characters long")
		return false
	}
	if !isValidClass(class) {
		ReturnModule.BadRequest(w, "Invalid class")
		return false
	}
	return true
}

func isValidClass(class string) bool {
	file, err := os.Open("data/classes.txt")
	if err != nil {
		log.Fatalln("Error opening classes.txt file: ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		validClass := scanner.Text()
		if class == validClass {
			return true
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln("Error reading classes.txt file: ", err)
	}

	return false
}

func containsBadWords(name string) bool {
	file, err := os.Open("data/bad_words.txt")
	if err != nil {
		log.Fatalln("Error opening bad_words.txt file: ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, ",")
		for _, badWord := range words {
			badWord = strings.TrimSpace(badWord)

			// Use case insensitive matching
			re := regexp.MustCompile(`(?i)\b` + regexp.QuoteMeta(badWord) + `\b`)
			if re.MatchString(name) {
				return true
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln("Error reading bad_words.txt file: ", err)
	}

	return false
}
