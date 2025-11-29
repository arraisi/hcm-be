package salesorder

import "github.com/elgris/sqrl"

type GetSalesOrderRequest struct {
	SoNumber *string `json:"so_number"`
	SpkID    *string `json:"spk_id"`
	OutletID *string `json:"outlet_id"`
}

// Apply applies the request parameters to the given SelectBuilder
func (req GetSalesOrderRequest) Apply(q *sqrl.SelectBuilder) {
	if req.SpkID != nil {
		q.Where(sqrl.Eq{"i_spk_id": req.SpkID})
	}
	if req.SoNumber != nil {
		q.Where(sqrl.Eq{"c_so_number": req.SoNumber})
	}
	if req.OutletID != nil {
		q.Where(sqrl.Eq{"i_outlet_id": req.OutletID})
	}
}

type GetSalesOrderAccessoriesRequest struct {
	SalesOrderID      *string `json:"sales_order_id"`
	PackageID         *string `json:"package_id"`
	AccessoriesNumber *string `json:"accessories_number"`
}

// Apply applies the request parameters to the given SelectBuilder
func (req GetSalesOrderAccessoriesRequest) Apply(q *sqrl.SelectBuilder) {
	if req.SalesOrderID != nil {
		q.Where(sqrl.Eq{"i_sales_order_id": req.SalesOrderID})
	}
	if req.PackageID != nil {
		q.Where(sqrl.Eq{"i_package_id": req.PackageID})
	}
	if req.AccessoriesNumber != nil {
		q.Where(sqrl.Eq{"c_accessories_number": req.AccessoriesNumber})
	}
}

type DeleteSalesOrderAccessoriesRequest struct {
	SalesOrderID *string `json:"sales_order_id"`
}

// Apply applies the request parameters to the given DeleteBuilder
func (req DeleteSalesOrderAccessoriesRequest) Apply(q *sqrl.DeleteBuilder) {
	if req.SalesOrderID != nil {
		q.Where(sqrl.Eq{"i_sales_order_id": req.SalesOrderID})
	}
}

type DeleteSalesOrderAccessoriesPartRequest struct {
	AccessoriesID     *string `json:"accessories_id"`
	AccessoriesNumber *string `json:"accessories_number"`
}

// Apply applies the request parameters to the given DeleteBuilder
func (req DeleteSalesOrderAccessoriesPartRequest) Apply(q *sqrl.DeleteBuilder) {
	if req.AccessoriesID != nil {
		q.Where(sqrl.Eq{"i_accessories_id": req.AccessoriesID})
	}
	if req.AccessoriesNumber != nil {
		q.Where(sqrl.Eq{"c_accessories_number": req.AccessoriesNumber})
	}
}

type DeleteSalesOrderDeliveryPlanRequest struct {
	SalesOrderID *string `json:"sales_order_id"`
}

// Apply applies the request parameters to the given DeleteBuilder
func (req DeleteSalesOrderDeliveryPlanRequest) Apply(q *sqrl.DeleteBuilder) {
	if req.SalesOrderID != nil {
		q.Where(sqrl.Eq{"i_sales_order_id": req.SalesOrderID})
	}
}

type DeleteSalesOrderInsurancePolicyRequest struct {
	SalesOrderID *string `json:"sales_order_id"`
}

// Apply applies the request parameters to the given DeleteBuilder
func (req DeleteSalesOrderInsurancePolicyRequest) Apply(q *sqrl.DeleteBuilder) {
	if req.SalesOrderID != nil {
		q.Where(sqrl.Eq{"i_sales_order_id": req.SalesOrderID})
	}
}

type DeleteSalesOrderPaymentRequest struct {
	SalesOrderID *string `json:"sales_order_id"`
}

// Apply applies the request parameters to the given DeleteBuilder
func (req DeleteSalesOrderPaymentRequest) Apply(q *sqrl.DeleteBuilder) {
	if req.SalesOrderID != nil {
		q.Where(sqrl.Eq{"i_sales_order_id": req.SalesOrderID})
	}
}
