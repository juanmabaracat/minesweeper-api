# minesweeper-api
Minesweeper API implementation to create and save games.

## Notes
This API is designed to provide a client with the possibility of logging in users, creating new games, saving them 
and being able to obtain a list of the player's games.
Given the time available for the challenge, some parts were left out.

### Technical details:
- Hexagonal Architecture
- gin-gonic: HTTP web framework
- Go Modules: dependencies management
- Testify: testing tool
- Mockery: A mock code autogenerator for golang
- In memory DB
- Docker

What was left:
- frontend using react library
- JWT authentication
- CassandraDB
- gocql: Cassandra client for Go
- Docker compose to bundle database and application
- Deploy in AWS
- Some details like limiting board size, more test coverage, some endpoints.

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
```
curl -i -X POST '127.0.0.1:8080/minesweeper/api/game' -d '{"player_id": 123, "width": 3, "height": 3, "mines": 2}
```

#### Request body example
```

{
	"player_id": 123,
	"width": 3,
	"height": 3,
	"mines": 2
}
```

#### Response
```
{
    "id": 979070430525095337,
    "player_id": 123,
    "status": "PLAYING",
    "Board": {
        "width": 3,
        "height": 3,
        "mines": 2,
        "cells": [
            {
                "X": 0,
                "Y": 0,
                "has_mine": false,
                "was_revealed": false,
                "Mark": "NONE"
            },
            {
                "X": 0,
                "Y": 1,
                "has_mine": false,
                "was_revealed": false,
                "Mark": "NONE"
            },
            {
                "X": 0,
                "Y": 2,
                "has_mine": false,
                "was_revealed": false,
                "Mark": "NONE"
            },
            {
                "X": 1,
                "Y": 0,
                "has_mine": false,
                "was_revealed": false,
                "Mark": "NONE"
            },
            {
                "X": 1,
                "Y": 1,
                "has_mine": false,
                "was_revealed": false,
                "Mark": "NONE"
            },
            {
                "X": 1,
                "Y": 2,
                "has_mine": false,
                "was_revealed": false,
                "Mark": "NONE"
            },
            {
                "X": 2,
                "Y": 0,
                "has_mine": true,
                "was_revealed": false,
                "Mark": "NONE"
            },
            {
                "X": 2,
                "Y": 1,
                "has_mine": false,
                "was_revealed": false,
                "Mark": "NONE"
            },
            {
                "X": 2,
                "Y": 2,
                "has_mine": true,
                "was_revealed": false,
                "Mark": "NONE"
            }
        ]
    },
    "date_created": "2020-05-08T16:12:10Z",
    "date_finished": ""
}
```
