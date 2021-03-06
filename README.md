# instapush 
[![GoDoc](https://godoc.org/github.com/jonahgeorge/instapush?status.png)](https://godoc.org/github.com/jonahgeorge/instapush) 
[![Build Status](https://travis-ci.org/jonahgeorge/instapush.svg)](https://travis-ci.org/jonahgeorge/instapush)

Unofficial implementation of the [Instapush](https://instapush.im/) API v1 in Golang.<br/>
Send custom push notifications to your Android/iOS device.<br/>

## Basic Usage
``` go
import "github.com/jonahgeorge/instapush"

func main() {
  app := instapush.App{Id: "INSTAPUSH_APP_ID", Secret: "INSTAPUSH_APP_SECERT"}

  app.Send("SignUp", map[string]interface{}{
      "First Name": "James",
      "Last Name":  "Bond",
  })
}
```

### Creating an application
``` go
  client := instapush.Client{Token: "INSTAPUSH_USER_TOKEN"}

  // returns the json response with title, id, secret, and status
  // later iteration will return an app
  res, err := client.AddApp("NEW_APP_NAME")
```

### Listing applications
``` go
  client := instapush.Client{Token: "INSTAPUSH_USER_TOKEN"}

  // returns a slice of apps
  apps, err := client.ListApps()
```

### Retrieving an application
``` go
  client := instapush.Client{Token: "INSTAPUSH_USER_TOKEN"}

  // returns an app
  app, err := client.FindApp("INSTAPUSH_APP_TITLE")
```

### Creating an event
``` go
  app := instapush.App{Id: "INSTAPUSH_APP_ID", Secret: "INSTAPUSH_APP_SECERT"}

  res, err := app.AddEvent("SignUp", 
    []string{"First Name", "Last Name", "{First Name} {Last Name} signed up for your app!"})

  fmt.Printf("%s\n", res)
```

### Listing events
``` go
  app := instapush.App{Id: "INSTAPUSH_APP_ID", Secret: "INSTAPUSH_APP_SECERT"}

  res, err := app.ListEvents()

  fmt.Printf("%s\n", res)
```

### Sending a notification
``` go
  app := instapush.App{Id: "INSTAPUSH_APP_ID", Secret: "INSTAPUSH_APP_SECERT"}

  res, err := app.Send("SignUp", map[string]interface{}{
      "First Name": "James",
      "Last Name":  "Bond",
  })

  fmt.Printf("%s\n", res)
```

## License
>Copyright (c) 2014 [Jonah George](http://jonahgeorge.com)

>MIT License

>Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

>The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

>THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
