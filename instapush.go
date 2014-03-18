package instapush

import (
	"bytes"
	"encoding/json"
	"fmt"
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

type Tracker struct {
	key   string
	value string
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
	resource := endpoint + "apps/list"

	client := &http.Client{}
	req, err := http.NewRequest("GET", resource, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("x-instapush-token", c.Token)
	req.Header.Set("Content-Type", "application/json")

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

func (c Client) AddApp() ([]byte, error) {
	// todo
	//resource := endpoint + "apps/add"
	return nil, nil
}

func (a App) ListEvents() ([]byte, error) {
	resource := endpoint + "events/list"

	client := &http.Client{}
	req, err := http.NewRequest("GET", resource, nil)
	req.Header.Set("x-instapush-appid", a.Id)
	req.Header.Set("x-instapush-appsecret", a.Secret)
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	contents, err := ioutil.ReadAll(res.Body)

	return contents, err
}

func (a App) AddEvent() ([]byte, error) {
	// todo
	//resource := endpoint + "events/add"
	return nil, nil
}

func (a App) Send(event string, trackers interface{}) ([]byte, error) {
	resource := endpoint + "post"

	data := map[string]interface{}{
		"event":    event,
		"trackers": trackers,
	}

	d, err := json.Marshal(data)
	fmt.Printf("%s\n", d)

	client := &http.Client{}
	req, err := http.NewRequest("POST", resource, bytes.NewBuffer(d))
	req.Header.Set("x-instapush-appid", a.Id)
	req.Header.Set("x-instapush-appsecret", a.Secret)
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	contents, err := ioutil.ReadAll(res.Body)

	return contents, err
}
