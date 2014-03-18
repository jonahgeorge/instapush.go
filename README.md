# [Instapush](https://instapush.im/) for Golang [![wercker status](https://app.wercker.com/status/2aa99fba574aaf114e73c78b690d68ea/s/ "wercker status")](https://app.wercker.com/project/bykey/2aa99fba574aaf114e73c78b690d68ea)

Unofficial implementation of the Instapush API in Golang.<br/>
Experimental API, subject to change.

## Documentation
http://godoc.org/github.com/jonahgeorge/go-instapush

## Setup

## Examples
``` go
func main() {
  client := instapush.NewClient("USER_TOKEN")
  apps, _ := client.ListApps()

  for _, value := range apps {
      fmt.Printf("title: %s\n", value.Title)
      fmt.Printf("id: %s\n", value.Id)
      fmt.Printf("secret: %s\n", value.Secret)
  }
}
```

