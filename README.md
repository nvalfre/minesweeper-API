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
  
