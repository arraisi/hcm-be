package leads

import (
	"context"
	"fmt"

	"github.com/arraisi/hcm-be/internal/domain/dto/leads"
	"github.com/elgris/sqrl"
)

// ListLeads retrieves a paginated list of leads with filters
func (r *repository) ListLeads(ctx context.Context, req leads.ListLeadsRequest) ([]leads.LeadListItem, int64, error) {
	var result []leads.LeadListItem

	// Build base query with joins
	query := sqrl.
		Select(
			"l.i_leads_id",
			"CONCAT(c.n_first_name, ' ', c.n_last_name) AS customer_name",
			"c.c_phone_number",
			"l.c_tam_lead_score AS stage",
			"l.c_leads_follow_up_status AS last_follow_up_status",
			"l.c_leads_type AS last_follow_up_type",
			"l.d_created_datetime AS created_datetime",
			"l.c_leads_source AS lead_source",
			"CASE WHEN l.c_leads_follow_up_status IN ('CONTACTED', 'ON_CONSIDERATION', 'INTERESTED') THEN 'Connected' ELSE 'Not Connected' END AS connection_status",
			"l.n_outlet_name",
			"l.c_model",
			"l.c_variant",
		).
		From("dbo.tm_leads l").
		LeftJoin("dbo.tr_customer c ON l.i_customer_id = c.i_id").
		OrderBy("l.d_created_datetime DESC")

	// Apply filters
	if req.Stage != nil && *req.Stage != "" {
		query = query.Where(sqrl.Eq{"l.c_tam_lead_score": *req.Stage})
	}

	if req.ConnectionStatus != nil && *req.ConnectionStatus != "" {
		if *req.ConnectionStatus == "Connected" {
			query = query.Where("l.c_leads_follow_up_status IN ('CONTACTED', 'ON_CONSIDERATION', 'INTERESTED')")
		} else {
			query = query.Where("l.c_leads_follow_up_status NOT IN ('CONTACTED', 'ON_CONSIDERATION', 'INTERESTED')")
		}
	}

	if req.LastFollowUpType != nil && *req.LastFollowUpType != "" {
		query = query.Where(sqrl.Eq{"l.c_leads_type": *req.LastFollowUpType})
	}

	if req.LeadSource != nil && *req.LeadSource != "" {
		query = query.Where(sqrl.Eq{"l.c_leads_source": *req.LeadSource})
	}

	if req.CreatedDateFrom != nil {
		query = query.Where(sqrl.GtOrEq{"l.d_created_datetime": *req.CreatedDateFrom})
	}

	if req.CreatedDateTo != nil {
		query = query.Where(sqrl.LtOrEq{"l.d_created_datetime": *req.CreatedDateTo})
	}

	// Build and execute query
	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return nil, 0, fmt.Errorf("failed to build query: %w", err)
	}

	// Add SQL Server pagination using OFFSET-FETCH
	offset := (req.Page - 1) * req.PageSize
	sqlQuery = fmt.Sprintf("%s OFFSET %d ROWS FETCH NEXT %d ROWS ONLY", sqlQuery, offset, req.PageSize)

	sqlQuery = r.db.Rebind(sqlQuery)
	err = r.db.SelectContext(ctx, &result, sqlQuery, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to execute query: %w", err)
	}

	// Get total count
	countQuery := sqrl.
		Select("COUNT(*)").
		From("dbo.tm_leads l").
		LeftJoin("dbo.tr_customer c ON l.i_customer_id = c.i_id")

	// Apply same filters to count query
	if req.Stage != nil && *req.Stage != "" {
		countQuery = countQuery.Where(sqrl.Eq{"l.c_tam_lead_score": *req.Stage})
	}
	if req.ConnectionStatus != nil && *req.ConnectionStatus != "" {
		if *req.ConnectionStatus == "Connected" {
			countQuery = countQuery.Where("l.c_leads_follow_up_status IN ('CONTACTED', 'ON_CONSIDERATION', 'INTERESTED')")
		} else {
			countQuery = countQuery.Where("l.c_leads_follow_up_status NOT IN ('CONTACTED', 'ON_CONSIDERATION', 'INTERESTED')")
		}
	}
	if req.LastFollowUpType != nil && *req.LastFollowUpType != "" {
		countQuery = countQuery.Where(sqrl.Eq{"l.c_leads_type": *req.LastFollowUpType})
	}
	if req.LeadSource != nil && *req.LeadSource != "" {
		countQuery = countQuery.Where(sqrl.Eq{"l.c_leads_source": *req.LeadSource})
	}
	if req.CreatedDateFrom != nil {
		countQuery = countQuery.Where(sqrl.GtOrEq{"l.d_created_datetime": *req.CreatedDateFrom})
	}
	if req.CreatedDateTo != nil {
		countQuery = countQuery.Where(sqrl.LtOrEq{"l.d_created_datetime": *req.CreatedDateTo})
	}

	countSQL, countArgs, err := countQuery.ToSql()
	if err != nil {
		return nil, 0, fmt.Errorf("failed to build count query: %w", err)
	}

	var total int64
	countSQL = r.db.Rebind(countSQL)
	err = r.db.GetContext(ctx, &total, countSQL, countArgs...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get total count: %w", err)
	}

	return result, total, nil
}
