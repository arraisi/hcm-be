package leads

import (
	"context"
	"database/sql"
	"errors"
	"github.com/arraisi/hcm-be/internal/domain"
	dtoLeads "github.com/arraisi/hcm-be/internal/domain/dto/leads"

	"github.com/jmoiron/sqlx"
)

func (s *service) UpsertLeads(ctx context.Context, tx *sqlx.Tx, req domain.Leads) (string, error) {
	leads, err := s.leadsRepo.GetLeads(ctx, dtoLeads.GetLeadsRequest{
		LeadsID: &req.LeadsID,
	})
	if err == nil {
		// Found → update
		req.ID = leads.ID
		err = s.leadsRepo.UpdateLeads(ctx, tx, req)
		if err != nil {
			return leads.ID, err
		}
		return leads.ID, nil
	}

	// Not found → create
	if errors.Is(err, sql.ErrNoRows) {
		err := s.leadsRepo.CreateLeads(ctx, tx, &req)
		if err != nil {
			return req.LeadsID, err
		}
		return req.LeadsID, nil
	}

	// other error
	return leads.ID, err
}
