package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/robdimsdale/godoist"
	"github.com/robdimsdale/godoist/httphelper"
)

var (
	email    = flag.String("email", "", "Email with which to connect")
	password = flag.String("password", "", "Password with which to connect")

	u godoist.User
)

func main() {
	flag.Parse()

	godoist.HTTPHelper = &httphelper.ActualHTTPHelper{}

	c, err := godoist.NewClient(*email, *password)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error making client: %v\n", err))
	}

	err = c.Login()
	if err != nil {
		log.Fatal(fmt.Sprintf("Error logging in: %v\n", err))
	}

	result, err := c.Ping()
	if err != nil {
		log.Fatal(fmt.Sprintf("Error pinging: %v\n", err))
	}

	fmt.Println("Ping response: " + result)
}
