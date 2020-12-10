**_Golang minesweeper game - rest api_**

# minesweeper-API
minesweeper-API-server

####Description: https://github.com/deviget/minesweeper-API

**Core tasks:**
    Design and implement a documented RESTful API for the game (think of a mobile app for your API)
    Implement an API client library for the API designed above. Ideally, in a different language, of your preference, to the one used for the API
    When a cell with no adjacent mines is revealed, all adjacent squares will be revealed (and repeat)
    Ability to 'flag' a cell with a question mark or red flag
    Detect when game is over
    Persistence
    Time tracking
    Ability to start a new game and preserve/resume the old ones
    Ability to select the game parameters: number of rows, columns, and mines
    Ability to support multiple users/accounts
### Usage
####Endpoints Example
#####- Create game 3x3x1 transforms => 8x8x10

`	
			"request": {
				"method": "POST",
				"header": [],
				"url": {
					"raw": "localhost:8080/game/new?rows=3&columns=3&mines=1&name=myTestGame",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"game",
						"new"
					],
					"query": [
						{
							"key": "rows",
							"value": "3"
						},
						{
							"key": "columns",
							"value": "3"
						},
						{
							"key": "mines",
							"value": "1"
						},
						{
							"key": "name",
							"value": "myTestGame"
						}
					]
				}
			},
			"response":{
                           "Status": 200,
                           "Message": {
                               "name": "myTestGame2",
                               "uuid": "cb5e4288-8038-4fbd-9fce-933a7b770430",
                               "rows": 8,
                               "cols": 8,
                               "mines": 10,
                               "grid": [
                                   [
                                       {
                                           "column": 0,
                                           "row": 0,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": false,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       },
                                       {
                                           "column": 0,
                                           "row": 1,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": false,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       },
                                       {
                                           "column": 0,
                                           "row": 2,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": false,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       },
                                       {
                                           "column": 0,
                                           "row": 3,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": false,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       },
                                       {
                                           "column": 0,
                                           "row": 4,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": false,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       },
                                       {
                                           "column": 0,
                                           "row": 5,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": false,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       },
                                       {
                                           "column": 0,
                                           "row": 6,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": false,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       },
                                       {
                                           "column": 0,
                                           "row": 7,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": false,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       }
                                   ],
                                   [
                                       {
                                           "column": 1,
                                           "row": 0,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": false,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       },
                                       {
                                           "column": 1,
                                           "row": 1,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": false,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       },
                                       {
                                           "column": 1,
                                           "row": 2,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": false,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       },
                                       {
                                           "column": 1,
                                           "row": 3,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": false,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       },
                                       {
                                           "column": 1,
                                           "row": 4,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": false,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       },
                                       {
                                           "column": 1,
                                           "row": 5,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": false,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       },
                                       {
                                           "column": 1,
                                           "row": 6,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": true,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       },
                                       {
                                           "column": 1,
                                           "row": 7,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": false,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       }
                                   ],
                                   [
                                       {
                                           "column": 2,
                                           "row": 0,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": false,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       },
                                       {
                                           "column": 2,
                                           "row": 1,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": false,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       },
                                       {
                                           "column": 2,
                                           "row": 2,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": false,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       },
                                       {
                                           "column": 2,
                                           "row": 3,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": false,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       },
                                       {
                                           "column": 2,
                                           "row": 4,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": false,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       },
                                       {
                                           "column": 2,
                                           "row": 5,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": true,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       },
                                       {
                                           "column": 2,
                                           "row": 6,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": false,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       },
                                       {
                                           "column": 2,
                                           "row": 7,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": false,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       }
                                   ],
                                   [
                                       {
                                           "column": 3,
                                           "row": 0,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": false,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       },
                                       {
                                           "column": 3,
                                           "row": 1,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": true,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       },
                                       {
                                           "column": 3,
                                           "row": 2,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": false,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       },
                                       {
                                           "column": 3,
                                           "row": 3,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": true,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       },
                                       {
                                           "column": 3,
                                           "row": 4,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": true,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       },
                                       {
                                           "column": 3,
                                           "row": 5,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": false,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       },
                                       {
                                           "column": 3,
                                           "row": 6,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": false,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       },
                                       {
                                           "column": 3,
                                           "row": 7,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": false,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       }
                                   ],
                                   [
                                       {
                                           "column": 4,
                                           "row": 0,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": true,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       },
                                       {
                                           "column": 4,
                                           "row": 1,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": false,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       },
                                       {
                                           "column": 4,
                                           "row": 2,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": false,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       },
                                       {
                                           "column": 4,
                                           "row": 3,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": false,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       },
                                       {
                                           "column": 4,
                                           "row": 4,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": false,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       },
                                       {
                                           "column": 4,
                                           "row": 5,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": false,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       },
                                       {
                                           "column": 4,
                                           "row": 6,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": false,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       },
                                       {
                                           "column": 4,
                                           "row": 7,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": true,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       }
                                   ],
                                   [
                                       {
                                           "column": 5,
                                           "row": 0,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": false,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       },
                                       {
                                           "column": 5,
                                           "row": 1,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": false,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       },
                                       {
                                           "column": 5,
                                           "row": 2,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": false,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       },
                                       {
                                           "column": 5,
                                           "row": 3,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": false,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       },
                                       {
                                           "column": 5,
                                           "row": 4,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": false,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       },
                                       {
                                           "column": 5,
                                           "row": 5,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": false,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       },
                                       {
                                           "column": 5,
                                           "row": 6,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": true,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       },
                                       {
                                           "column": 5,
                                           "row": 7,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": false,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       }
                                   ],
                                   [
                                       {
                                           "column": 6,
                                           "row": 0,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": false,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       },
                                       {
                                           "column": 6,
                                           "row": 1,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": false,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       },
                                       {
                                           "column": 6,
                                           "row": 2,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": false,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       },
                                       {
                                           "column": 6,
                                           "row": 3,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": false,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       },
                                       {
                                           "column": 6,
                                           "row": 4,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": false,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       },
                                       {
                                           "column": 6,
                                           "row": 5,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": true,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       },
                                       {
                                           "column": 6,
                                           "row": 6,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": false,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       },
                                       {
                                           "column": 6,
                                           "row": 7,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": false,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       }
                                   ],
                                   [
                                       {
                                           "column": 7,
                                           "row": 0,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": true,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       },
                                       {
                                           "column": 7,
                                           "row": 1,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": false,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       },
                                       {
                                           "column": 7,
                                           "row": 2,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": false,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       },
                                       {
                                           "column": 7,
                                           "row": 3,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": false,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       },
                                       {
                                           "column": 7,
                                           "row": 4,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": false,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       },
                                       {
                                           "column": 7,
                                           "row": 5,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": false,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       },
                                       {
                                           "column": 7,
                                           "row": 6,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": false,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       },
                                       {
                                           "column": 7,
                                           "row": 7,
                                           "value": 0,
                                           "adjacent_bombs": 0,
                                           "mine": false,
                                           "clicked": false,
                                           "flagged": false,
                                           "opened": false,
                                           "covered": false
                                       }
                                   ]
                               ],
                               "clicks": 0,
                               "flags": 0,
                               "open_cells": 0,
                               "covered_cells": 0
                           }
                       }
		},`
		
##### Click position

			`"request": {
				"method": "PUT",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\r\n    \"column\": 0,\r\n    \"row\": 1,\r\n    \"flag\": false\r\n}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "localhost:8080/game/movement?name=myTestGame2&uuid=a57340e4-28bb-40c8-9897-3d0b2fe397a7",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"game",
						"movement"
					],
					"query": [
						{
							"key": "name",
							"value": "myTestGame"
						},
						{
							"key": "uuid",
							"value": "a57340e4-28bb-40c8-9897-3d0b2fe397a7"
						}
					]
				}
			},
			"response": [{
                             "Status": 200,
                             "Message": {
                                 "GameStatus": {
                                     "alive": true,
                                     "status": "Game in progress",
                                     "won": ""
                                 },
                                 "Cell": {
                                     "column": 1,
                                     "row": 0,
                                     "value": 0,
                                     "adjacent_bombs": 0,
                                     "mine": true,
                                     "clicked": false,
                                     "flagged": false,
                                     "opened": false,
                                     "covered": false
                                 },
                                 "Game": {
                                     "name": "myTestGame2",
                                     "uuid": "959de865-1304-4558-a310-a4a67ae92b2b",
                                     "rows": 8,
                                     "cols": 8,
                                     "mines": 10,
                                     "grid": [
                                         [
                                             {
                                                 "column": 0,
                                                 "row": 0,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 0,
                                                 "row": 1,
                                                 "value": 0,
                                                 "adjacent_bombs": 1,
                                                 "mine": false,
                                                 "clicked": true,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": true
                                             },
                                             {
                                                 "column": 0,
                                                 "row": 2,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 0,
                                                 "row": 3,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 0,
                                                 "row": 4,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 0,
                                                 "row": 5,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 0,
                                                 "row": 6,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 0,
                                                 "row": 7,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             }
                                         ],
                                         [
                                             {
                                                 "column": 1,
                                                 "row": 0,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": true,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 1,
                                                 "row": 1,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 1,
                                                 "row": 2,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 1,
                                                 "row": 3,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 1,
                                                 "row": 4,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 1,
                                                 "row": 5,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": true,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 1,
                                                 "row": 6,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 1,
                                                 "row": 7,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             }
                                         ],
                                         [
                                             {
                                                 "column": 2,
                                                 "row": 0,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 2,
                                                 "row": 1,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": true,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 2,
                                                 "row": 2,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 2,
                                                 "row": 3,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 2,
                                                 "row": 4,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 2,
                                                 "row": 5,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 2,
                                                 "row": 6,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 2,
                                                 "row": 7,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             }
                                         ],
                                         [
                                             {
                                                 "column": 3,
                                                 "row": 0,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 3,
                                                 "row": 1,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 3,
                                                 "row": 2,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": true,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 3,
                                                 "row": 3,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 3,
                                                 "row": 4,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": true,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 3,
                                                 "row": 5,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 3,
                                                 "row": 6,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 3,
                                                 "row": 7,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             }
                                         ],
                                         [
                                             {
                                                 "column": 4,
                                                 "row": 0,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 4,
                                                 "row": 1,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 4,
                                                 "row": 2,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 4,
                                                 "row": 3,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 4,
                                                 "row": 4,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": true,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 4,
                                                 "row": 5,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 4,
                                                 "row": 6,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 4,
                                                 "row": 7,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             }
                                         ],
                                         [
                                             {
                                                 "column": 5,
                                                 "row": 0,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 5,
                                                 "row": 1,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 5,
                                                 "row": 2,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 5,
                                                 "row": 3,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 5,
                                                 "row": 4,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 5,
                                                 "row": 5,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 5,
                                                 "row": 6,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 5,
                                                 "row": 7,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": true,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             }
                                         ],
                                         [
                                             {
                                                 "column": 6,
                                                 "row": 0,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 6,
                                                 "row": 1,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": true,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 6,
                                                 "row": 2,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": true,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 6,
                                                 "row": 3,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 6,
                                                 "row": 4,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 6,
                                                 "row": 5,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": true,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 6,
                                                 "row": 6,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 6,
                                                 "row": 7,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             }
                                         ],
                                         [
                                             {
                                                 "column": 7,
                                                 "row": 0,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 7,
                                                 "row": 1,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 7,
                                                 "row": 2,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 7,
                                                 "row": 3,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 7,
                                                 "row": 4,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 7,
                                                 "row": 5,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 7,
                                                 "row": 6,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 7,
                                                 "row": 7,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             }
                                         ]
                                     ],
                                     "clicks": 1,
                                     "flags": 0,
                                     "open_cells": 0,
                                     "covered_cells": 1
                                 }
                             }
                         },
                         {
                             "Status": 200,
                             "Message": {
                                 "GameStatus": {
                                     "alive": false,
                                     "status": "Game finished!",
                                     "won": "You lose!"
                                 },
                                 "Cell": {
                                     "column": 1,
                                     "row": 0,
                                     "value": 0,
                                     "adjacent_bombs": 0,
                                     "mine": false,
                                     "clicked": false,
                                     "flagged": false,
                                     "opened": false,
                                     "covered": false
                                 },
                                 "Game": {
                                     "name": "myTestGame2",
                                     "uuid": "fe02cd48-0962-40cb-b8e4-1300d00b6905",
                                     "rows": 8,
                                     "cols": 8,
                                     "mines": 10,
                                     "grid": [
                                         [
                                             {
                                                 "column": 0,
                                                 "row": 0,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 0,
                                                 "row": 1,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": true,
                                                 "clicked": true,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 0,
                                                 "row": 2,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 0,
                                                 "row": 3,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 0,
                                                 "row": 4,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 0,
                                                 "row": 5,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 0,
                                                 "row": 6,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": true,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 0,
                                                 "row": 7,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             }
                                         ],
                                         [
                                             {
                                                 "column": 1,
                                                 "row": 0,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 1,
                                                 "row": 1,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 1,
                                                 "row": 2,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": true,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 1,
                                                 "row": 3,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 1,
                                                 "row": 4,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": true,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 1,
                                                 "row": 5,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 1,
                                                 "row": 6,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 1,
                                                 "row": 7,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             }
                                         ],
                                         [
                                             {
                                                 "column": 2,
                                                 "row": 0,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 2,
                                                 "row": 1,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 2,
                                                 "row": 2,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 2,
                                                 "row": 3,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": true,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 2,
                                                 "row": 4,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 2,
                                                 "row": 5,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 2,
                                                 "row": 6,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 2,
                                                 "row": 7,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             }
                                         ],
                                         [
                                             {
                                                 "column": 3,
                                                 "row": 0,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": true,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 3,
                                                 "row": 1,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 3,
                                                 "row": 2,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 3,
                                                 "row": 3,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": true,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 3,
                                                 "row": 4,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 3,
                                                 "row": 5,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 3,
                                                 "row": 6,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 3,
                                                 "row": 7,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             }
                                         ],
                                         [
                                             {
                                                 "column": 4,
                                                 "row": 0,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 4,
                                                 "row": 1,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 4,
                                                 "row": 2,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 4,
                                                 "row": 3,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 4,
                                                 "row": 4,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 4,
                                                 "row": 5,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 4,
                                                 "row": 6,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 4,
                                                 "row": 7,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             }
                                         ],
                                         [
                                             {
                                                 "column": 5,
                                                 "row": 0,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 5,
                                                 "row": 1,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 5,
                                                 "row": 2,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 5,
                                                 "row": 3,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 5,
                                                 "row": 4,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": true,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 5,
                                                 "row": 5,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 5,
                                                 "row": 6,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 5,
                                                 "row": 7,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             }
                                         ],
                                         [
                                             {
                                                 "column": 6,
                                                 "row": 0,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 6,
                                                 "row": 1,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 6,
                                                 "row": 2,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 6,
                                                 "row": 3,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 6,
                                                 "row": 4,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": true,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 6,
                                                 "row": 5,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 6,
                                                 "row": 6,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 6,
                                                 "row": 7,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             }
                                         ],
                                         [
                                             {
                                                 "column": 7,
                                                 "row": 0,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 7,
                                                 "row": 1,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 7,
                                                 "row": 2,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 7,
                                                 "row": 3,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 7,
                                                 "row": 4,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 7,
                                                 "row": 5,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 7,
                                                 "row": 6,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 7,
                                                 "row": 7,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": true,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             }
                                         ]
                                     ],
                                     "clicks": 1,
                                     "flags": 0,
                                     "open_cells": 0,
                                     "covered_cells": 0
                                 }
                             }
                         },
                         {
                             "Status": 500,
                             "Message": "Game Name: 0xc0002b9830, UUID: 0xc0002b9860, Error: game not found"
                         }]`
		
##### Flag Click position

			`"request": {
				"method": "PUT",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\r\n    \"column\": 0,\r\n    \"row\": 1,\r\n    \"flag\": true\r\n}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "localhost:8080/game/movement/flag?name=myTestGame2&uuid=a57340e4-28bb-40c8-9897-3d0b2fe397a7",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"game",
						"movement"
					],
					"query": [
						{
							"key": "name",
							"value": "myTestGame"
						},
						{
							"key": "uuid",
							"value": "a57340e4-28bb-40c8-9897-3d0b2fe397a7"
						}
					]
				}
			},
			"response": [{
                             "Status": 200,
                             "Message": {
                                 "GameStatus": {
                                     "alive": true,
                                     "status": "Game in progress",
                                     "won": ""
                                 },
                                 "Cell": {
                                     "column": 1,
                                     "row": 0,
                                     "value": 0,
                                     "adjacent_bombs": 0,
                                     "mine": true,
                                     "clicked": false,
                                     "flagged": false,
                                     "opened": false,
                                     "covered": false
                                 },
                                 "Game": {
                                     "name": "myTestGame2",
                                     "uuid": "959de865-1304-4558-a310-a4a67ae92b2b",
                                     "rows": 8,
                                     "cols": 8,
                                     "mines": 10,
                                     "grid": [
                                         [
                                             {
                                                 "column": 0,
                                                 "row": 0,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 0,
                                                 "row": 1,
                                                 "value": 0,
                                                 "adjacent_bombs": 1,
                                                 "mine": false,
                                                 "clicked": true,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": true
                                             },
                                             {
                                                 "column": 0,
                                                 "row": 2,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 0,
                                                 "row": 3,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 0,
                                                 "row": 4,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 0,
                                                 "row": 5,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 0,
                                                 "row": 6,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 0,
                                                 "row": 7,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             }
                                         ],
                                         [
                                             {
                                                 "column": 1,
                                                 "row": 0,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": true,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 1,
                                                 "row": 1,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 1,
                                                 "row": 2,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 1,
                                                 "row": 3,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 1,
                                                 "row": 4,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 1,
                                                 "row": 5,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": true,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 1,
                                                 "row": 6,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 1,
                                                 "row": 7,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             }
                                         ],
                                         [
                                             {
                                                 "column": 2,
                                                 "row": 0,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 2,
                                                 "row": 1,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": true,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 2,
                                                 "row": 2,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 2,
                                                 "row": 3,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 2,
                                                 "row": 4,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 2,
                                                 "row": 5,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 2,
                                                 "row": 6,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 2,
                                                 "row": 7,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             }
                                         ],
                                         [
                                             {
                                                 "column": 3,
                                                 "row": 0,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 3,
                                                 "row": 1,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 3,
                                                 "row": 2,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": true,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 3,
                                                 "row": 3,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 3,
                                                 "row": 4,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": true,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 3,
                                                 "row": 5,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 3,
                                                 "row": 6,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 3,
                                                 "row": 7,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             }
                                         ],
                                         [
                                             {
                                                 "column": 4,
                                                 "row": 0,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 4,
                                                 "row": 1,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 4,
                                                 "row": 2,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 4,
                                                 "row": 3,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 4,
                                                 "row": 4,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": true,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 4,
                                                 "row": 5,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 4,
                                                 "row": 6,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 4,
                                                 "row": 7,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             }
                                         ],
                                         [
                                             {
                                                 "column": 5,
                                                 "row": 0,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 5,
                                                 "row": 1,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 5,
                                                 "row": 2,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 5,
                                                 "row": 3,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 5,
                                                 "row": 4,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 5,
                                                 "row": 5,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 5,
                                                 "row": 6,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 5,
                                                 "row": 7,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": true,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             }
                                         ],
                                         [
                                             {
                                                 "column": 6,
                                                 "row": 0,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 6,
                                                 "row": 1,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": true,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 6,
                                                 "row": 2,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": true,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 6,
                                                 "row": 3,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 6,
                                                 "row": 4,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 6,
                                                 "row": 5,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": true,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 6,
                                                 "row": 6,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 6,
                                                 "row": 7,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             }
                                         ],
                                         [
                                             {
                                                 "column": 7,
                                                 "row": 0,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 7,
                                                 "row": 1,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 7,
                                                 "row": 2,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 7,
                                                 "row": 3,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 7,
                                                 "row": 4,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 7,
                                                 "row": 5,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 7,
                                                 "row": 6,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 7,
                                                 "row": 7,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             }
                                         ]
                                     ],
                                     "clicks": 1,
                                     "flags": 0,
                                     "open_cells": 0,
                                     "covered_cells": 1
                                 }
                             }
                         },
                         {
                             "Status": 200,
                             "Message": {
                                 "GameStatus": {
                                     "alive": false,
                                     "status": "Game finished!",
                                     "won": "You lose!"
                                 },
                                 "Cell": {
                                     "column": 1,
                                     "row": 0,
                                     "value": 0,
                                     "adjacent_bombs": 0,
                                     "mine": false,
                                     "clicked": false,
                                     "flagged": false,
                                     "opened": false,
                                     "covered": false
                                 },
                                 "Game": {
                                     "name": "myTestGame2",
                                     "uuid": "fe02cd48-0962-40cb-b8e4-1300d00b6905",
                                     "rows": 8,
                                     "cols": 8,
                                     "mines": 10,
                                     "grid": [
                                         [
                                             {
                                                 "column": 0,
                                                 "row": 0,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 0,
                                                 "row": 1,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": true,
                                                 "clicked": true,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 0,
                                                 "row": 2,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 0,
                                                 "row": 3,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 0,
                                                 "row": 4,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 0,
                                                 "row": 5,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 0,
                                                 "row": 6,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": true,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 0,
                                                 "row": 7,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             }
                                         ],
                                         [
                                             {
                                                 "column": 1,
                                                 "row": 0,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 1,
                                                 "row": 1,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 1,
                                                 "row": 2,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": true,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 1,
                                                 "row": 3,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 1,
                                                 "row": 4,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": true,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 1,
                                                 "row": 5,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 1,
                                                 "row": 6,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 1,
                                                 "row": 7,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             }
                                         ],
                                         [
                                             {
                                                 "column": 2,
                                                 "row": 0,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 2,
                                                 "row": 1,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 2,
                                                 "row": 2,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 2,
                                                 "row": 3,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": true,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 2,
                                                 "row": 4,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 2,
                                                 "row": 5,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 2,
                                                 "row": 6,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 2,
                                                 "row": 7,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             }
                                         ],
                                         [
                                             {
                                                 "column": 3,
                                                 "row": 0,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": true,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 3,
                                                 "row": 1,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 3,
                                                 "row": 2,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 3,
                                                 "row": 3,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": true,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 3,
                                                 "row": 4,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 3,
                                                 "row": 5,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 3,
                                                 "row": 6,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 3,
                                                 "row": 7,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             }
                                         ],
                                         [
                                             {
                                                 "column": 4,
                                                 "row": 0,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 4,
                                                 "row": 1,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 4,
                                                 "row": 2,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 4,
                                                 "row": 3,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 4,
                                                 "row": 4,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 4,
                                                 "row": 5,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 4,
                                                 "row": 6,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 4,
                                                 "row": 7,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             }
                                         ],
                                         [
                                             {
                                                 "column": 5,
                                                 "row": 0,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 5,
                                                 "row": 1,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 5,
                                                 "row": 2,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 5,
                                                 "row": 3,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 5,
                                                 "row": 4,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": true,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 5,
                                                 "row": 5,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 5,
                                                 "row": 6,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 5,
                                                 "row": 7,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             }
                                         ],
                                         [
                                             {
                                                 "column": 6,
                                                 "row": 0,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 6,
                                                 "row": 1,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 6,
                                                 "row": 2,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 6,
                                                 "row": 3,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 6,
                                                 "row": 4,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": true,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 6,
                                                 "row": 5,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 6,
                                                 "row": 6,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 6,
                                                 "row": 7,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             }
                                         ],
                                         [
                                             {
                                                 "column": 7,
                                                 "row": 0,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 7,
                                                 "row": 1,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 7,
                                                 "row": 2,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 7,
                                                 "row": 3,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 7,
                                                 "row": 4,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 7,
                                                 "row": 5,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 7,
                                                 "row": 6,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": false,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             },
                                             {
                                                 "column": 7,
                                                 "row": 7,
                                                 "value": 0,
                                                 "adjacent_bombs": 0,
                                                 "mine": true,
                                                 "clicked": false,
                                                 "flagged": false,
                                                 "opened": false,
                                                 "covered": false
                                             }
                                         ]
                                     ],
                                     "clicks": 1,
                                     "flags": 0,
                                     "open_cells": 0,
                                     "covered_cells": 0
                                 }
                             }
                         },
                         {
                             "Status": 500,
                             "Message": "Game Name: 0xc0002b9830, UUID: 0xc0002b9860, Error: game not found"
                         }]`
                         
#####Pause Game
		`{
			"name": "Pause game",
			"request": {
				"method": "PUT",
				"header": [],
				"url": {
					"raw": "localhost:8080/game/pause?name=myTestGame",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"game",
						"pause"
					],
					"query": [
						{
							"key": "name",
							"value": "myTestGame"
						}
					]
				}
			},
			"response": [{
                                                      "Status": 200,
                                                      "Message": {
                                                          "alive": false,
                                                          "status": "Game paused.",
                                                          "won": "You lose!"
                                                      }
                                                  },
                                                  {
                                                      "Status": 400,
                                                      "Message": "game not found"
                                                  }]
		},`
		
#####Pause Game

		`{
			"name": "Resume game",
			"request": {
				"method": "PUT",
				"header": [],
				"url": {
					"raw": "localhost:8080/game/resume?name=myTestGame",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"game",
						"resume"
					],
					"query": [
						{
							"key": "name",
							"value": "myTestGame"
						}
					]
				}
			},
			"response": [
                         {
                             "Status": 400,
                             "Message": "game not found"
                         },{
                              "Status": 200,
                              "Message": {
                                  "alive": false,
                                  "status": "Game in progress",
                                  "won": "You lose!"
                              }
                          }],
                          "request": {
                          				"method": "PUT",
                          				"header": [],
                          				"body": {
                          					"mode": "raw",
                          					"raw": "{\r\n    \"column\": 0,\r\n    \"row\": 1,\r\n    \"flag\": false\r\n}",
                          					"options": {
                          						"raw": {
                          							"language": "json"
                          						}
                          					}
                          				},
                          				"url": {
                          					"raw": "localhost:8080/game/movement?name=myTestGame&uuid=a57340e4-28bb-40c8-9897-3d0b2fe397a7",
                          					"host": [
                          						"localhost"
                          					],
                          					"port": "8080",
                          					"path": [
                          						"game",
                          						"movement"
                          					],
                          					"query": [
                          						{
                          							"key": "name",
                          							"value": "myTestGame"
                          						},
                          						{
                          							"key": "uuid",
                          							"value": "a57340e4-28bb-40c8-9897-3d0b2fe397a7"
                          						}
                          					]
                          				}
                          			},
                          			"response": [{
                          				"Status": 200,
                          				"Message": {
                          					"alive": false,
                          					"status": "Game paused.",
                          					"won": "You lose!"
                          				}
                          			}]
		}`
}
### Backlog
**Backend Requierements:**
  - Setup environment
  - Generate connection with on memory api service
  - Generate ping controller
  - Push to heroku and validate working
  - Generate a new rest api on go with a default gin router
  - Define domain and add structs for game, cell, banner-status, actions
  - Add game constants, start, over, cell flag, cell positions, status, etc.
  - Add mappings & controllers, map each request unmarshall json body by request domain defined: 
    - POST: Start new game - weithht(column & rows) & (Number of mines)
    - POST: Restart current game game name
    - POST: Pause game
    - POST: Resume game
    - GET: Last Game Action
    - GET: Game Action History
    - POST: Post new click movement
    - GET: Game history details
  - Generate services:
    - Game service
      -> Generate a criteria value by params, validate some maximum dificult && dimensions
      -> Make the area
      -> Get randon position under the columns & rows maximums area, dividing by the maximum until get the count of bombs done.
      -> Validate actual position status to put  bomb and count
      -> Validate and add hints into others positions iterating by nodes
      -> Pause & Resume games
      -> End game validation for each position in status flagged, visited or revealed, if the difference between the number of bombs (not revealed) and the flag bombs safed is equals to the number of positions opens, win.
      -> End game by action validation exploit bomb
      -> Force endgame if timer is done
      -> Add time tracking service with channels
      -> Add movement x, y publish 
        -> Validation and exploit if bomb, end game
        -> Reveal hints if ok and reveal a random value area not visited next to the current visit, it should be just safe zone (WITHOUT bombs)
      -> Add flag movement x, y publish to ELUDE this position.
    - History service
  - Generate DAOS
    - Game
    
**Frontend - Client side - Pretended: flutter app with dart.**
  - Setup env
  - Setup rest providers with http calls
  - Setup routes & views
  - Build home view, if have a current game resume option, else start game
  - Build game start options view
  - Build board render view with maximum columns with some positions object
  - Build movements events on dashboard positions(onclick, doble click or right click or a flag)
  - Build game history view just list by user.
  - Deploy onto github with hummingbird
  
