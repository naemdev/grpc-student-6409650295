package main

import (
	"context"
	"log"
	"time"

	pb "grpc-student/studentpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewStudentServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// GetStudent
	res, err := client.GetStudent(ctx, &pb.StudentRequest{
		Id: 101,
	})
	if err != nil {
		log.Fatalf("Error calling GetStudent: %v", err)
	}

	log.Println("Student Info:")
	log.Printf("ID: %d", res.Id)
	log.Printf("Name: %s", res.Name)
	log.Printf("Major: %s", res.Major)
	log.Printf("Email: %s", res.Email)
	log.Printf("Faculty: %s", res.Faculty)
	log.Printf("Year: %d", res.Year)
	log.Printf("Phone: %s", res.Phone)

	// ListStudents
	resList, err := client.ListStudents(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("Error calling ListStudents: %v", err)
	}

	log.Println("Student List:")
	for _, s := range resList.Student {
		log.Printf("ID: %d", s.Id)
		log.Printf("Name: %s", s.Name)
		log.Printf("Major: %s", s.Major)
		log.Printf("Email: %s", s.Email)
		log.Printf("Faculty: %s", s.Faculty)
		log.Printf("Year: %d", s.Year)
		log.Printf("Phone: %s", s.Phone)
		log.Println("-----")
	}
}
