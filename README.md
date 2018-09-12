# ArgoScuolaNext API in Go
[![GoDoc](https://godoc.org/github.com/hearot/ArgoScuolaNext-Go/argoscuolanext?status.svg)](https://godoc.org/github.com/hearot/ArgoScuolaNext-Go/argoscuolanext) [![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](./LICENSE) [![Dev @hearot](https://img.shields.io/badge/Dev-%20@hearot-blue.svg)](https://telegram.me/hearot)
                                                                                                                                                                  
Implementazione delle API di ArgoScuolaNext al fine d'essere utilizzate per visualizzare le statistiche e le informazioni di uno studente.

[ArgoScuolaNext APIs in Php](https://github.com/hearot/ArgoScuolaNext)

[ArgoScuolaNext APIs in Python](https://github.com/hearot/ArgoScuolaNext-Python)

[English description of the client](README-en.md)

## Tabella dei contenuti
  - [0. Installazione](#installazione)
  - [1. Importare le API](#importare-le-api)
  - [2. Log in](#log-in)
  - [3. Richiamare un metodo](#richiamare-un-metodo)
  - [4. Documentazione](https://godoc.org/github.com/hearot/ArgoScuolaNext-Go/argoscuolanext)

## Installazione
Puoi installare facilmente questo modulo utilizzando `go get`:
```bash
go get github.com/hearot/argoscuolanext-go/argoscuolanext
```

O, se vuoi aggiornare il modulo:
```bash
go get -u github.com/hearot/argoscuolanext-go/argoscuolanext
```

## Importare le API
Devi usare `import` per importare tutto il pacchetto.
```go
package main

import "github.com/hearot/argoscuolanext-go/argoscuolanext"
```

## Log in
Per utilizzare le API dovrai prima utilizzare lo strutto `Credentials` ed il metodo `Login` per accedere.
```go
package main

import (
    "github.com/hearot/argoscuolanext-go/argoscuolanext"
    "log"
)

func main() {
    credentials := argoscuolanext.Credentials{
        Username: "IL_TUO_USERNAME",
        Password: "LA_TUA_PASSWORD",
        SchoolCode: "IL_TUO_CODICE_SCUOLA",
    }

    _, err := credentials.Login()

    if err != nil {
        log.Fatal(err)
    }
}
```

### Richiamare un metodo
Seguendo la [documentazione](https://godoc.org/github.com/hearot/ArgoScuolaNext-Go/argoscuolanext), puoi richiamare qualsiasi metodo. Di seguito un esempio:
```go
package main

import (
    "github.com/hearot/argoscuolanext-go/argoscuolanext"
    "log"
    "time"
)

func main() {
    credentials := argoscuolanext.Credentials{
        Username: "IL_TUO_USERNAME",
        Password: "LA_TUA_PASSWORD",
        SchoolCode: "IL_TUO_CODICE_SCUOLA",
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
