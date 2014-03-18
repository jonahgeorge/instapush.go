package instapush

import (
	"io/ioutil"
	"net/http"
)

const (
	endpoint = "https://api.instapush.im/v1/post"
)

type Client struct {
	Token string
}

type App struct {
	Title  string `json:"title"`
	Id     string `json:"appId"`
	Secret string `json:"appSecret"`
}

func NewClient(token string) Client {
	var c Client
	c.Token = token
	return c
}

func (c Client) ListApps() ([]byte, error) {
	url := endpoint + "apps/list"

	client := &http.Client{}
	req, err := htt.NewRequest("GET", url, nil)
	req.Header.Set("x-instapush-token", c.Token)

	res, err := client.Do(req)
	contents, err := ioutil.ReadAll(res.Body)

	var apps []App
	err = json.Unmarshal(contents, apps)
	return apps, err
}

func (a App) ListEvents() ([]byte, error) {
	url := endpoint + "events/list"

	client := &http.Client{}
	req, err := htt.NewRequest("GET", url, nil)
	req.Header.Set("x-instapush-appid", a.Id)
	req.Header.Set("x-instapush-appsecret", a.Secret)

	res, err := client.Do(req)
	contents, err := ioutil.ReadAll(res.Body)

	return contents, err
}
