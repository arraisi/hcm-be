package creditsimulation

type BranchResponse struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

type OutletResponse struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

type AssetGroupResponse struct {
	AssetGroupCode string `json:"assetGroupCode"`
	AssetGroupName string `json:"assetGroupName"`
	AssetKindCode  string `json:"assetKindCode"`
	AssetKindName  string `json:"assetKindName"`
}

type AssetTypeResponse struct {
	AssetTypeCode           string  `json:"assetTypeCode"`
	AssetTypeName           string  `json:"assetTypeName"`
	Price                   float64 `json:"price"`
	DpBottomLimit           float64 `json:"dpBottomLimit"`
	DpTopLimit              float64 `json:"dpTopLimit"`
	AdministrationJournalTo string `json:"administrationJournalTo"`
	MinAdminLiability       string `json:"minAdminLiability"` 
	MaxAdminLiability       string `json:"maxAdminLiability"`
}

type InsuranceAssetType struct {
	Sequence int    `json:"sequence"`
	Code     string `json:"code"`
	Audit    interface{} `json:"audit"`
	RowNum   interface{} `json:"rowNum"`
	Name     interface{} `json:"name"`
	Category interface{} `json:"category"`
}

type InstallmentResponse struct {
	PriceListId             string `json:"priceListId"`
	PriceListTitle          string `json:"priceListTitle"`
	AssetTypeCode           string `json:"assetTypeCode"`
	AssetTypeName           string `json:"assetTypeName"`
	Price                   string `json:"price"` 
	MinAdminLiability       float64 `json:"minAdminLiability"`
	MaxAdminLiability       float64 `json:"maxAdminLiability"`
	DpBottomLimitPercentage float64 `json:"dpBottomLimitPercentage"`
	DpBottomLimit           string `json:"dpBottomLimit"` 
	DpTopLimitPercentage    float64 `json:"dpTopLimitPercentage"`
	DpTopLimit              string `json:"dpTopLimit"` 
	InstallmentMin          string `json:"installmentMin"`
	InstallmentMax          string `json:"installmentMax"` 
	ContractPaymentType     interface{} `json:"contractPaymentType"` 
	InsuranceYearPaidDetail int    `json:"insuranceYearPaidDetail"`
	MinInsuranceAmount      float64 `json:"minInsuranceAmount"`
	MaxInsuranceAmount      float64 `json:"maxInsuranceAmount"`
	MinTenor                int    `json:"minTenor"`
	MaxTenor                int    `json:"maxTenor"`
	MinInterestRate         float64 `json:"minInterestRate"`
	MaxInterestRate         float64 `json:"maxInterestRate"`
	BungaTenor12            int    `json:"bungaTenor12"`
	BungaTenor36            int    `json:"bungaTenor36"`
	InsuranceAssetType      InsuranceAssetType `json:"insuranceAssetType"`
}

type AggregatedInstallmentResponse struct {
	InstallmentMin string `json:"installmentMin"`
	InstallmentMax string `json:"installmentMax"`
}

type CreditSimulationDetailResponse struct {
	TenorName   string `json:"tenorName"`
	DownPayment string `json:"downPayment"` 
	Installment string `json:"installment"` 
	PriceListId string `json:"priceListId"`
	TenorVale   int    `json:"tenorVale"`   
	Title       string `json:"title"`
}