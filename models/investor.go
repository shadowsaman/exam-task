package models

type InvestorPrimeryKey struct {
	Id string `json:"id"`
}

type CreateInvestor struct {
	Name string `json:"name"`
}

type Investor struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type InvestorFoyda struct {
	InvestorId string  `json:"investor_id"`
	Name       string  `json:"name"`
	Sum        float64 `json:"sum"`
}

type UpdateInvestor struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type UpdateInvestorSwag struct {
	Name string `json:"name"`
}

type GetListInvestorRequest struct {
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
}

type GetListInvestorResponse struct {
	Count     int64       `json:"count"`
	Investors []*Investor `json:"investors"`
}

type GetListInvestorResponseFoyda struct {
	Count         int64            `json:"count"`
	InvestorFoyda []*InvestorFoyda `json:"investors"`
}

type Empty struct{}
