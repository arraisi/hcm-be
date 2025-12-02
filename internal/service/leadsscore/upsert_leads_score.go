package usedcar

import (
	"context"
	"database/sql"
	"errors"
	"github.com/arraisi/hcm-be/internal/domain"

	"github.com/jmoiron/sqlx"
)

func (s *service) UpsertLeadsScore(ctx context.Context, tx *sqlx.Tx, req domain.LeadsScore) (string, error) {
	leadsScore, err := s.repo.GetLeadsScore(ctx, req.LeadsID)
	if err == nil {
		// Found → update
		req.ID = leadsScore.ID
		err = s.repo.UpdateLeadsScore(ctx, tx, req)
		if err != nil {
			return leadsScore.ID, err
		}
		return leadsScore.ID, nil
	}

	// Not found → create
	if errors.Is(err, sql.ErrNoRows) {
		leadsScoreID, err := s.repo.CreateLeadsScore(ctx, tx, &req)
		if err != nil {
			return leadsScoreID, err
		}
		return leadsScoreID, nil
	}

	// other error
	return leadsScore.ID, err
}
