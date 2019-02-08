package services

import (
	"context"
	"io"
	"log"
	"time"

	pb "orami.com/techshare/pkg"
)

func PrintPersonGRPC(client pb.TechShareClient, id int64) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	person, err := client.GetPerson(ctx, &pb.PersonRequest{Id: id})
	if err != nil {
		log.Print(err)
	}

	log.Printf("%+v\n", person)
}

func PrintPeopleGRPC(client pb.TechShareClient, limit int64) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	stream, err := client.ListPeople(ctx, &pb.PeopleRequest{Limit: limit})
	if err != nil {
		log.Fatal(err)
	}

	for {
		person, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("%v.ListPeople(_) = _, %v", client, err)
		}

		log.Printf("%+v\n", person)
	}
}
