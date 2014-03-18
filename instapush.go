package instapush

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	endpoint = "https://api.instapush.im/v1/"
)

type Client struct {
	Token string
}

type App struct {
	Title  string `json:"title"`
	Id     string `json:"appID"`
	Secret string `json:"appSecret"`
}

type Response struct {
	Message string `json:"msg"`
	Error   bool   `json:"error"`
	Status  int    `json:"status"`
}

func NewClient(token string) Client {
	var c Client
	c.Token = token
	return c
}

func (c Client) ListApps() ([]App, error) {
	url := endpoint + "apps/list"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("x-instapush-token", c.Token)

	res, err := client.Do(req)
	contents, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var apps []App
	err = json.Unmarshal(contents, &apps)
	if err != nil {
		log.Fatal(err)
	}
	return apps, err
}

func (a App) ListEvents() ([]byte, error) {
	url := endpoint + "events/list"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("x-instapush-appid", a.Id)
	req.Header.Set("x-instapush-appsecret", a.Secret)

	res, err := client.Do(req)
	contents, err := ioutil.ReadAll(res.Body)

	return contents, err
}
