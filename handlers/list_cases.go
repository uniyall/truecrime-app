package handlers

import (
	"encoding/json"
	"net/http"

	"prateekkuniyal.dev/desicrimepod/models"
)

type CasesHandler struct {
	Api string
}

func (c *CasesHandler) writeJSONResponse(w http.ResponseWriter, data any) {
	w.Header().Add("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	if err := enc.Encode(data); err != nil {
		// error log this
	}
}

func (c *CasesHandler) GetCases(w http.ResponseWriter, r *http.Request) {
	cases := []models.Case{
		{
			ID:       "1",
			Name:     "2008 Noida Double Murder Case (Arushi Talwar)",
			State:    "Uttar Pradesh",
			Year:     2008,
			YtVideos: nil,
		},
		{
			ID:       "2",
			Name:     "Burari Deaths",
			State:    "Delhi",
			Year:     2018,
			YtVideos: nil,
		},
		{
			ID:       "3",
			Name:     "2006 Noida serial murders",
			State:    "Uttar Pradesh",
			Year:     2006,
			YtVideos: nil,
		},
		{
			ID:       "4",
			Name:     "Bikini Killer Cases",
			State:    "Uttar Pradesh",
			Year:     1980,
			YtVideos: nil,
		},
	}

	c.writeJSONResponse(w, cases)
}

func (c *CasesHandler) GetRandomCases(w http.ResponseWriter, r *http.Request) {
	randomCase := models.Case{
		ID:       "1",
		Name:     "2008 Noida Double Murder Case (Arushi Talwar)",
		State:    "Uttar Pradesh",
		Year:     2008,
		YtVideos: nil,
	}

	c.writeJSONResponse(w, randomCase)

}
