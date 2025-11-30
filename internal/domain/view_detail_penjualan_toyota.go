package domain

type ViewDetailPenjualanToyota struct {
	Branch               string `json:"branch" db:"Branch"`
	BranchCode           string `json:"branch_code" db:"Branch Code"`
	Outlet               string `json:"outlet" db:"Outlet"`
	OutletCode           string `json:"outlet_code" db:"Outlet Code"`
	TypeTransaction      string `json:"type_transaction" db:"Type_transaction"`
	InvoiceDate          string `json:"invoice_date" db:"Invoice Date"`
	SPKDate              string `json:"spk_date" db:"SPK Date"`
	TanggalPanyerahanDEC string `json:"tanggal_panyerahan_dec" db:"Tanggal Panyerahan (DEC)"`
	Year                 int    `json:"year" db:"Year"`
	Period               string `json:"period" db:"Period"`
	Month                string `json:"month" db:"Month"`
	SalesBy              string `json:"sales_by" db:"Sales By"`
	PaymentType          string `json:"payment_type" db:"Payment Type"`
	CustomerID           string `json:"customer_id" db:"Customer Id"`
	CustomerName         string `json:"customer_name" db:"Customer Name"`
	CustomerGroup        string `json:"customer_group" db:"Customer Group"`
	CustomerAddress      string `json:"customer_address" db:"Customer Address"`
	City                 string `json:"city" db:"City"`
	Phone                string `json:"phone" db:"Phone"`
	SpvName              string `json:"spv_name" db:"Spv Name"`
	SlmName              string `json:"slm_name" db:"Slm Name"`
	ModelTAM             string `json:"model_tam" db:"Model (TAM)"`
	Model                string `json:"model" db:"Model"`
	ProductType          string `json:"product_type" db:"Product Type"`
	ProductID            string `json:"product_id" db:"Product Id"`
	Katashiki            string `json:"katashiki" db:"Katashiki"`
	Suffix               string `json:"suffix" db:"Suffix"`
	ProductName          string `json:"product_name" db:"Product Name"`
	Karoseri             string `json:"karoseri" db:"Karoseri"`
	Chassis              string `json:"chassis" db:"Chassis"`
	Engine               string `json:"engine" db:"Engine"`
	ProdYear             int    `json:"prod_year" db:"Prod Year"`
	ColourID             string `json:"colour_id" db:"Colour Id"`
	Colour               string `json:"colour" db:"Colour"`
	Tenor                int    `json:"tenor" db:"Tenor"`
	TypeAsuransi         string `json:"type_asuransi" db:"Type Asuransi"`
	NamaAsuransi         string `json:"nama_asuransi" db:"Nama Asuransi"`
	UnitSales            int    `json:"unit_sales" db:"Unit Sales"`
}

// TableName returns the database table name for the Order model
func (o *ViewDetailPenjualanToyota) TableName() string {
	return "dbo.VW_Detail_Penjualan_toyota"
}

// Columns returns the list of database columns for the Order model
func (o *ViewDetailPenjualanToyota) Columns() []string {
	return []string{
		"Branch",
		"Branch Code",
		"Outlet",
		"Outlet Code",
		"Type_transaction",
		"Invoice Date",
		"SPK Date",
		"Tanggal Panyerahan (DEC)",
		"Year",
		"Period",
		"Month",
		"Sales By",
		"Payment Type",
		"Customer Id",
		"Customer Name",
		"Customer Group",
		"Customer Address",
		"City",
		"Phone",
		"Spv Name",
		"Slm Name",
		"Model (TAM)",
		"Model",
		"Product Type",
		"Product Id",
		"Katashiki",
		"Suffix",
		"Product Name",
		"Karoseri",
		"Chassis",
		"Engine",
		"Prod Year",
		"Colour Id",
		"Colour",
		"Tenor",
		"Type Asuransi",
		"Nama Asuransi",
		"Unit Sales",
	}
}
