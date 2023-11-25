# Jungle run backend

## Usage
To start the server run
```
go run server.go
```
you can now open `http://localhost:8080/api/leaderboard` in your browser (it should just return an empty `[]` for now). Then, to add new users with scores run:
```
curl --data "name=Alex&class=6N1&score=5" http://localhost:8080/api/submit
```
Try reloading the leaderboard page and you should see that this user is now on the leaderboard.

## Notes
- For now, the data is stored in memory and not in a database.
- The leaderboard is not sorted, but when the database is implemented it will be.
- Currently, this is designed so that at the end of the game it submits the score once, but it may be changed it the future to have a create and update endpoint so that the scores are updated in real time.
