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
	cols, _ := model.ToCreateMap()

	// Prepend i_id as the first column
	columns := append([]string{"i_id"}, cols...)

	// Build value placeholders for bulk insert
	valueStrings := make([]string, 0, len(req))
	valueArgs := make([]interface{}, 0, len(req)*len(columns))

	placeholderCount := 0
	for i := range req {
		// Generate and assign UUID
		req[i].ID = uuid.NewString()

		_, vals := req[i].ToCreateMap()

		// Prepend ID as the first value to match column order
		values := append([]interface{}{req[i].ID}, vals...)

		// Create placeholders for this row using @p1, @p2... for SQL Server
		placeholders := make([]string, len(columns))
		for j := range columns {
			placeholderCount++
			placeholders[j] = fmt.Sprintf("@p%d", placeholderCount)
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

	_, err := tx.ExecContext(ctx, query, valueArgs...)
	if err != nil {
		return fmt.Errorf("failed to bulk insert ro leads: %w", err)
	}

	return nil
}
