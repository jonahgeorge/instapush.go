# instapush [![wercker status](https://app.wercker.com/status/2aa99fba574aaf114e73c78b690d68ea/s/ "wercker status")](https://app.wercker.com/project/bykey/2aa99fba574aaf114e73c78b690d68ea)

Unofficial implementation of the [Instapush](https://instapush.im/) API in Golang.<br/>
Allows you to send custom Push notifications to your Android/iOS device.<br/>

<br/>
Experimental API, may subject to change.

## Documentation
http://godoc.org/github.com/jonahgeorge/instapush

## Installation
`go get github.com/jonahgeorge/instapush`

## Usage
``` go
package main

import (
  "fmt"
  "github.com/jonahgeorge/instapush"
)

func main() {
  client := instapush.NewClient("USER_TOKEN")
  apps, _ := client.ListApps()

  for _, value := range apps {
      fmt.Printf("title: %s\n", value.Title)
      fmt.Printf("id: %s\n", value.Id)
      fmt.Printf("secret: %s\n", value.Secret)
  }

  trackers := map[string]interface{}{
      "First Name": "James",
      "Last Name":  "Bond",
  }

  // [todo] - better way to select apps
  b, e := apps[0].Send("Sign_Up", trackers)
  if e != nil {
      log.Fatal(e)
  }

  fmt.Printf("%s\n", b)
}
```

## License
Copyright (c) 2014 Jonah George

MIT License

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
"Software"), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
