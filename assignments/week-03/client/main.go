package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "../CustomerFeedback"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

var c pb.CustomerFeedbackClient

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c = pb.NewCustomerFeedbackClient(conn)

	for {
		fmt.Println("----Choose Menu----")
		fmt.Println("1 : Add Feedback")
		fmt.Println("2 : Get Feedback By BookingCode")
		fmt.Println("3 : Get Feedback By PassengerID")
		fmt.Println("4 : Delete Feedback")
		var menuID int
		_, err := fmt.Scan(&menuID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("You choosed : ", menuID)
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		var passengerId int32
		var bookingId int32
		var feedback string
		switch menuID {
		case 1:
			fmt.Println("-- Please input PassengerID--")
			fmt.Scan(&passengerId)
			fmt.Println("-- Please input BookingCode--")
			fmt.Scan(&bookingId)
			fmt.Println("-- Please input Feedback--")
			fmt.Scan(&feedback)
			addFeedback(ctx, &pb.PassengerFeedback{
				PassengerId: passengerId,
				BookingId:   bookingId,
				Feedback:    feedback})
		case 2:
			fmt.Println("-- Please input BookingCode--")
			fmt.Scan(&bookingId)
			getFeedbackByBookingCode(ctx, bookingId)
		case 3:
			var passengerID int32
			fmt.Println("-- Please input PassengerID--")
			fmt.Scan(&passengerID)
			getFeedbackByPassagerID(ctx, passengerID)
		case 4:
			var passengerID int32
			fmt.Println("-- Please input PassengerID--")
			fmt.Scan(&passengerID)
			deleteFeedbackByPassagerID(ctx, passengerID)
		}
	}
}
func getFeedbackByBookingCode(ctx context.Context, bookingId int32) {
	r, err := c.GetFeedbackByBookingCode(ctx, &pb.GetFeedbackByBookingCodeRequest{BookingId: bookingId})
	if err != nil {
		log.Printf("could not get feedbacks by booking id %v becasue of %v", bookingId, err)
	} else {
		log.Printf("GetFeedbackByBookingCode response: %s", r.GetPassengerFeedback())
	}
}
func getFeedbackByPassagerID(ctx context.Context, passengerID int32) {
	r, err := c.GetFeedbackByPassengerID(ctx, &pb.GetFeedbackByPassengerIDRequest{PassengerId: passengerID})
	if err != nil {
		log.Printf("could not get feedbacks by passager id %v because of %v", passengerID, err)
	} else {
		log.Printf("GetFeedbackByPassagerID response: %s", r.GetFeedbacks())
	}
}

func deleteFeedbackByPassagerID(ctx context.Context, passengerID int32) {
	r, err := c.DeleteFeedbackByPassengerID(ctx, &pb.DeleteFeedbackByPassengerIDRequest{PassengerId: passengerID})
	if err != nil {
		log.Printf("could not delete feedbacks by passager id %v because of %v", passengerID, err)
	} else {
		log.Printf("DeleteFeedbackByPassagerID response: %s", r.GetMsg())
	}
}

func addFeedback(ctx context.Context, passengerFeedback *pb.PassengerFeedback) {
	r, err := c.AddFeedback(ctx, &pb.AddFeedbackRequest{PassengerFeedback: passengerFeedback})
	if err != nil {
		log.Printf("could not add the feedback by params %v because of %v", &passengerFeedback, err)
	} else {
		log.Printf("AddFeedback response: %s", r.GetMsg())
	}
}
