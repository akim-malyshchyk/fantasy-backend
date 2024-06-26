package handlers

import "database/sql"

type HandlerContext struct {
	DB      *sql.DB
	BaseUrl string
}

type Tournament struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	DateStart string `json:"dateStart"`
	DateEnd   string `json:"dateEnd"`
	TownName  string `json:"town"`
}

type TournamentInfo struct {
	ID                 int            `json:"id"`
	Name               string         `json:"name"`
	DateStart          string         `json:"dateStart"`
	DateEnd            string         `json:"dateEnd"`
	DifficultyForecast float32        `json:"difficultyForecast"`
	QuestionQty        map[string]int `json:"questionQty"`
}

type TournamentTownData struct {
	TournamentId int `json:"id"`
	TownId       int `json:"idtown"`
}

type Team struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Town Town   `json:"town"`
}

type Country struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Town struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Country Country `json:"country"`
}

type TeamData struct {
	Team           Team    `json:"team"`
	Current        Team    `json:"current"`
	QuestionsTotal int     `json:"questionsTotal"`
	Position       float32 `json:"position"`
}
