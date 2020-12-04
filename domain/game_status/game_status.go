package game_status

const WaitingForStart = "Start game"
const Started = "Game in progress"
const Paused = "Game paused."
const Finished = "Game finished!"
const Lose = "You lose!"
const Won = "You Won!"

type GameStatus struct {
	Status string `json:"status"`
	Won    string `json:"won"`
}
