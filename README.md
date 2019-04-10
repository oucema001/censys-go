# censys-go
[![Build Status](https://api.travis-ci.org/oucema001/censys-go.svg?branch=master)](https://travis-ci.org/oucema001/censys-go)
[![Build status](https://ci.appveyor.com/api/projects/status/5rd44yg4p5umtq31?svg=true)](https://ci.appveyor.com/project/oucema001/censys-go)
[![codecov](https://codecov.io/gh/oucema001/censys-go/branch/master/graph/badge.svg)](https://codecov.io/gh/oucema001/censys-go)
[![GoDoc](https://godoc.org/github.com/oucema001/censys-go/censys?status.svg)](https://godoc.org/github.com/oucema001/censys-go/censys)
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/oucema001/censys-go)](https://goreportcard.com/report/github.com/oucema001/censys-go)


To start working with censys you have to create an account and get an application Key and application secret First. You can do this at https://censys.io.

### Installation

Download the package:

```bash
go get github.com/oucema001/censys-go
```

You can now use the library in your own projects :-)

### Usage

```go
package main

import (
    "log"
    
    "github.com/oucema001/censys-go"
)

func main() {
    client := censys.NewClient(nil, "MY_APPID", "MY_APPSECRET")
    dns, err := client.GetDNSResolve([]string{"google.com", "ya.ru"})
    
    a, err := client.Search(context.Background(), "www.google.com", censys.WEBSITES)
	if err != nil {
		log.Panic(err)
    } else {
        fmt.Println("alexa rank : %d",a.Results[0].AlexaRank)
    }
}
```
This will output : 

```bash
alexa rank : 1
```

###Implemented Censys APIs : 

#### Account
- [x] /api/v1/account
#### Search
- [x] /api/v1/search/ipv4
- [x] /api/v1/search/websites
- [x] /api/v1/search/certificates
#### View
- /api/v1/view/:index/:id
- [x] /api/v1/view/ipv4/:id
- [x] /api/v1/view/websites/:id
- [x] /api/v1/view/certificates/:id
#### Data
- [ ] /api/v1/data
#### Report
- /api/v1/report/:index
- [x] /api/v1/report/ipv4
- [ ] /api/v1/report/websites
- [ ] /api/v1/report/certificates

### Links
* [Censys API documentation](https://censys.io/api/v1/docs/report)
