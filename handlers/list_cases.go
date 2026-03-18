package handlers

import (
	"encoding/json"
	"net/http"

	"prateekkuniyal.dev/desicrimepod/models"
)

type GetCases struct{}


func (c GetCases) ServeHTTP(w http.ResponseWriter,r *http.Request) {
	myCases := []models.Case {
		{
			ID: "1",
			Name: "2008 Noida Double Murder Case",
			State: "Uttar Pradesh",
			Year: 2008,
			YtVideos: nil,
		},
	}

	w.Header().Add("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	if err := enc.Encode(myCases); err != nil {
		// error log this 
	}
}