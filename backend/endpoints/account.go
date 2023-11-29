package endpoints

import (
	"bufio"
	"fmt"
	"jungle-rush/backend/database"
	ReturnModule "jungle-rush/backend/modules/return_module"
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
		ReturnModule.CustomError(w, r, "Bad Request: name and class cannot be empty", 400)
		return
	}

	// Check for bad words
	if containsBadWords(name) {
		ReturnModule.CustomError(w, r, "Bad Request: Name contains inappropriate words", 400)
		return
	}

	score, err := strconv.Atoi(r.Form.Get("score"))
	if err != nil {
		ReturnModule.CustomError(w, r, "Bad Request: score must be an int", 400)
		return
	}

	// Limit of PostgreSQL int https://www.postgresql.org/docs/15/datatype-numeric.html
	randomId := rand.Intn(2147483647)
	database.SubmitResult(randomId, name, class, score)
	ReturnModule.ID(w, r, randomId)
}

func containsBadWords(name string) bool {
	file, err := os.Open("bad_words.txt")
	if err != nil {
		fmt.Println("Error opening bad_words.txt file:", err)
		return false
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
		fmt.Println("Error reading bad_words.txt file:", err)
	}

	return false
}
