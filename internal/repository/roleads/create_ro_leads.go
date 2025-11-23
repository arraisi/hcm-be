package roleads

import (
	"context"
	"fmt"
	"strings"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func (r *repository) CreateRoLeads(ctx context.Context, tx *sqlx.Tx, req []domain.RoLeads) error {
	if len(req) == 0 {
		return nil
	}

	model := &domain.RoLeads{}
	columns, _ := model.ToCreateMap()
	columns = append(columns, "i_id")

	// Build value placeholders for bulk insert
	valueStrings := make([]string, 0, len(req))
	valueArgs := make([]interface{}, 0, len(req)*len(columns))

	placeholderCount := 0
	for _, roLead := range req {
		// Generate and assign UUID
		roLead.ID = uuid.NewString()

		_, values := roLead.ToCreateMap()
		values = append(values, roLead.ID)

		// Create placeholders for this row
		placeholders := make([]string, len(columns))
		for i := range columns {
			placeholderCount++
			placeholders[i] = fmt.Sprintf("$%d", placeholderCount)
		}
		valueStrings = append(valueStrings, fmt.Sprintf("(%s)", strings.Join(placeholders, ", ")))
		valueArgs = append(valueArgs, values...)
	}

	// Build the final query
	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES %s",
		model.TableName(),
		strings.Join(columns, ", "),
		strings.Join(valueStrings, ", "),
	)

	query = r.db.Rebind(query)
	_, err := tx.ExecContext(ctx, query, valueArgs...)
	if err != nil {
		return fmt.Errorf("failed to bulk insert ro leads: %w", err)
	}

	return nil
}
