# ArgoScuolaNext API in Go
[![GoDoc](https://godoc.org/github.com/hearot/ArgoScuolaNext-Go/argoscuolanext?status.svg)](https://godoc.org/github.com/hearot/ArgoScuolaNext-Go/argoscuolanext) [![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](./LICENSE) [![Dev @hearot](https://img.shields.io/badge/Dev-%20@hearot-blue.svg)](https://telegram.me/hearot)

Implementation of ArgoScuolaNext APIs, intended to be used to view a student's statistics and information.

[ArgoScuolaNext APIs in Php](https://github.com/hearot/ArgoScuolaNext)

[ArgoScuolaNext APIs in Python](https://github.com/hearot/ArgoScuolaNext-Python)

[Italian description](README.md)

## Table of Contents
  - [0. Installation](#installation)
  - [1. Import APIs](#import-apis)
  - [2. Log in](#log-in)
  - [3. Call a method](#call-a-method)
  - [4. Documentazione](https://godoc.org/github.com/hearot/ArgoScuolaNext-Go/argoscuolanext)

## Installation
You can easily install that package using `go get`:
```bash
go get github.com/hearot/argoscuolanext-go/argoscuolanext
```

Or, if you want to upgrade the package:
```bash
go get -u github.com/hearot/argoscuolanext-go/argoscuolanext
```

## Import APIs
You have to use `import` to import all argoscuolanext package.
```go
package main

import "github.com/hearot/argoscuolanext-go/argoscuolanext"
```

## Log in
To log in you have to define your own Credentials struct and call the Login() function on it.
```go
package main

import (
    "github.com/hearot/argoscuolanext-go/argoscuolanext"
    "log"
)

func main() {
    credentials := argoscuolanext.Credentials{
        Username: "YOUR_USERNAME",
        Password: "YOUR_PASSWORD",
        SchoolCode: "YOUR_SCHOOL_CODE",
    }

    _, err := credentials.Login()

    if err != nil {
        log.Fatal(err)
    }
}
```

### Call a method
Following the [documentation](https://godoc.org/github.com/hearot/ArgoScuolaNext-Go/argoscuolanext), you can call all methods. Here an example:
```go
package main

import (
    "github.com/hearot/argoscuolanext-go/argoscuolanext"
    "log"
    "time"
)

func main() {
    credentials := argoscuolanext.Credentials{
        Username: "YOUR_USERNAME",
        Password: "YOUR_PASSWORD",
        SchoolCode: "YOUR_SCHOOL_CODE",
    }

    session, err := credentials.Login()

    if err != nil {
        log.Fatal(err)
    }

    response, err := session.Assenze()

    if err != nil {
        log.Fatal(err)
    }
    
    log.Print(response)
}
```
