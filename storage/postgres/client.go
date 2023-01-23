package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"app/models"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/google/uuid"
)

type ClientRepo struct {
	db *pgxpool.Pool
}

func NewClientRepo(db *pgxpool.Pool) *ClientRepo {
	return &ClientRepo{
		db: db,
	}
}

func (r *ClientRepo) Insert(ctx context.Context, client *models.CreateClient) (string, error) {

	var (
		id = uuid.New().String()
	)

	query := `
		INSERT INTO client (
			id,
			first_name,
			last_name,
			address,
			phone_number,
			updated_at
		) VALUES ($1, $2, $3, $4, $5, now())
	`

	_, err := r.db.Exec(ctx, query,
		id,
		client.FirstName,
		client.LastName,
		client.Address,
		client.PhoneNumber,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *ClientRepo) GetByID(ctx context.Context, req *models.ClientPrimeryKey) (*models.Client, error) {

	var (
		id          sql.NullString
		firstName   sql.NullString
		lastName    sql.NullString
		address     sql.NullString
		phoneNumber sql.NullString
		createdAt   sql.NullString
		updatedAt   sql.NullString
	)

	query := `
		SELECT
			id,
			first_name,
			last_name,
			address,
			phone_number,
			created_at,
			updated_at
		FROM client
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&firstName,
		&lastName,
		&address,
		&phoneNumber,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	resp := &models.Client{
		Id:          id.String,
		FirstName:   firstName.String,
		LastName:    lastName.String,
		Address:     address.String,
		PhoneNumber: phoneNumber.String,
		CreatedAt:   createdAt.String,
		UpdatedAt:   updatedAt.String,
	}

	return resp, err
}

func (r *ClientRepo) GetList(ctx context.Context, req *models.GetListClientRequest) (*models.GetListClientResponse, error) {
	var (
		offset = "OFFSET 0"
		limit  = "LIMIT 10"
		resp   = &models.GetListClientResponse{}
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
			first_name,
			last_name,
			address,
			phone_number,
			created_at,
			updated_at
		FROM client
	`

	query += offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {

		var (
			id          sql.NullString
			firstName   sql.NullString
			lastName    sql.NullString
			address     sql.NullString
			phoneNumber sql.NullString
			createdAt   sql.NullString
			updatedAt   sql.NullString
		)

		err = rows.Scan(
			&resp.Count,
			&id,
			&firstName,
			&lastName,
			&address,
			&phoneNumber,
			&createdAt,
			&updatedAt,
		)

		resp.Clients = append(resp.Clients, &models.Client{
			Id:          id.String,
			FirstName:   firstName.String,
			LastName:    lastName.String,
			Address:     address.String,
			PhoneNumber: phoneNumber.String,
			CreatedAt:   createdAt.String,
			UpdatedAt:   updatedAt.String,
		})
	}

	return resp, err
}

func (r *ClientRepo) Update(ctx context.Context, client *models.UpdateClient) error {
	query := `
		UPDATE
			client
		SET
			first_name = $2,
			last_name = $3,
			address = $4,
			phone_number = $5,
,			updated_at = now()
		WHERE id = $1
	`

	_, err := r.db.Exec(ctx, query,
		client.Id,
		client.FirstName,
		client.LastName,
		client.Address,
		client.PhoneNumber,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *ClientRepo) Delete(ctx context.Context, req *models.ClientPrimeryKey) error {

	_, err := r.db.Exec(ctx, "delete from client where id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}
