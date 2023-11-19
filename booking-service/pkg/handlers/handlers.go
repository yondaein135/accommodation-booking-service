package handlers

import (
	"accommodation-booking-system/booking-service/pkg/database"
	"accommodation-booking-system/booking-service/pkg/models"
	"accommodation-booking-system/booking-service/proto"
	"context"
	"errors"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strings"
)

type BookingServiceServer struct {
	proto.UnimplementedBookingServiceServer
	DBManager *database.DatabaseManager
}

func NewBookingServiceServer(dbManager *database.DatabaseManager) *BookingServiceServer {
	return &BookingServiceServer{DBManager: dbManager}
}

func convertToProtoBooking(booking *models.Booking) *proto.Booking {
	return &proto.Booking{
		Id:          uint64(booking.ID),
		UserId:      booking.UserID,
		RoomId:      booking.RoomID,
		StartDate:   timestamppb.New(booking.StartDate),
		EndDate:     timestamppb.New(booking.EndDate),
		Price:       booking.Price,
		Occupants:   booking.Occupants,
		IsPaid:      booking.IsPaid,
		IsCancelled: booking.IsCancelled,
	}
}

func (s *BookingServiceServer) CheckRoomAvailability(ctx context.Context, req *proto.CheckRoomAvailabilityRequest) (*proto.CheckRoomAvailabilityResponse, error) {
	var bookings []models.Booking

	result := s.DBManager.DB.Where("room_id = ? AND NOT (end_date <= ? OR start_date >= ?) AND is_cancelled = FALSE", req.RoomId, req.StartDate.AsTime(), req.EndDate.AsTime()).Find(&bookings)

	if result.Error != nil {
		return nil, fmt.Errorf("error checking room availability: %v", result.Error)
	}

	isAvailable := len(bookings) == 0
	fmt.Println(isAvailable)
	return &proto.CheckRoomAvailabilityResponse{
		IsAvailable: isAvailable,
	}, nil
}

func (s *BookingServiceServer) CreateBooking(ctx context.Context, req *proto.CreateBookingRequest) (*proto.CreateBookingResponse, error) {
	startDate := req.StartDate.AsTime()
	endDate := req.EndDate.AsTime()

	var existingBookings []models.Booking
	result := s.DBManager.DB.Where("room_id = ? AND NOT (end_date <= ? OR start_date >= ?) AND is_cancelled = FALSE", req.RoomId, startDate, endDate).Find(&existingBookings)
	if result.Error != nil {
		return nil, fmt.Errorf("error checking existing bookings: %v", result.Error)
	}

	if len(existingBookings) > 0 {
		return nil, errors.New("the room is already booked for the selected dates")
	}

	booking := models.Booking{
		UserID:    req.UserId,
		RoomID:    req.RoomId,
		StartDate: startDate,
		EndDate:   endDate,
		Price:     req.Price,
		Occupants: req.Occupants,
	}
	result = s.DBManager.DB.Create(&booking)
	if result.Error != nil {
		return nil, result.Error
	}

	return &proto.CreateBookingResponse{
		Booking: convertToProtoBooking(&booking),
	}, nil
}

func (s *BookingServiceServer) GetBookingHistory(ctx context.Context, req *proto.GetBookingHistoryRequest) (*proto.GetBookingHistoryResponse, error) {
	var conditions []string
	var args []interface{}

	conditions = append(conditions, "user_id = ?")
	args = append(args, req.UserId)

	if req.StartDate != nil {
		conditions = append(conditions, "start_date >= ?")
		args = append(args, req.StartDate.AsTime())
	}
	if req.EndDate != nil {
		conditions = append(conditions, "end_date <= ?")
		args = append(args, req.EndDate.AsTime())
	}
	if req.BookingDate != nil {
		conditions = append(conditions, "created_at <= ?")
		args = append(args, req.BookingDate.AsTime())
	}

	if req.PriceMin > 0 && req.PriceMax >= req.PriceMin {
		conditions = append(conditions, "price >= ?")
		args = append(args, req.PriceMin)
		conditions = append(conditions, "price <= ?")
		args = append(args, req.PriceMax)
	}

	var bookings []models.Booking
	query := s.DBManager.DB.Where(strings.Join(conditions, " AND "), args...)
	result := query.WithContext(ctx).Find(&bookings)
	if result.Error != nil {
		return nil, result.Error
	}
	protoBookings := make([]*proto.Booking, 0, len(bookings))
	for _, b := range bookings {
		protoBookings = append(protoBookings, convertToProtoBooking(&b))
	}
	if len(bookings) > 0 {
		fmt.Println("Booking History:")
		for i, b := range bookings {
			fmt.Printf("\nBooking %d:\n", i+1)
			fmt.Printf("  - Booking ID: %d\n", b.ID)
			fmt.Printf("  - Room ID: %d\n", b.RoomID)
			fmt.Printf("  - Start Date: %s\n", b.StartDate.Format("2006-01-02"))
			fmt.Printf("  - End Date: %s\n", b.EndDate.Format("2006-01-02"))
			fmt.Printf("  - Price: %.2f\n", b.Price)
			fmt.Printf("  - Occupants: %d\n", b.Occupants)
			fmt.Printf("  - Is Paid: %t\n", b.IsPaid)
			fmt.Printf("  - Is Cancelled: %t\n", b.IsCancelled)
		}
	} else {
		fmt.Println("No bookings found.")
	}

	return &proto.GetBookingHistoryResponse{
		Bookings: protoBookings,
	}, nil
}

func (s *BookingServiceServer) CancelBooking(ctx context.Context, req *proto.CancelBookingRequest) (*proto.CancelBookingResponse, error) {
	// We assume that already cancelled booking won't be received by this endpoint
	var booking models.Booking

	result := s.DBManager.DB.First(&booking, req.BookingId)
	if result.Error != nil {
		return nil, result.Error
	}

	if booking.UserID != req.UserId {
		return nil, errors.New("unauthorized: user did not make this booking")
	}

	booking.IsCancelled = true
	if result := s.DBManager.DB.Save(&booking); result.Error != nil {
		return nil, fmt.Errorf("failed to cancel booking: %v", result.Error)
	}

	return &proto.CancelBookingResponse{IsCancelled: true}, nil
}

func (s *BookingServiceServer) UpdatePaymentStatus(ctx context.Context, req *proto.UpdatePaymentStatusRequest) (*proto.UpdatePaymentStatusResponse, error) {
	var booking models.Booking
	result := s.DBManager.DB.First(&booking, req.BookingId)
	if result.Error != nil {
		return nil, result.Error
	}

	booking.IsPaid = req.IsPaid
	result = s.DBManager.DB.Save(&booking)
	if result.Error != nil {
		return nil, result.Error
	}

	return &proto.UpdatePaymentStatusResponse{Success: true}, nil
}
