package main

import (
	"log"
	"os"

	pb "github.com/adhurjaty/hager-microservices/user-service/proto/user"
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"golang.org/x/net/context"
)

func main() {

	cmd.Init()

	// Create new greeter client
	client := pb.NewUserServiceClient("go.micro.srv.user", microclient.DefaultClient)

	username := "adhurjaty"
	email := "adhurjaty@novacoast.com"
	password := "test123"
	company := "Novacoast"
	firstName := "Anil"
	lastName := "Dhurjaty"

	log.Println(username, email, password)

	r, err := client.Create(context.TODO(), &pb.User{
		Username:  username,
		Email:     email,
		Password:  password,
		Company:   company,
		FirstName: firstName,
		LastName:  lastName,
	})
	if err != nil {
		log.Fatalf("Could not create: %v", err)
	}
	log.Printf("Created: %s", r.User.Id)

	getAll, err := client.GetAll(context.Background(), &pb.Request{})
	if err != nil {
		log.Fatalf("Could not list users: %v", err)
	}
	for _, v := range getAll.Users {
		log.Println(v)
	}

	authResponse, err := client.Auth(context.TODO(), &pb.User{
		Username: username,
		Password: password,
	})

	if err != nil {
		log.Fatalf("Could not authenticate user: %s error: %v\n", username, err)
	}

	log.Printf("Your access token is: %s \n", authResponse.Token)

	// let's just exit because
	os.Exit(0)
}
