# ArgoScuolaNext APIs
[![GoDoc](https://godoc.org/github.com/hearot/ArgoScuolaNext-Go/argoscuolanext?status.svg)](https://godoc.org/github.com/hearot/ArgoScuolaNext-Go/argoscuolanext)

Client that uses ArgoScuolaNext APIs to manage and view your informations on it.

[ArgoScuolaNext APIs in Php](https://github.com/hearot/ArgoScuolaNext)

[ArgoScuolaNext APIs in Python](https://github.com/hearot/ArgoScuolaNext-Python)

[Italian description of the client](README.md)

## Table of Contents
  - [0. Installation](#installation)
  - [1. Import APIs](#import-apis)
  - [2. Log in](#log-in)
    - [What happened today](#what-happened-today)
    - [Absences](#absences)
    - [Disciplinary notes](#disciplinary-notes)
    - [Daily marks](#daily-marks)
    - [Final marks](#final-marks)
    - [Homework](#homework)
    - [Lesson topics](#lesson-topics)
    - [Class reminder](#class-reminder)
    - [Class schedule](#class-schedule)
    - [Teachers](#teachers)
  - [3. Logout](#log-out)
  - [4. Documentation](https://godoc.org/github.com/hearot/ArgoScuolaNext-Go/argoscuolanext)

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
You must use `import` to import all argoscuolanext package.
```go
import "github.com/hearot/argoscuolanext-go/argoscuolanext"
```

## Log in
To log in you have to define your `Credentials` struct and call the `Login()` function on it.
```go
import (
    "github.com/hearot/argoscuolanext-go/argoscuolanext"
    "log"
)

func main() {
    credentials = argoscuolanext.Credentials{
        Username: "YOUR_USERNAME",
        Password: "YOUR_PASSWORD",
        SchoolCode: "YOUR_SCHOOL_CODE",
    }

    _, err = credentials.Login()

    if err != nil {
        log.Fatal(err)
    }
}
```

### What happened today
You can call `oggi` query by using `Session.Oggi()` function. You have to pass a `time.Time` object as parameter, use `time.Now()` for today.

```go
import (
    "github.com/hearot/argoscuolanext-go/argoscuolanext"
    "log"
    "time"
)

func main() {
    credentials = argoscuolanext.Credentials{
        Username: "YOUR_USERNAME",
        Password: "YOUR_PASSWORD",
        SchoolCode: "YOUR_SCHOOL_CODE",
    }

    session, err = credentials.Login()

    if err != nil {
        log.Fatal(err)
    }

    _, err = session.Oggi(time.Now())

    if err != nil {
        log.Fatal(err)
    }
}
```

Example output:
```json
{
   "dati":[
      {
         "dati":{
            "datGiorno":"2017-10-14",
            "desMateria":"MATEMATICA",
            "numAnno":2017,
            "prgMateria":"prgMateria",
            "prgClasse":"prgClasse",
            "desCompiti":"Maths: Study fractions.",
            "prgScuola":"prgScuola",
            "docente":"(Prof. NAME OF YOUR TEACHER)",
            "codMin":"schoolCode"
         },
         "giorno":"2017-10-14",
         "numAnno":2017,
         "prgAlunno":"prgAlunno",
         "prgScheda":"prgScheda",
         "prgScuola":"prgScuola",
         "tipo":"COM",
         "titolo":"Compiti assegnati",
         "ordine":40,
         "codMin":"schoolCode"
      },
      {
         "dati":{
            "datGiorno":"2017-10-14",
            "desMateria":"LINGUA E LETTERATURA ITALIANA",
            "numAnno":2017,
            "prgMateria":"prgMateria",
            "prgClasse":"prgClasse",
            "prgScuola":"prgScuola",
            "desArgomento":"Italian test.",
            "docente":"(Prof. NAME OF YOUR TEACHER)",
            "codMin":"schoolCode"
         },
         "giorno":"2017-10-14",
         "numAnno":2017,
         "prgAlunno":"prgAlunno",
         "prgScheda":"prgScheda",
         "prgScuola":"prgScuola",
         "tipo":"ARG",
         "titolo":"Argomenti lezione",
         "ordine":50,
         "codMin":"schoolCode"
      }
   ],
   "abilitazioni":{
      "ORARIO_SCOLASTICO":true,
      "VALUTAZIONI_PERIODICHE":true,
      "COMPITI_ASSEGNATI":true,
      "TABELLONE_SCRUTINIO_FINALE":true,
      "CURRICULUM_VISUALIZZA_FAMIGLIA":false,
      "CONSIGLIO_DI_ISTITUTO":true,
      "NOTE_DISCIPLINARI":false,
      "ACCESSO_CON_CONTROLLO_SCHEDA":true,
      "VOTI_GIUDIZI":false,
      "VALUTAZIONI_GIORNALIERE":true,
      "IGNORA_OPZIONE_VOTI_DOCENTI":false,
      "ARGOMENTI_LEZIONE":true,
      "CONSIGLIO_DI_CLASSE":false,
      "VALUTAZIONI_SOSPESE_PERIODICHE":false,
      "PIN_VOTI":false,
      "PAGELLE_ONLINE":true,
      "RECUPERO_DEBITO_INT":false,
      "RECUPERO_DEBITO_SF":false,
      "PROMEMORIA_CLASSE":true,
      "VISUALIZZA_BACHECA_PUBBLICA":false,
      "CURRICULUM_MODIFICA_FAMIGLIA":false,
      "TABELLONE_PERIODI_INTERMEDI":false,
      "TASSE_SCOLASTICHE":true,
      "DOCENTI_CLASSE":false,
      "VISUALIZZA_ASSENZE_REG_PROF":true,
      "VISUALIZZA_CURRICULUM":false,
      "ASSENZE_PER_DATA":true,
      "RICHIESTA_CERTIFICATI":false,
      "ACCESSO_SENZA_CONTROLLO":true,
      "PRENOTAZIONE_ALUNNI":false,
      "MODIFICA_RECAPITI":true,
      "PAGELLINO_ONLINE":false,
      "MEDIA_PESATA":false,
      "GIUSTIFICAZIONI_ASSENZE":false
   },
   "nuoviElementi":0
}
```

### Absences
You can call `assenze` query by using `Session.Assenze()` function.
```go
import (
    "github.com/hearot/argoscuolanext-go/argoscuolanext"
    "log"
)

func main() {
    credentials = argoscuolanext.Credentials{
        Username: "YOUR_USERNAME",
        Password: "YOUR_PASSWORD",
        SchoolCode: "YOUR_SCHOOL_CODE",
    }

    session, err = credentials.Login()

    if err != nil {
        log.Fatal(err)
    }

    _, err = session.Assenze()

    if err != nil {
        log.Fatal(err)
    }
}
```

Example output:
```json
{
   "dati":[
      {
         "codEvento":"A",
         "numOra":"",
         "datGiustificazione":"2017-03-27",
         "prgScuola":"prgScuola",
         "prgScheda":"prgScheda",
         "binUid":"binUid",
         "codMin":"schoolCode",
         "datAssenza":"2017-03-25",
         "numAnno":"2016",
         "prgAlunno":"prgAlunno",
         "flgDaGiustificare":"1",
         "giustificataDa":"(Prof. NAME OF YOUR TEACHER)",
         "desAssenza":"",
         "registrataDa":"(Prof. NAME OF YOUR TEACHER)"
      }
   ]
}
```

### Disciplinary notes
You can call `notedisciplinari` query by using `Session.Notedisciplinari()` function.
```go
import (
    "github.com/hearot/argoscuolanext-go/argoscuolanext"
    "log"
)

func main() {
    credentials = argoscuolanext.Credentials{
        Username: "YOUR_USERNAME",
        Password: "YOUR_PASSWORD",
        SchoolCode: "YOUR_SCHOOL_CODE",
    }

    session, err = credentials.Login()

    if err != nil {
        log.Fatal(err)
    }

    _, err = session.Notedisciplinari()

    if err != nil {
        log.Fatal(err)
    }
}
```

Example output:
```json
{
   "dati":[
      {
         "prgAlunno":"prgAlunno",
         "numAnno":"2016",
         "flgVisualizzata":"S",
         "prgAnagrafe":"prgAnagrafe",
         "prgNota":"prgNota",
         "prgScheda":"prgScheda",
         "prgScuola":"prgScuola",
         "desNota":"The student hasn't done the homeworks.",
         "datNota":"2018-10-14",
         "docente":"(Prof. NAME OF YOUR TEACHER)",
         "codMin":"schoolCode"
      }
   ]
}
```

### Daily marks
You can call `votigiornalieri` query by using `Session.Votigiornalieri()` function.
```go
import (
    "github.com/hearot/argoscuolanext-go/argoscuolanext"
    "log"
)

func main() {
    credentials = argoscuolanext.Credentials{
        Username: "YOUR_USERNAME",
        Password: "YOUR_PASSWORD",
        SchoolCode: "YOUR_SCHOOL_CODE",
    }

    session, err = credentials.Login()

    if err != nil {
        log.Fatal(err)
    }

    _, err = session.Votigiornalieri()

    if err != nil {
        log.Fatal(err)
    }
}
```

Example output:
```json
{
   "dati":[
      {
         "datGiorno":"2017-04-19",
         "desMateria":"GEOGRAFIA",
         "prgMateria":"prgMateria",
         "prgScuola":"prgScuola",
         "prgScheda":"prgScheda",
         "codVotoPratico":"N",
         "decValore":"7.5",
         "codMin":"schoolCode",
         "desProva":"",
         "codVoto":"7\u00bd",
         "numAnno":"2016",
         "prgAlunno":"prgAlunno",
         "desCommento":"",
         "docente":"(Prof NAME OF YOUR TEACHER)\n)"
      }
   ]
}
```

### Final marks
You can call `votiscrutinio` query by using `Session.Votiscrutinio()` function.
```go
import (
    "github.com/hearot/argoscuolanext-go/argoscuolanext"
    "log"
)

func main() {
    credentials = argoscuolanext.Credentials{
        Username: "YOUR_USERNAME",
        Password: "YOUR_PASSWORD",
        SchoolCode: "YOUR_SCHOOL_CODE",
    }

    session, err = credentials.Login()

    if err != nil {
        log.Fatal(err)
    }

    _, err = session.Votiscrutinio()

    if err != nil {
        log.Fatal(err)
    }
}
```

Example output:
```json
{
   "dati":[
      {
         "ordineMateria":"2",
         "desMateria":"LINGUA E LET. ITA.",
         "votoOrale":{
            "codVoto":"7"
         },
         "prgMateria":"prgMateria",
         "prgScuola":"prgScuola",
         "prgScheda":"prgScheda",
         "votoUnico":"1",
         "prgPeriodo":"1",
         "assenze":"1",
         "codMin":"schoolCode",
         "suddivisione":"SO",
         "numAnno":"2016",
         "prgAlunno":"prgAlunno",
         "giudizioSintetico":"",
         "prgClasse":"prgClasse"
      }
   ]
}
```

### Homeworks
You can call `compiti` query by using `Session.Compiti()` function.
```go
import (
    "github.com/hearot/argoscuolanext-go/argoscuolanext"
    "log"
)

func main() {
    credentials = argoscuolanext.Credentials{
        Username: "YOUR_USERNAME",
        Password: "YOUR_PASSWORD",
        SchoolCode: "YOUR_SCHOOL_CODE",
    }

    session, err = credentials.Login()

    if err != nil {
        log.Fatal(err)
    }

    _, err = session.Compiti()

    if err != nil {
        log.Fatal(err)
    }
}
```

Example output:
```json
{
   "dati":[
      {
         "datGiorno":"2017-04-22",
         "desMateria":"S.I. BIOLOGIA",
         "numAnno":"2016",
         "prgMateria":"prgMateria",
         "prgClasse":"prgClasse",
         "desCompiti":"Do exercise number 3 at page 31.",
         "prgScuola":"2",
         "docente":"(Prof. NAME OF YOUR TEACHER)",
         "codMin":"schoolCode"
      }
   ]
}
```

### Lesson topics
You can call `argomenti` query by using `Session.Argomenti()` function.
```go
import (
    "github.com/hearot/argoscuolanext-go/argoscuolanext"
    "log"
)

func main() {
    credentials = argoscuolanext.Credentials{
        Username: "YOUR_USERNAME",
        Password: "YOUR_PASSWORD",
        SchoolCode: "YOUR_SCHOOL_CODE",
    }

    session, err = credentials.Login()

    if err != nil {
        log.Fatal(err)
    }

    _, err = session.Argomenti()

    if err != nil {
        log.Fatal(err)
    }
}
```

Example output:
```json
{
   "dati":[
      {
         "datGiorno":"2017-04-22",
         "desMateria":"S.I. BIOLOGIA",
         "numAnno":"2016",
         "prgMateria":"prgMateria",
         "prgClasse":"prgClasse",
         "prgScuola":"prgScuola",
         "desArgomento":"We have watched a video.",
         "docente":"(Prof. NAME OF YOUR TEACHER)",
         "codMin":"schoolCode\n)"
      }
   ]
}g
```

### Class reminder
You can call `promemoria` query by using `Session.Promemoria()` function.
```go
import (
    "github.com/hearot/argoscuolanext-go/argoscuolanext"
    "log"
)

func main() {
    credentials = argoscuolanext.Credentials{
        Username: "YOUR_USERNAME",
        Password: "YOUR_PASSWORD",
        SchoolCode: "YOUR_SCHOOL_CODE",
    }

    session, err = credentials.Login()

    if err != nil {
        log.Fatal(err)
    }

    _, err = session.Promemoriai()

    if err != nil {
        log.Fatal(err)
    }
}
```

Example output:
```json
{
   "dati":[
      {
         "desAnnotazioni":"IT Test",
         "datGiorno":"2017-05-11",
         "numAnno":"2016",
         "prgProgressivo":"prgProgressivo",
         "prgClasse":"prgClasse",
         "prgAnagrafe":"prgAnagrafe",
         "prgScuola":"prgScuola",
         "desMittente":"NAME OF YOUR TEACHER",
         "codMin":"schoolCode\n)"
      }
   ]
}
```

### Class schedule
You can call `orario` query by using `Session.Orario()` function.
```go
import (
    "github.com/hearot/argoscuolanext-go/argoscuolanext"
    "log"
)

func main() {
    credentials = argoscuolanext.Credentials{
        Username: "YOUR_USERNAME",
        Password: "YOUR_PASSWORD",
        SchoolCode: "YOUR_SCHOOL_CODE",
    }

    session, err = credentials.Login()

    if err != nil {
        log.Fatal(err)
    }

    _, err = session.Orario()

    if err != nil {
        log.Fatal(err)
    }
}
```

Example output:
```json
{
   "dati":[
      {
         "numOra":"1",
         "giorno":"Luned\u00ec",
         "prgClasse":"prgClasse",
         "prgScuola":"prgScuola",
         "lezioni":[
            {
               "materia":"DIRITTO ED ECON.",
               "docente":"(Prof. NAME OF YOUR TEACHER)"
            }
         ],
         "numGiorno":"1",
         "codMin":"schoolCode"
      }
   ]
}
```

### Teachers
You can call `docenticlasse` query by using `Session.Docenticlasse()` function.
```go
import (
    "github.com/hearot/argoscuolanext-go/argoscuolanext"
    "log"
)

func main() {
    credentials = argoscuolanext.Credentials{
        Username: "YOUR_USERNAME",
        Password: "YOUR_PASSWORD",
        SchoolCode: "YOUR_SCHOOL_CODE",
    }

    session, err = credentials.Login()

    if err != nil {
        log.Fatal(err)
    }

    _, err = session.Docenticlasse()

    if err != nil {
        log.Fatal(err)
    }
}
```

Example output:
```json
{
   "dati":[
      {
         "prgClasse":"1967",
         "prgAnagrafe":"prgAnagrafe",
         "prgScuola":"prgScuola",
         "materie":"(S.I. BIOLOGIA)",
         "docente":{
            "email":"",
            "nome":"NAME",
            "cognome":"OF YOUR TEACHER"
         },
         "codMin":"schoolCode"
      }
   ]
}
```

## Log out
To log out you have to assign `nil` to the object.
```go
import (
    "github.com/hearot/argoscuolanext-go/argoscuolanext"
    "log"
)

func main() {
    credentials = argoscuolanext.Credentials{
        Username: "YOUR_USERNAME",
        Password: "YOUR_PASSWORD",
        SchoolCode: "YOUR_SCHOOL_CODE",
    }

    session, err = credentials.Login()

    if err != nil {
        log.Fatal(err)
    }

    session = nil
}
```