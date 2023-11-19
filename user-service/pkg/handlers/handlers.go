package handlers

import (
	"accommodation-booking-system/user-service/pkg/database"
	"accommodation-booking-system/user-service/pkg/models"
	"accommodation-booking-system/user-service/proto"
	"context"
	"fmt"
)

type UserServiceServer struct {
	proto.UnimplementedUserServiceServer
	DBManager *database.DatabaseManager
}

func NewUserServiceServer(dbManager *database.DatabaseManager) *UserServiceServer {
	return &UserServiceServer{DBManager: dbManager}
}

func (s *UserServiceServer) CreateUser(ctx context.Context, req *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
	user := models.User{
		Name:  req.Name,
		Email: req.Email,
	}
	result := s.DBManager.DB.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &proto.CreateUserResponse{
		User: &proto.User{
			Id:    uint64(user.ID),
			Name:  user.Name,
			Email: user.Email,
		},
	}, nil
}

func (s *UserServiceServer) GetUser(ctx context.Context, req *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	var user models.User
	result := s.DBManager.DB.First(&user, "id = ?", req.UserId)
	if result.Error != nil {
		return nil, result.Error
	}

	return &proto.GetUserResponse{
		User: &proto.User{
			Id:    uint64(user.ID),
			Name:  user.Name,
			Email: user.Email,
		},
	}, nil
}

func (s *UserServiceServer) UpdateBillingAddress(ctx context.Context, req *proto.UpdateBillingAddressRequest) (*proto.UpdateBillingAddressResponse, error) {
	newAddress := models.BillingAddress{
		UserID:       uint(req.UserId),
		AddressLine1: req.BillingAddress.AddressLine1,
		AddressLine2: req.BillingAddress.AddressLine2,
		City:         req.BillingAddress.City,
		State:        req.BillingAddress.State,
		Country:      req.BillingAddress.Country,
		PostalCode:   req.BillingAddress.PostalCode,
	}

	result := s.DBManager.DB.Create(&newAddress)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to add billing address: %v", result.Error)
	}

	return &proto.UpdateBillingAddressResponse{Success: true}, nil
}
