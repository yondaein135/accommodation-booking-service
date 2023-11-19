package handlers

import (
	"accommodation-booking-system/payment-service/pkg/database"
	"accommodation-booking-system/payment-service/pkg/models"
	"accommodation-booking-system/payment-service/proto"
	"context"
	"fmt"
	"log"
)

type PaymentServiceServer struct {
	proto.UnimplementedPaymentServiceServer
	DBManager     *database.DatabaseManager
	BookingClient proto.BookingServiceClient
}

func NewPaymentServiceServer(dbManager *database.DatabaseManager, bookingClient proto.BookingServiceClient) *PaymentServiceServer {
	return &PaymentServiceServer{
		DBManager:     dbManager,
		BookingClient: bookingClient,
	}
}

func mockSendData(request *proto.PaymentDetails) bool {
	log.Println("sent payment details to vendor, payment has been processed.")
	return true
}

func (s *PaymentServiceServer) ProcessPayment(ctx context.Context, req *proto.PaymentRequest) (*proto.PaymentResponse, error) {
	paymentDetails := req.PaymentDetails
	success := mockSendData(paymentDetails)
	if success {
		_, err := s.BookingClient.UpdatePaymentStatus(ctx, &proto.UpdatePaymentStatusRequest{
			BookingId: req.BookingId,
			IsPaid:    true,
		})
		if err != nil {
			log.Printf("Failed to update payment status in booking service: %v", err)
		}
	}
	message := "Payment processed successfully"

	payment := models.Payment{
		BookingID: req.BookingId,
		Amount:    req.Amount,
		Success:   success,
	}
	result := s.DBManager.DB.Create(&payment)
	if result.Error != nil {
		return nil, result.Error
	}

	return &proto.PaymentResponse{
		PaymentId: fmt.Sprintf("%d", payment.ID),
		Success:   success,
		Message:   message,
	}, nil
}
