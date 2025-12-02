package leads

import (
	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/elgris/sqrl"
)

// TradeInRequest represents the trade-in request
type TradeInRequest struct {
	LeadsID                      string   `json:"leads_id" validate:"required"`
	TradeInFlag                  bool     `json:"trade_in_flag"`
	VIN                          string   `json:"vin"`
	PoliceNumber                 string   `json:"police_number"`
	KatashikiSuffix              string   `json:"katashiki_suffix"`
	ColorCode                    string   `json:"color_code"`
	UsedCarBrand                 string   `json:"used_car_brand"`
	UsedCarModel                 string   `json:"used_car_model"`
	UsedCarVariant               string   `json:"used_car_variant"`
	UsedCarColor                 string   `json:"used_car_color"`
	InitialEstimatedUsedCarPrice *float64 `json:"initial_estimated_used_car_price"`
	Mileage                      *float64 `json:"mileage"`
	Province                     string   `json:"province"`
	UsedCarYear                  *float64 `json:"used_car_year"`
	UsedCarTransmission          string   `json:"used_car_transmission"`
	UsedCarFuel                  string   `json:"used_car_fuel"`
	UsedCarEngineCapacity        *float64 `json:"used_car_engine_capacity"`
	UsedCarMover                 string   `json:"used_car_mover"`
}

// GetTradeInRequest represents the request parameters for getting a trade-in
type GetTradeInRequest struct {
	ID           *string
	LeadsID      *string
	VIN          *string
	PoliceNumber *string
}

// Apply applies the request parameters to the given SelectBuilder
func (req GetTradeInRequest) Apply(q *sqrl.SelectBuilder) {
	if req.ID != nil {
		q.Where(sqrl.Eq{"i_id": *req.ID})
	}
	if req.LeadsID != nil {
		q.Where(sqrl.Eq{"i_leads_id": *req.LeadsID})
	}
	if req.VIN != nil {
		q.Where(sqrl.Eq{"c_vin": *req.VIN})
	}
	if req.PoliceNumber != nil {
		q.Where(sqrl.Eq{"c_police_number": *req.PoliceNumber})
	}
}

// GetTradeInsRequest represents the request parameters for getting multiple trade-ins
type GetTradeInsRequest struct {
	LeadsID  *string
	Page     int
	PageSize int
}

// Apply applies the request parameters to the given SelectBuilder
func (req GetTradeInsRequest) Apply(q *sqrl.SelectBuilder) {
	if req.LeadsID != nil {
		q.Where(sqrl.Eq{"i_leads_id": *req.LeadsID})
	}

	if req.PageSize > 0 {
		// Calculate offset: (page - 1) * pageSize
		offset := 0
		if req.Page > 1 {
			offset = (req.Page - 1) * req.PageSize
		}
		// Use pageSize + 1 to detect if there's a next page
		limit := req.PageSize + 1
		q.Suffix("OFFSET ? ROWS FETCH NEXT ? ROWS ONLY", offset, limit)
	}
}

// GetTradeInsResponse represents the response for getting multiple trade-ins
type GetTradeInsResponse struct {
	Data       []domain.LeadsTradeIn `json:"data"`
	Pagination TradeInPagination     `json:"pagination"`
}

// TradeInPagination represents pagination information for trade-ins
type TradeInPagination struct {
	Page     int  `json:"page"`
	PageSize int  `json:"page_size"`
	HasNext  bool `json:"has_next"`
}

// ToDomain converts the request to the internal LeadsTradeIn model
func (req *TradeInRequest) ToDomain(leadsID string) domain.LeadsTradeIn {
	return domain.LeadsTradeIn{
		LeadsID:                      leadsID,
		TradeInFlag:                  req.TradeInFlag,
		VIN:                          req.VIN,
		PoliceNumber:                 req.PoliceNumber,
		KatashikiSuffix:              req.KatashikiSuffix,
		ColorCode:                    req.ColorCode,
		UsedCarBrand:                 req.UsedCarBrand,
		UsedCarModel:                 req.UsedCarModel,
		UsedCarVariant:               req.UsedCarVariant,
		UsedCarColor:                 req.UsedCarColor,
		InitialEstimatedUsedCarPrice: req.InitialEstimatedUsedCarPrice,
		Mileage:                      req.Mileage,
		Province:                     req.Province,
		UsedCarYear:                  req.UsedCarYear,
		UsedCarTransmission:          req.UsedCarTransmission,
		UsedCarFuel:                  req.UsedCarFuel,
		UsedCarEngineCapacity:        req.UsedCarEngineCapacity,
		UsedCarMover:                 req.UsedCarMover,
	}
}
