package storage

import (
	"app/models"
	"context"
)

type StorageI interface {
	CloseDB()
	Investor() InvestorRepoI
	Car() CarRepoI
	Client() ClientRepoI
	Order() OrderRepoI
}

type InvestorRepoI interface {
	Insert(context.Context, *models.CreateInvestor) (string, error)
	GetByID(context.Context, *models.InvestorPrimeryKey) (*models.Investor, error)
	GetList(context.Context, *models.GetListInvestorRequest) (*models.GetListInvestorResponse, error)
	GetListFoyda(context.Context, *models.GetListInvestorRequest) (*models.GetListInvestorResponseFoyda, error)
	Update(context.Context, *models.UpdateInvestor) error
	Delete(context.Context, *models.InvestorPrimeryKey) error
}

type CarRepoI interface {
	Insert(context.Context, *models.CreateCar) (string, error)
	GetByID(context.Context, *models.CarPrimeryKey) (*models.Car, error)
	GetList(context.Context, *models.GetListCarRequest) (*models.GetListCarResponse, error)
	Update(context.Context, *models.UpdateCar) error
	Delete(context.Context, *models.CarPrimeryKey) error
}

type ClientRepoI interface {
	Insert(context.Context, *models.CreateClient) (string, error)
	GetByID(context.Context, *models.ClientPrimeryKey) (*models.Client, error)
	GetList(context.Context, *models.GetListClientRequest) (*models.GetListClientResponse, error)
	Update(context.Context, *models.UpdateClient) error
	Delete(context.Context, *models.ClientPrimeryKey) error
}

type OrderRepoI interface {
	Insert(context.Context, *models.CreateOrder) (string, error)
	GetByID(context.Context, *models.OrderPrimeryKey) (*models.Order, error)
	GetList(context.Context, *models.GetListOrderRequest) (*models.GetListOrderResponse, error)
	Update(context.Context, *models.UpdateOrder) error
	UpdateStatus(context.Context, *models.UpdateStatus) error
	Delete(context.Context, *models.OrderPrimeryKey) error
}
