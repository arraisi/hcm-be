package domain

type Employee struct {
	ID           string `db:"i_id"`
	EmployeeID   string `db:"i_employee_id"`
	EmployeeName string `db:"n_employee_name"`
	Email        string `db:"c_email"`
	PhoneNumber  string `db:"c_phone_number"`
}

func (e *Employee) TableName() string {
	return "tr_employee"
}

func (e *Employee) Columns() []string {
	return []string{
		"i_id",
		"i_employee_id",
		"n_employee_name",
		"c_email",
		"c_phone_number",
	}
}

func (e *Employee) SelectColumns() []string {
	return []string{
		"CAST(i_id AS CHAR) AS i_id",
		"CAST(i_employee_id AS CHAR) AS i_employee_id",
		"n_employee_name",
		"c_email",
		"c_phone_number",
	}
}

func (e *Employee) ToCreateMap() ([]string, []interface{}) {
	columns := make([]string, 0, len(e.SelectColumns()))
	values := make([]interface{}, 0, len(e.SelectColumns()))

	columns = append(columns, "i_employee_id")
	values = append(values, e.EmployeeID)

	columns = append(columns, "n_employee_name")
	values = append(values, e.EmployeeName)

	columns = append(columns, "c_email")
	values = append(values, e.Email)

	columns = append(columns, "c_phone_number")
	values = append(values, e.PhoneNumber)

	return columns, values
}

func (e *Employee) ToUpdateMap() map[string]interface{} {
	updateMap := make(map[string]interface{})

	updateMap["i_employee_id"] = e.EmployeeID
	updateMap["n_employee_name"] = e.EmployeeName
	updateMap["c_email"] = e.Email
	updateMap["c_phone_number"] = e.PhoneNumber

	return updateMap
}
