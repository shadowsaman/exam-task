package models

type OrderPrimeryKey struct {
	Id string `json:"id"`
}

type CreateOrder struct {
	CarId      string  `json:"car_id"`
	ClientId   string  `json:"client_id"`
	TotalPrice float64 `json:"total_price"`
	PaidPrice  float64 `json:"paid_price"`
	DayCount   int     `json:"day_count"`
	GiveKm     int     `json:"give_km"`
	ReceiveKm  int     `json:"recieve_km"`
}

type Order struct {
	Id         string  `json:"id"`
	CarId      string  `json:"car_id"`
	ClientId   string  `json:"client_id"`
	TotalPrice float64 `json:"total_price"`
	PaidPrice  float64 `json:"paid_price"`
	DayCount   int     `json:"day_count"`
	GiveKm     int     `json:"give_km"`
	ReceiveKm  int     `json:"recieve_km"`
	Status     string  `json:"status"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
}

type UpdateOrder struct {
	Id         string  `json:"id"`
	CarId      string  `json:"car_id"`
	ClientId   string  `json:"client_id"`
	TotalPrice float64 `json:"total_price"`
	PaidPrice  float64 `json:"paid_price"`
	DayCount   int     `json:"day_count"`
	GiveKm     int     `json:"give_km"`
	ReceiveKm  int     `json:"recieve_km"`
	Status     string  `json:"status"`
}

type UpdateOrderSwag struct {
	CarId      string  `json:"car_id"`
	ClientId   string  `json:"client_id"`
	TotalPrice float64 `json:"total_price"`
	PaidPrice  float64 `json:"paid_price"`
	DayCount   int     `json:"day_count"`
	GiveKm     int     `json:"give_km"`
	ReceiveKm  int     `json:"recieve_km"`
	Status     string  `json:"status"`
}

type UpdateStatus struct {
	Id     string `json:"id"`
	Status string `json:"status"`
}

type UpdateStatusSwag struct {
	Status string `json:"status"`
}

type GetListOrderRequest struct {
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
}

type GetListOrderResponse struct {
	Count  int64    `json:"count"`
	Orders []*Order `json:"orders"`
}
