package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"app/models"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/google/uuid"
)

type OrderRepo struct {
	db *pgxpool.Pool
}

func NewOrderRepo(db *pgxpool.Pool) *OrderRepo {
	return &OrderRepo{
		db: db,
	}
}

func (r *OrderRepo) Insert(ctx context.Context, order *models.CreateOrder) (string, error) {

	var (
		id = uuid.New().String()
	)

	query := `
		INSERT INTO "order" (
			id,
			car_id,
			client_id,
			total_price,
			paid_price,
			day_count,
			give_km,
			receive_km,
			updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, now())
	`

	_, err := r.db.Exec(ctx, query,
		id,
		order.CarId,
		order.ClientId,
		order.TotalPrice,
		order.PaidPrice,
		order.DayCount,
		order.GiveKm,
		order.ReceiveKm,
	)

	if err != nil {
		return "", err
	}

	var (
		km    = order.CarId
		price = order.CarId
	)

	Querykm := `
			UPDATE "order" 
			SET    
					give_km = km
			FROM   car 
			`
	if km != "" {
		km = fmt.Sprintf(`WHERE  "order".car_id =  '%s'%s`, order.CarId, " AND  km IS DISTINCT FROM give_km ")
		Querykm += km
	}

	_, err = r.db.Exec(ctx, Querykm)
	if err != nil {
		return "", err
	}

	QueryTotalPrice := `
		UPDATE "order"
		SET
			total_price = price * day_count
		from car 
	`
	if price != "" {
		price = fmt.Sprintf(`WHERE "order".car_id = '%s'`, order.CarId)
		QueryTotalPrice += price
	}

	_, err = r.db.Exec(ctx, QueryTotalPrice)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *OrderRepo) GetByID(ctx context.Context, req *models.OrderPrimeryKey) (*models.Order, error) {

	var (
		id         sql.NullString
		carId      sql.NullString
		clientId   sql.NullString
		totalPrice sql.NullFloat64
		paidPrice  sql.NullFloat64
		dayCount   sql.NullInt64
		giveKm     sql.NullInt64
		receiveKm  sql.NullInt64
		status     sql.NullString
		createdAt  sql.NullString
		updatedAt  sql.NullString
	)

	query := `
		SELECT
			id,
			car_id,
			client_id,
			total_price,
			paid_price,
			day_count,
			give_km,
			receive_km,
			status,
			created_at,
			updated_at
		FROM "order"
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&carId,
		&clientId,
		&totalPrice,
		&paidPrice,
		&dayCount,
		&giveKm,
		&receiveKm,
		&status,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	resp := &models.Order{
		Id:         id.String,
		CarId:      carId.String,
		ClientId:   clientId.String,
		TotalPrice: totalPrice.Float64,
		PaidPrice:  paidPrice.Float64,
		DayCount:   int(dayCount.Int64),
		GiveKm:     int(giveKm.Int64),
		ReceiveKm:  int(receiveKm.Int64),
		Status:     status.String,
		CreatedAt:  createdAt.String,
		UpdatedAt:  updatedAt.String,
	}

	return resp, err
}

func (r *OrderRepo) GetList(ctx context.Context, req *models.GetListOrderRequest) (*models.GetListOrderResponse, error) {
	var (
		offset = "OFFSET 0"
		limit  = "LIMIT 10"
		resp   = &models.GetListOrderResponse{}
	)

	if req.Offset > 0 {
		offset = fmt.Sprintf("OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf("LIMIT %d", req.Limit)
	}

	query := `
		SELECT
			COUNT(*) OVER(),
			id,
			car_id,
			client_id,
			total_price,
			paid_price,
			day_count,
			give_km,
			receive_km,
			status,
			created_at,
			updated_at
		FROM "order"
	`

	query += offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {

		var (
			id         sql.NullString
			carId      sql.NullString
			clientId   sql.NullString
			totalPrice sql.NullFloat64
			paidPrice  sql.NullFloat64
			dayCount   sql.NullInt64
			giveKm     sql.NullInt64
			receiveKm  sql.NullInt64
			status     sql.NullString
			createdAt  sql.NullString
			updatedAt  sql.NullString
		)

		err = rows.Scan(
			&resp.Count,
			&id,
			&carId,
			&clientId,
			&totalPrice,
			&paidPrice,
			&dayCount,
			&giveKm,
			&receiveKm,
			&status,
			&createdAt,
			&updatedAt,
		)

		resp.Orders = append(resp.Orders, &models.Order{
			Id:         id.String,
			CarId:      carId.String,
			ClientId:   clientId.String,
			TotalPrice: totalPrice.Float64,
			PaidPrice:  paidPrice.Float64,
			DayCount:   int(dayCount.Int64),
			GiveKm:     int(giveKm.Int64),
			ReceiveKm:  int(receiveKm.Int64),
			Status:     status.String,
			CreatedAt:  createdAt.String,
			UpdatedAt:  updatedAt.String,
		})
	}

	return resp, err
}

func (r *OrderRepo) Update(ctx context.Context, order *models.UpdateOrder) error {

	query := `
		UPDATE
			"order"
		SET
			car_id = $2
			client_id = $3
			total_price = $4
			paid_price = $5
			day_count = $6
			give_km = $7
			receive_km = $8
			status = $9
			updated_at = now()
		WHERE id = $1
	`

	_, err := r.db.Exec(ctx, query,
		order.Id,
		order.CarId,
		order.ClientId,
		order.TotalPrice,
		order.PaidPrice,
		order.DayCount,
		order.GiveKm,
		order.ReceiveKm,
		order.Status,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *OrderRepo) UpdateStatus(ctx context.Context, order *models.UpdateStatus) error {

	query := `
		UPDATE 
			"order"
		SET
			status = $2
		WHERE id = $1
	`

	_, err := r.db.Exec(ctx, query, order.Id, order.Status)

	if err != nil {
		return err
	}

	return nil
}

func (r *OrderRepo) Delete(ctx context.Context, req *models.OrderPrimeryKey) error {

	query := `
	delete from "order" where id = $1
	`

	_, err := r.db.Exec(ctx, query, req.Id)
	if err != nil {
		return err
	}

	return nil
}
