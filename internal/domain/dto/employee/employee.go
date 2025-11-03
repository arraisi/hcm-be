package employee

import "github.com/elgris/sqrl"

type GetEmployeeRequest struct {
	ID         *string
	EmployeeID *string
	Email      *string
}

// Apply applies the request parameters to the given SelectBuilder
func (req GetEmployeeRequest) Apply(q *sqrl.SelectBuilder) {
	if req.ID != nil {
		q.Where(sqrl.Eq{"i_id": req.ID})
	}
	if req.EmployeeID != nil {
		q.Where(sqrl.Eq{"i_employee_id": req.EmployeeID})
	}
	if req.Email != nil {
		q.Where(sqrl.Eq{"c_email": req.Email})
	}
}
