package main

import (
	"context"
	"log"
	"net"

	pb "../CustomerFeedback"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

const (
	port = ":50051"
)

type server struct{}

var database []pb.PassengerFeedback = []pb.PassengerFeedback{
	pb.PassengerFeedback{
		BookingId:   1,
		PassengerId: 1,
		Feedback:    "Good",
	},
	pb.PassengerFeedback{
		BookingId:   2,
		PassengerId: 2,
		Feedback:    "Normal",
	},
	pb.PassengerFeedback{
		BookingId:   3,
		PassengerId: 2,
		Feedback:    "Bad",
	},
}

func ifNotExistedBookingId(bookingId int32) bool {
	for _, fb := range database {
		if fb.BookingId == bookingId {
			return false
		}
	}
	return true
}

func addNewFeedback(newFeedback pb.PassengerFeedback) {
	database = append(database, newFeedback)
}

func deleteFeedbackByPassengerId(passengerId int32) {
	var filteredDatabase []pb.PassengerFeedback
	for _, fb := range database {
		if fb.PassengerId != passengerId {
			filteredDatabase = append(filteredDatabase, fb)
		}
	}
	database = filteredDatabase
}

func getFeedbackByPassengerId(passengerId int32) []*pb.PassengerFeedback {
	listPassengerFeedback := []*pb.PassengerFeedback{}
	for _, fb := range database {
		if fb.PassengerId == passengerId {
			listPassengerFeedback = append(listPassengerFeedback, &fb)
		}
	}
	return listPassengerFeedback
}

func getFeedbackByBookingId(bookingId int32) *pb.PassengerFeedback {
	for _, fb := range database {
		if fb.BookingId == bookingId {
			return &fb
		}
	}
	return nil
}

func (s *server) AddFeedback(ctx context.Context, in *pb.AddFeedbackRequest) (*pb.AddFeedbackResponse, error) {
	log.Printf("AddFeedback request: %v", in)
	newPassangerFeedback := pb.PassengerFeedback{
		Feedback:    in.PassengerFeedback.GetFeedback(),
		BookingId:   in.PassengerFeedback.GetBookingId(),
		PassengerId: in.PassengerFeedback.GetPassengerId(),
	}
	if ifNotExistedBookingId(newPassangerFeedback.BookingId) {
		addNewFeedback(newPassangerFeedback)
		log.Println("Added the feedback successfully")
		return &pb.AddFeedbackResponse{Msg: "Added the feedback"}, nil
	}
	log.Println("Cannot add the feedback")
	return &pb.AddFeedbackResponse{Msg: "Cannot add the feedback"}, status.Error(400, "The feedback are already existed")
}

func (s *server) GetFeedbackByPassengerID(ctx context.Context, in *pb.GetFeedbackByPassengerIDRequest) (*pb.GetFeedbackByPassengerIDResponse, error) {
	log.Printf("GetFeedbackByPassengerID request: %v", in)
	listPassengerFeedback := getFeedbackByPassengerId(in.PassengerId)
	log.Println("Get the feedback by passengerId successfully")
	return &pb.GetFeedbackByPassengerIDResponse{Feedbacks: listPassengerFeedback}, nil
}

func (s *server) GetFeedbackByBookingCode(ctx context.Context, in *pb.GetFeedbackByBookingCodeRequest) (*pb.GetFeedbackByBookingCodeResponse, error) {
	log.Printf("GetFeedbackByBookingID request: %v", in)
	passengerFeedback := getFeedbackByBookingId(in.BookingId)
	if passengerFeedback != nil {
		log.Println("Get the feedback by bookingId successfully")
		return &pb.GetFeedbackByBookingCodeResponse{PassengerFeedback: passengerFeedback}, nil
	}
	log.Println("Cannot find any feedback with the booking id")
	return &pb.GetFeedbackByBookingCodeResponse{PassengerFeedback: nil}, status.Error(400, "Cannot find any feedback with the booking id")
}

func (s *server) DeleteFeedbackByPassengerID(ctx context.Context, in *pb.DeleteFeedbackByPassengerIDRequest) (*pb.DeleteFeedbackByPassengerIDResponse, error) {
	log.Printf("DeleteFeedbackByPassengerID request: %v", in)
	deleteFeedbackByPassengerId(in.PassengerId)
	log.Println("Deleted the feedback successfully")
	return &pb.DeleteFeedbackByPassengerIDResponse{Msg: "Deleted"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCustomerFeedbackServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
