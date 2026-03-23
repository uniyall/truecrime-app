package models

type Case struct {
	Id         int     `json:"case_id"`
	Name       string  `json:"title"`
	Summary    string  `json:"summary"`
	OccuredOn  string  `json:"occured_on"`
	Status     string  `json:"status"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	CategoryId int     `json:"category_id"`
	LocationId int     `json:"location_id"`
}

type Category struct {
	Id   string `json:"category_id"`
	Name string `json:"name"`
}
