package handlers

import (
	"encoding/json"
	"net/http"

	"prateekkuniyal.dev/desicrimepod/data"
	"prateekkuniyal.dev/desicrimepod/logger"
)

type CasesHandler struct {
	storage data.MovieStorage
	logger  *logger.Logger
}

func NewCaseHandler(storage data.MovieStorage, logger *logger.Logger) *CasesHandler {
	return &CasesHandler{
		storage: storage,
		logger:  logger,
	}
}

func (c *CasesHandler) writeJSONResponse(w http.ResponseWriter, data any) {
	w.Header().Add("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	if err := enc.Encode(data); err != nil {
		c.logger.Error("Unable to encode response to JSON", err)
	}
}

func (c *CasesHandler) GetCases(w http.ResponseWriter, r *http.Request) {
	cases, err := c.storage.GetTopCases()
	if err != nil {
		http.Error(w, "Unable to fetch data", http.StatusInternalServerError)
		c.logger.Error("Unable to obtain cases from Storage", err)
	}
	c.writeJSONResponse(w, cases)
}

func (c *CasesHandler) GetRandomCases(w http.ResponseWriter, r *http.Request) {
	randomCase, err := c.storage.GetTopCases()
	if err != nil {
		http.Error(w, "Unable to fetch data", http.StatusInternalServerError)
		c.logger.Error("Unable to obtain cases from Storage", err)
	}

	c.writeJSONResponse(w, randomCase)

}
