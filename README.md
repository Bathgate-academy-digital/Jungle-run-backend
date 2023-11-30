# Jungle run backend

## Usage
First, set up credentials for a PostgreSQL server
```
export DB_HOST="127.0.0.1" 
export DB_NAME="leaderdboard" 
export DB_USER="<username>"
export DB_PASS="<password>"
```

Then, to start the server run
```
go run server.go
```
you can now open `http://localhost:8080/api/leaderboard` in your browser (it should just return an empty `[]` for now). To add new users with scores run
```
curl --data "name=Colin&class=5N2&score=0" http://localhost:8080/api/submit
```
This will return an ID, keep this for later. Try reloading the leaderboard page and you should see that this user is now on the leaderboard!

You can now update the score by using the ID you just obtained (this will update only existing users - NOT create new users)
```
curl --data "id=123456&score=5" http://localhost:8080/api/update
```

## Security
We have various security measures in place some of which are:
- Class verification
- User verification
- Banned word filtration
- Passwordless
