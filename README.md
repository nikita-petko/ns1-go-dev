[![Build Status](https://travis-ci.org/ns1/ns1-go.svg?branch=v2)](https://travis-ci.org/ns1/ns1-go) [![GoDoc](https://godoc.org/gopkg.in/nkpetko/ns1-go.v2/ns1-go.v2?status.svg)](https://godoc.org/gopkg.in/nkpetko/ns1-go.v2/ns1-go.v2)

# NS1 Golang SDK

> This project is in [active development](https://github.com/ns1/community/blob/master/project_status/ACTIVE_DEVELOPMENT.md).

The golang client for the NS1 API: https://ns1.com/api/

# Installing

```
$ go get gopkg.in/nkpetko/ns1-go.v2/ns1-go.v2
```

Examples
========

[See more](https://github.com/ns1/ns1-go/tree/v2/rest/_examples)


```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	api "gopkg.in/nkpetko/ns1-go.v2/ns1-go.v2/rest"
)

func main() {
	k := os.Getenv("NS1_APIKEY")
	if k == "" {
		fmt.Println("NS1_APIKEY environment variable is not set, giving up")
		os.Exit(1)
	}

	httpClient := &http.Client{Timeout: time.Second * 10}
	client := api.NewClient(httpClient, api.SetAPIKey(k))

	zones, _, err := client.Zones.List()
	if err != nil {
		log.Fatal(err)
	}

	for _, z := range zones {
		fmt.Println(z.Zone)
	}

}
```

Contributing
============
Pull Requests and issues are welcome. See the [NS1 Contribution Guidelines](https://github.com/ns1/community) for more information.

Run tests:

```
make test
```

Local dev: use `go mod replace` in client code to point to local checkout of
this repository.

Consider running `./script/install-git-hooks` to install local git hooks for this
project.

# LICENSE

Apache2 - see the included LICENSE file for more information
