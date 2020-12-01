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
  - Generate connection with mongo client
  - Generate ping controller
  - Push to heroku and validate working
  - Generate a new rest api on go with a defautl router
  - Define domain and add structs for account, user, game, current-game, cell, banner-status, session, actions
  - Add game constants, start, over, cell flag, cell positions, status, etc.
  - Add mappings & controllers, map each request unmarshall json body by request domain defined: 
    - POST: Create user. 
    - POST: Update user. 
    - POST: Get user. 
    - POST: Login
    - POST: Logout 
    - POST: Start new game - Time(minutes), weithht(column & rows) & (Difficulty || Number of mines)
    - POST: Pause game
    - POST: Resume game
    - POST: Abandon game
    - GET: Last Game Action
    - GET: Game Action History
    - POST: Post new action movement
    - GET: Game history details
  - Generate services:
    - User Service
      -> Create user
      -> Update user
      -> Get user
    - Auth service
      -> Login
      -> Logout
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
    - Time tracking
      -> Add time tracking service with channels
    - Action service
      -> Add movement x, y publish 
        -> Validation and exploit if bomb, end game
        -> Reveal hints if ok and reveal a random value area not visited next to the current visit, it should be just safe zone (WITHOUT bombs)
      -> Add flag movement x, y publish to ELUDE this position.
    - History service
  - Generate DAOS
    - User
    - Game
    - Action
    - Session
    
**Frontend - Client side - Pretended: flutter app with dart.**
  - Setup env
  - Setup rest providers with http calls
  - Setup routes & views
  - Build a login & logout flutter view
  - Build home view, if have a current game resume option, else start game
  - Build game start options view
  - Build board render view with maximum colums with some position object
  - Build movements events on dashboard positions(onclick, dobleclick or rightclick for flag)
  - Build game history view just list by user.
  - Deploy onto github with hummingbird
  
