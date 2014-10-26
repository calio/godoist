package godoist

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

type Client struct {
	email    string
	password string

	user User
}

func NewClient(email string, password string) (Client, error) {
	if email == "" {
		return Client{}, errors.New("email must not be empty")
	}
	if password == "" {
		return Client{}, errors.New("password must not be empty")
	}

	return Client{
		email:    email,
		password: password,
	}, nil
}

func (c Client) ApiToken() string {
	return c.user.Api_token
}

func (c *Client) Login() error {
	resp, err := HTTPHelper.PostForm("https://api.todoist.com/api/login", url.Values{"email": {c.email}, "password": {c.password}})
	if err != nil {
		return err
	}

	body, err := HTTPHelper.ResponseBodyAsBytes(resp)
	if err != nil {
		return err
	}

	if body == nil {
		return errors.New("Nil response body")
	}

	if len(body) == 0 {
		return errors.New("Empty response body")
	}

	if string(body) == "\"LOGIN_ERROR\"" {
		return errors.New("login failed - incorrect email/password?")
	}

	err = json.Unmarshal(body, &c.user)
	if err != nil {
		return err
	}

	return nil
}

func (c Client) Ping() (string, error) {
	resp, err := HTTPHelper.PostForm("https://api.todoist.com/api/ping", url.Values{"token": {c.ApiToken()}})
	if err != nil {
		return "", err
	}

	switch resp.StatusCode {
	case http.StatusUnauthorized, http.StatusForbidden:
		return "", errors.New("Invalid login")
	}

	body, err := HTTPHelper.ResponseBodyAsBytes(resp)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
