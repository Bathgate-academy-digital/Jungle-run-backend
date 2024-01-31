package admin_module

import (
	"log"
	"net/http"
)

func IsCorrectCreds(r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	if username == "admin" && password == "5fe88ee2442925b67c5aa328ae3c65445b66d195fef92e705360931d3c2f037b" && ok {
		return true
	}
	log.Printf("Wrong credentials: username=%v password=%v", username, password)
	return false
}
