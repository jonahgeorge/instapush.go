package instapush

import (
	"bytes"
	"encoding/json"
	"errors"
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

// Returns a list of apps associated with the user token
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

// Returns the app that matches the title argument
func (c Client) FindApp(title string) (App, error) {
	apps, err := c.ListApps()
	if err != nil {
		return App{}, err
	}

	for key, _ := range apps {
		if apps[key].Title == title {
			return apps[key], err
		}
	}

	return App{}, errors.New("App '" + title + "' not found.")
}

// Uses a client object to create a new app. Return message contains id and secret.
func (c Client) AddApp(title string) ([]byte, error) {
	resource := endpoint + "apps/add"

	data := map[string]interface{}{
		"title": title,
	}
	d, err := json.Marshal(data)

	client := &http.Client{}
	req, err := http.NewRequest("POST", resource, bytes.NewBuffer(d))
	req.Header.Set("x-instapush-token", c.Token)
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	contents, err := ioutil.ReadAll(res.Body)

	return contents, err
}

// Returns a list of events for the app
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

// Creates a new event linked to this app
func (a App) AddEvent(title string, trackers []string, message string) ([]byte, error) {
	resource := endpoint + "events/add"

	data := map[string]interface{}{
		"title":    title,
		"trackers": trackers,
		"message":  message,
	}
	d, err := json.Marshal(data)

	client := &http.Client{}
	req, err := http.NewRequest("POST", resource, bytes.NewBuffer(d))
	req.Header.Set("x-instapush-appid", a.Id)
	req.Header.Set("x-instapush-appsecret", a.Secret)
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	contents, err := ioutil.ReadAll(res.Body)

	return contents, err
}

// Send a new push notification that matches the string argument. The tracker fills the arguments for the event.
func (a App) Send(event string, trackers interface{}) ([]byte, error) {
	resource := endpoint + "post"

	data := map[string]interface{}{
		"event":    event,
		"trackers": trackers,
	}

	d, err := json.Marshal(data)

	client := &http.Client{}
	req, err := http.NewRequest("POST", resource, bytes.NewBuffer(d))
	req.Header.Set("x-instapush-appid", a.Id)
	req.Header.Set("x-instapush-appsecret", a.Secret)
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	contents, err := ioutil.ReadAll(res.Body)

	return contents, err
}
