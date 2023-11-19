package handlers

import (
	"accommodation-booking-system/property-service/pkg/database"
	"accommodation-booking-system/property-service/pkg/models"
	"accommodation-booking-system/property-service/proto"
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type PropertyServiceServer struct {
	proto.UnimplementedPropertyServiceServer
	DBManager     *database.DatabaseManager
	BookingClient proto.BookingServiceClient
}

func convertRoomToProto(room *models.Room) *proto.Room {
	return &proto.Room{
		Id:         room.ID,
		PropertyId: room.PropertyID,
		Type:       room.Type,
		Capacity:   int32(room.Capacity),
		Price:      room.Price,
	}
}

func NewPropertyServiceServer(dbManager *database.DatabaseManager, bookingClient proto.BookingServiceClient) *PropertyServiceServer {
	return &PropertyServiceServer{
		DBManager:     dbManager,
		BookingClient: bookingClient,
	}
}

func filterAvailableRooms(rooms []models.Room, startDate, endDate time.Time, priceMin, priceMax float32, bookingClient proto.BookingServiceClient) []*proto.Room {
	var availableRooms []*proto.Room

	for _, room := range rooms {
		fmt.Println(room)
		if room.Price >= priceMin && room.Price <= priceMax {
			resp, err := bookingClient.CheckRoomAvailability(context.Background(), &proto.CheckRoomAvailabilityRequest{
				RoomId:    room.ID,
				StartDate: timestamppb.New(startDate),
				EndDate:   timestamppb.New(endDate),
			})
			if err == nil && resp.IsAvailable {
				availableRooms = append(availableRooms, convertRoomToProto(&room))
			}
		}
	}

	return availableRooms
}

func (s *PropertyServiceServer) ListProperties(ctx context.Context, req *proto.ListPropertiesRequest) (*proto.ListPropertiesResponse, error) {
	desiredStartDate := req.DesiredStartDate.AsTime()
	desiredEndDate := req.DesiredEndDate.AsTime()

	var properties []models.Property
	result := s.DBManager.DB.Preload("Rooms").Find(&properties)
	if result.Error != nil {
		return nil, result.Error
	}

	var protoProperties []*proto.Property
	for _, p := range properties {
		availableRooms := filterAvailableRooms(p.Rooms, desiredStartDate, desiredEndDate, req.PriceMin, req.PriceMax, s.BookingClient)
		if len(availableRooms) > 0 {
			protoProperty := convertToProtoProperty(&p)
			protoProperty.Rooms = availableRooms
			protoProperties = append(protoProperties, protoProperty)
		}
	}

	return &proto.ListPropertiesResponse{
		Properties: protoProperties,
	}, nil
}

func convertToProtoProperty(property *models.Property) *proto.Property {
	protoRooms := []*proto.Room{}
	for _, r := range property.Rooms {
		protoRooms = append(protoRooms, &proto.Room{
			Id:         r.ID,
			PropertyId: r.PropertyID,
			Type:       r.Type,
		})
	}
	return &proto.Property{
		Id:       uint64(property.ID),
		Name:     property.Name,
		Location: property.Location,
		Rooms:    protoRooms,
	}
}
