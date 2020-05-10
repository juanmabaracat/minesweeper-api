# minesweeper-api
Minesweeper API implementation to create and save games.

## Notes
This API is designed to provide a client with the possibility of logging in users, creating new games, saving them 
and being able to obtain a list of the player's games.

### Technical details:
- [Hexagonal Architecture](https://alistair.cockburn.us/hexagonal-architecture/): Architectural pattern
- [gin-gonic](https://github.com/gin-gonic/gin): HTTP web framework
- [Go Modules](https://blog.golang.org/using-go-modules): dependency management
- [Testify](https://github.com/stretchr/testify): testing tool
- [Mockery](https://github.com/vektra/mockery): A mock code autogenerator for golang
- In memory DB
- [Docker](https://www.docker.com/)

### I would have liked to add:
- frontend using react library
- [jwt-go](https://github.com/dgrijalva/jwt-go): Go implementation of JSON Web Tokens (JWT)
- CassandraDB
- [gocql](https://github.com/gocql/gocql): Cassandra client for Go
- Deploy in AWS

### Endpoints

| Endpoint                         | Method     | Description
| -------------------------        | ---------- | -----------------------------------      |
| `/ping`                          | GET        | To the check if server is running        |
| `/minesweeper/api/game`          | GET        | Returns the id of all the player's games |
| `/minesweeper/api/game/:game_id` | GET        | Return a game for an specific id         |
| `/minesweeper/api/game`          | POST       | Create a new game                        |
| `/minesweeper/api/game/:game_id` | PUT        | Save the game                            |
| `/minesweeper/api/game/:game_id` | PATCH      | Modify game setting                      |
| `/minesweeper/api/user/login`    | POST       | User login. Returns a JWT token          |

## How to run the application:
clone the repository and move to the project folder:
```
$git clone https://github.com/juanmabaracat/minesweeper-api.git
$cd minesweeper-api
```
Use docker to build and run the application:
```
$docker build -t minesweeper-api .
$docker run -p 8080:8080 minesweeper-api:latest
```

## Examples
### Create a new game
### POST `/minesweeper/api/game`

| Code  | Description   |
| ----  | ------------- |
| 201   | Game created  |
| 400   | Bad request   |
| 500   | Server error  |

```
curl -i -X POST '127.0.0.1:8080/minesweeper/api/game' -d '{"player_id": 123, "width": 2, "height": 2, "mines": 1}
```

#### Request body example
```
{
	"player_id": 123,
	"width": 2,
	"height": 2,
	"mines": 1
}
```

#### Response
```
{
    "id": 15777586166085368371,
    "player_id": 123,
    "status": "playing",
    "board": {
        "width": 2,
        "height": 2,
        "mines": 1,
        "cells": [
            {
                "X": 0,
                "Y": 0,
                "has_mine": true,
                "was_revealed": false,
                "mark": "none"
            },
            {
                "X": 0,
                "Y": 1,
                "has_mine": false,
                "was_revealed": false,
                "mark": "none"
            },
            {
                "X": 1,
                "Y": 0,
                "has_mine": false,
                "was_revealed": false,
                "mark": "none"
            },
            {
                "X": 1,
                "Y": 1,
                "has_mine": false,
                "was_revealed": false,
                "mark": "none"
            }
        ]
    },
    "date_created": "2020-05-08T17:51:33Z",
    "date_finished": ""
}
```

### Error example, creating a game with invalid board size:
### POST `/minesweeper/api/game`
```
curl -i -X POST '127.0.0.1:8080/minesweeper/api/game' -d '{"player_id": 123, "width": 0, "height": 2, "mines": 1}
```

#### Request body example
```
{
	"player_id": 123,
	"width": 0,
	"height": 2,
	"mines": 1
}
```

#### Response
```
{
    "message": "invalid Board size",
    "status": 400,
    "error": "bad_request"
}
```

### Save game
### PUT `/minesweeper/api/game/:game_id`

| Code  | Description               |
| ----  | ------------------------  |
| 200   | Returns the game updated  |
| 400   | Bad request               |
| 404   | Game not found            |
| 500   | Server error              |
```
curl -i -X PUT '127.0.0.1:8080/minesweeper/api/game/15777586166085368371' -d {"status": "playing", "board": {...} }
```

#### Request body example
```
{
    "status": "playing",
    "board": {
        "width": 2,
        "height": 2,
        "mines": 1,
        "cells": [
            {
                "X": 0,
                "Y": 0,
                "has_mine": true,
                "was_revealed": false,
                "mark": "none"
            },
            {
                "X": 0,
                "Y": 1,
                "has_mine": false,
                "was_revealed": true,
                "mark": "none"
            },
            {
                "X": 1,
                "Y": 0,
                "has_mine": false,
                "was_revealed": true,
                "mark": "none"
            },
            {
                "X": 1,
                "Y": 1,
                "has_mine": false,
                "was_revealed": false,
                "mark": "none"
            }
        ]
    }
}
```

### Get a game
### GET `/minesweeper/api/game/:game_id`

| Code  | Description       |
| ----  | -------------     |
| 200   | returns the game  |
| 404   | Game not found    |
| 500   | Server error      |

```
curl -i -X GET '127.0.0.1:8080/minesweeper/api/game/15777586166085368371'
```

#### Response
```
{
    "id": 15777586166085368371,
    "player_id": 123,
    "status": "playing",
    "board": {
        "width": 2,
        "height": 2,
        "mines": 1,
        "cells": [
            {
                "X": 0,
                "Y": 0,
                "has_mine": true,
                "was_revealed": false,
                "mark": "none"
            },
            {
                "X": 0,
                "Y": 1,
                "has_mine": false,
                "was_revealed": true,
                "mark": "none"
            },
            {
                "X": 1,
                "Y": 0,
                "has_mine": false,
                "was_revealed": true,
                "mark": "none"
            },
            {
                "X": 1,
                "Y": 1,
                "has_mine": false,
                "was_revealed": false,
                "mark": "none"
            }
        ]
    },
    "date_created": "2020-05-08T18:14:12Z",
    "date_finished": ""
}
```

### Change game status
### PATCH `/minesweeper/api/game/:game_id`

| Code  | Description               |
| ----  | ------------------------  |
| 200   | Returns the game updated  |
| 400   | Bad request               |
| 404   | Game not found            |
| 500   | Server error              |

```
curl -i -X PATCH '127.0.0.1:8080/minesweeper/api/game/15777586166085368371' -d '{"status":"won"}
```

#### Request body example
```
{
	"status": "won"
}
```

#### Response
```
{
    "id": 15777586166085368371,
    "player_id": 123,
    "status": "won",
    "board": {
        ...
    },
    "date_created": "2020-05-08T18:13:12Z",
    "date_finished": "2020-05-08T18:14:24Z"
}
```