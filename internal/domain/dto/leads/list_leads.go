package leads

import "time"

// ListLeadsRequest represents the filter parameters for listing leads
type ListLeadsRequest struct {
	// Stage filter (Lead, Prospect, Hot Prospect, SPK, DEC)
	Stage *string `json:"stage" form:"stage"`

	// Connection Status filter (Connected, Not Connected)
	ConnectionStatus *string `json:"connection_status" form:"connection_status"`

	// Follow up terakhir (Simulasi Kredit, Kirim Catalog, Kirim Promo, Test Drive, Appraisal Trade In, Rencana SPK)
	LastFollowUpType *string `json:"last_follow_up_type" form:"last_follow_up_type"`

	// Range tanggal lead dibuat
	CreatedDateFrom *time.Time `json:"created_date_from" form:"created_date_from"`
	CreatedDateTo   *time.Time `json:"created_date_to" form:"created_date_to"`

	// Source lead
	LeadSource *string `json:"lead_source" form:"lead_source"`

	// Follow up selanjutnya (range tanggal)
	NextFollowUpFrom *time.Time `json:"next_follow_up_from" form:"next_follow_up_from"`
	NextFollowUpTo   *time.Time `json:"next_follow_up_to" form:"next_follow_up_to"`

	// Pagination
	Page     int `json:"page" form:"page"`
	PageSize int `json:"page_size" form:"page_size"`
}

// ListLeadsResponse represents the response for list leads
type ListLeadsResponse struct {
	Data       []LeadListItem `json:"data"`
	Pagination PaginationMeta `json:"pagination"`
}

// LeadListItem represents a single lead item in the list
type LeadListItem struct {
	LeadsID              string     `json:"leads_id" db:"i_leads_id"`
	CustomerName         string     `json:"customer_name" db:"customer_name"`                     // Nama (FirstName + LastName)
	PhoneNumber          string     `json:"phone_number" db:"c_phone_number"`                     // No HP
	Stage                string     `json:"stage" db:"stage"`                                     // Stage (Lead, Prospect, Hot Prospect, SPK, DEC)
	LastFollowUpStatus   string     `json:"last_follow_up_status" db:"last_follow_up_status"`     // Status follow up terakhir
	LastFollowUpType     *string    `json:"last_follow_up_type" db:"last_follow_up_type"`         // Type of last follow up
	LastFollowUpDatetime *time.Time `json:"last_follow_up_datetime" db:"last_follow_up_datetime"` // When was the last follow up
	CreatedDatetime      *time.Time `json:"created_datetime" db:"created_datetime"`               // Tanggal lead dibuat
	NextFollowUpDate     *time.Time `json:"next_follow_up_date" db:"next_follow_up_date"`         // Follow up target date
	LeadSource           string     `json:"lead_source" db:"lead_source"`
	ConnectionStatus     string     `json:"connection_status" db:"connection_status"` // Connected/Not Connected
	OutletName           *string    `json:"outlet_name" db:"n_outlet_name"`
	Model                *string    `json:"model" db:"c_model"`
	Variant              *string    `json:"variant" db:"c_variant"`
}

// PaginationMeta represents pagination metadata
type PaginationMeta struct {
	Page       int   `json:"page"`
	PageSize   int   `json:"page_size"`
	TotalItems int64 `json:"total_items"`
	TotalPages int   `json:"total_pages"`
}
