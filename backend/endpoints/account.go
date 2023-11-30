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
	"strconv"
	"strings"
)

func SubmitResult(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ReturnModule.MethodNotAllowed(w, r)
		return
	}

	r.ParseForm()
	name := r.Form.Get("name")
	class := r.Form.Get("class")

	if name == "" || class == "" {
		ReturnModule.BadRequest(w, r, "Name and class cannot be empty")
		return
	}

	if containsBadWords(name) {
		ReturnModule.BadRequest(w, r, "Name contains inappropriate words")
		return
	}

	score, err := strconv.Atoi(r.Form.Get("score"))
	if err != nil {
		ReturnModule.BadRequest(w, r, "Score must be an int")
		return
	}

	// Generate random ID up to limit of PostgreSQL int https://www.postgresql.org/docs/15/datatype-numeric.html
	randomId := rand.Intn(2147483647)
	err = database.SubmitResult(randomId, name, class, score)
	if err != nil {
		ReturnModule.InternalServerError(w, r, "Failed to create user")
		return
	}
	ReturnModule.ID(w, r, randomId)
}

func containsBadWords(name string) bool {
	file, err := os.Open("bad_words.txt")
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
