package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"encoding/json"
)

type user struct {
	Start_page       string
	Api_token        string
	Time_format      int
	Sort_order       int
	Full_name        string
	Mobile_number    string
	Mobile_host      string
	Timezone         string
	Id               int
	Date_format      int
	Premium_until    string
	Default_reminder string
	Email            string
}

var (
	email    = flag.String("email", "", "Email with which to connect")
	password = flag.String("password", "", "Password with which to connect")

	u user
)

func main() {
	flag.Parse()

	if *email == "" {
		log.Fatal("Email must be provided")
	}

	if *password == "" {
		log.Fatal("Password must be provided")
	}

	fmt.Println("Login response: " + login())

	fmt.Println("Ping response: " + ping())
}

func login() string {
	resp, err := http.PostForm("https://api.todoist.com/api/login", url.Values{"email": {*email}, "password": {*password}})
	if err != nil {
		log.Fatal(fmt.Sprint("Error logging in: %v\n", err))
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	json.Unmarshal(body, &u)
	return string(body)
}

func ping() string {
	resp, err := http.PostForm("https://api.todoist.com/api/ping", url.Values{"token": {u.Api_token}})
	if err != nil {
		log.Fatal(fmt.Sprint("Error pinging: %v\n", err))
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body)
}
