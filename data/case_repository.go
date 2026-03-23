package data

import (
	"database/sql"

	"prateekkuniyal.dev/desicrimepod/logger"
	"prateekkuniyal.dev/desicrimepod/models"
)

type CaseRepo struct {
	logger *logger.Logger
	db     *sql.DB
}

func NewCaseRepo(db *sql.DB, logger *logger.Logger) (*CaseRepo, error) {
	return &CaseRepo{
		db:     db,
		logger: logger,
	}, nil
}

const defaultLimit = 20

func (c *CaseRepo) GetTopCases() ([]models.Case, error) {
	query := `
	SELECT *
	FROM cases
	ORDER BY case_id DESC
	LIMIT $1
	`
	return c.getCases(query)
}

func (cs *CaseRepo) getCases(query string) ([]models.Case, error) {
	rows, err := cs.db.Query(query, defaultLimit)
	if err != nil {
		cs.logger.Error("Failed to query cases", err)
		return nil, err
	}

	defer rows.Close()

	var cases []models.Case
	for rows.Next() {
		var c models.Case
		if err := rows.Scan(
			&c.Id, &c.Name, &c.Summary, &c.OccuredOn, &c.Status, &c.Latitude, &c.Longitude, &c.CategoryId, &c.LocationId,
		); err != nil {
			cs.logger.Error("Failed to scan movie row", err)
			return nil, err
		}
		cases = append(cases, c)
	}
	return cases, nil

}
