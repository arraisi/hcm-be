package domain

import "time"

type CustomerReminder struct {
	ID                string `db:"i_id"`
	CustomerID        string `db:"i_customer_id"`         // FK to tr_customer (NOT NULL)
	CustomerVehicleID string `db:"i_customer_vehicle_id"` // FK to tm_customer_vehicle (NOT NULL)

	// Reminder detail (from Toyota webhook)
	OutletID                string     `db:"i_outlet_id"`                    // NOT NULL
	ExternalReminderID      string     `db:"i_reminder_id"`                  // "reminder_ID" from payload (NOT NULL)
	Activity                string     `db:"c_activity"`                     // SERVICE_BOOKING, etc (NOT NULL)
	ActivityPlanScheduledAt *time.Time `db:"d_activity_plan_scheduled_date"` // NULL
	AutoReminderStatus      *string    `db:"c_auto_reminder_status"`         // DELIVERED, READ, NON_MTOYOTA (NULL)
	ReminderMessage         *string    `db:"e_reminder_message"`             // NULL
	PriorityCall            *int       `db:"v_priority_call"`                // NULL
	ExtendedWarrantyStatus  *string    `db:"c_extended_warranty_status"`     // ELIGIBLE, NOT_ELIGIBLE, NOT_APPLICABLE (NULL)
	CustomerHabit           *string    `db:"c_customer_habit"`               // TIME_BASED, MILEAGE (NULL)
	LastHabit               *string    `db:"c_last_habit"`                   // PUNCTUAL, EARLY, LATE, INACTIVE, PASSIVE (NULL)
	NextServiceStatus       *string    `db:"c_next_service_status"`          // same values as last_habit (NULL)
	LastServiceDate         *time.Time `db:"d_last_service_date"`            // nullable
	NextServiceDate         *time.Time `db:"d_next_service_date"`            // nullable
	NCSStatus               *string    `db:"c_ncs_status"`                   // SAME_OUTLET, DIFFERENT (NULL)
	ProgramTab              *string    `db:"c_program_tab"`                  // T_CARE, GBSB, REGULAR (NULL)
	NextServiceStage        *int       `db:"v_next_service_stage"`           // NULL

	// Audit
	CreatedAt time.Time  `db:"d_created_at"` // NOT NULL (has default)
	CreatedBy *string    `db:"c_created_by"` // NULL
	UpdatedAt *time.Time `db:"d_updated_at"` // NULL (default getdate())
	UpdatedBy *string    `db:"c_updated_by"` // NULL
}

// TableName returns the database table name for the CustomerReminder model.
func (r *CustomerReminder) TableName() string {
	return "dbo.tr_customer_reminder"
}

// Columns returns the list of database columns for the CustomerReminder model.
func (r *CustomerReminder) Columns() []string {
	return []string{
		"i_id",
		"i_customer_id",
		"i_customer_vehicle_id",
		"i_outlet_id",
		"i_reminder_id",
		"c_activity",
		"d_activity_plan_scheduled_date",
		"c_auto_reminder_status",
		"e_reminder_message",
		"v_priority_call",
		"c_extended_warranty_status",
		"c_customer_habit",
		"c_last_habit",
		"c_next_service_status",
		"d_last_service_date",
		"d_next_service_date",
		"c_ncs_status",
		"c_program_tab",
		"v_next_service_stage",
		"d_created_at",
		"c_created_by",
		"d_updated_at",
		"c_updated_by",
	}
}

// SelectColumns returns the list of columns to select in queries for the CustomerReminder model.
func (r *CustomerReminder) SelectColumns() []string {
	return []string{
		"i_id",
		"i_outlet_id",
		"i_customer_id",
		"i_customer_vehicle_id",
		"i_reminder_id",
		"c_activity",
		"c_auto_reminder_status",
		"d_activity_plan_scheduled_date",
		"d_created_at",
		"c_created_by",
		"d_updated_at",
		"c_updated_by",
	}
}

func (r *CustomerReminder) ToCreateMap() (columns []string, values []interface{}) {
	columns = make([]string, 0, len(r.Columns()))
	values = make([]interface{}, 0, len(r.Columns()))

	if r.OutletID != "" {
		columns = append(columns, "i_outlet_id")
		values = append(values, r.OutletID)
	}
	if r.CustomerID != "" {
		columns = append(columns, "i_customer_id")
		values = append(values, r.CustomerID)
	}
	if r.CustomerVehicleID != "" {
		columns = append(columns, "i_customer_vehicle_id")
		values = append(values, r.CustomerVehicleID)
	}
	if r.ExternalReminderID != "" {
		columns = append(columns, "i_reminder_id")
		values = append(values, r.ExternalReminderID)
	}
	if r.Activity != "" {
		columns = append(columns, "c_activity")
		values = append(values, r.Activity)
	}
	if r.ActivityPlanScheduledAt != nil && !r.ActivityPlanScheduledAt.IsZero() {
		columns = append(columns, "d_activity_plan_scheduled_date")
		values = append(values, r.ActivityPlanScheduledAt.UTC())
	}
	if r.AutoReminderStatus != nil && *r.AutoReminderStatus != "" {
		columns = append(columns, "c_auto_reminder_status")
		values = append(values, *r.AutoReminderStatus)
	}
	if r.ReminderMessage != nil && *r.ReminderMessage != "" {
		columns = append(columns, "e_reminder_message")
		values = append(values, *r.ReminderMessage)
	}
	if r.PriorityCall != nil && *r.PriorityCall != 0 {
		columns = append(columns, "v_priority_call")
		values = append(values, *r.PriorityCall)
	}
	if r.ExtendedWarrantyStatus != nil && *r.ExtendedWarrantyStatus != "" {
		columns = append(columns, "c_extended_warranty_status")
		values = append(values, *r.ExtendedWarrantyStatus)
	}
	if r.CustomerHabit != nil && *r.CustomerHabit != "" {
		columns = append(columns, "c_customer_habit")
		values = append(values, *r.CustomerHabit)
	}
	if r.LastHabit != nil && *r.LastHabit != "" {
		columns = append(columns, "c_last_habit")
		values = append(values, *r.LastHabit)
	}
	if r.NextServiceStatus != nil && *r.NextServiceStatus != "" {
		columns = append(columns, "c_next_service_status")
		values = append(values, *r.NextServiceStatus)
	}
	if r.LastServiceDate != nil && !r.LastServiceDate.IsZero() {
		columns = append(columns, "d_last_service_date")
		values = append(values, r.LastServiceDate.UTC())
	}
	if r.NextServiceDate != nil && !r.NextServiceDate.IsZero() {
		columns = append(columns, "d_next_service_date")
		values = append(values, r.NextServiceDate.UTC())
	}
	if r.NCSStatus != nil && *r.NCSStatus != "" {
		columns = append(columns, "c_ncs_status")
		values = append(values, *r.NCSStatus)
	}
	if r.ProgramTab != nil && *r.ProgramTab != "" {
		columns = append(columns, "c_program_tab")
		values = append(values, *r.ProgramTab)
	}
	if r.NextServiceStage != nil && *r.NextServiceStage != 0 {
		columns = append(columns, "v_next_service_stage")
		values = append(values, *r.NextServiceStage)
	}

	// audit fields always set
	columns = append(columns, "d_created_at")
	values = append(values, r.CreatedAt.UTC())

	columns = append(columns, "c_created_by")
	values = append(values, r.CreatedBy) // pointer is fine; nil => NULL

	columns = append(columns, "d_updated_at")
	if r.UpdatedAt != nil {
		values = append(values, r.UpdatedAt.UTC())
	} else {
		values = append(values, nil)
	}

	columns = append(columns, "c_updated_by")
	values = append(values, r.UpdatedBy)

	return columns, values
}

func (r *CustomerReminder) ToUpdateMap() map[string]interface{} {
	updateMap := make(map[string]interface{})

	if r.AutoReminderStatus != nil && *r.AutoReminderStatus != "" {
		updateMap["c_auto_reminder_status"] = *r.AutoReminderStatus
	}
	if r.ReminderMessage != nil && *r.ReminderMessage != "" {
		updateMap["e_reminder_message"] = *r.ReminderMessage
	}
	if r.CustomerHabit != nil && *r.CustomerHabit != "" {
		updateMap["c_customer_habit"] = *r.CustomerHabit
	}
	if r.LastHabit != nil && *r.LastHabit != "" {
		updateMap["c_last_habit"] = *r.LastHabit
	}
	if r.NextServiceStatus != nil && *r.NextServiceStatus != "" {
		updateMap["c_next_service_status"] = *r.NextServiceStatus
	}
	if r.LastServiceDate != nil && !r.LastServiceDate.IsZero() {
		updateMap["d_last_service_date"] = r.LastServiceDate.UTC()
	}
	if r.NextServiceDate != nil && !r.NextServiceDate.IsZero() {
		updateMap["d_next_service_date"] = r.NextServiceDate.UTC()
	}
	if r.NCSStatus != nil && *r.NCSStatus != "" {
		updateMap["c_ncs_status"] = *r.NCSStatus
	}
	if r.ProgramTab != nil && *r.ProgramTab != "" {
		updateMap["c_program_tab"] = *r.ProgramTab
	}
	if r.NextServiceStage != nil && *r.NextServiceStage != 0 {
		updateMap["v_next_service_stage"] = *r.NextServiceStage
	}

	if r.UpdatedAt != nil {
		updateMap["d_updated_at"] = r.UpdatedAt.UTC()
	} else {
		updateMap["d_updated_at"] = nil
	}
	updateMap["c_updated_by"] = r.UpdatedBy

	return updateMap
}
