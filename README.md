# ArgoScuolaNext API in Go
Programma che utilizza le API di ArgoScuolaNext per gestire e vedere le tue informazioni.

[ArgoScuolaNext APIs in Php](https://github.com/hearot/ArgoScuolaNext)
[ArgoScuolaNext APIs in Python](https://github.com/hearot/ArgoScuolaNext-Python)

## Tabella dei contenuti
  - [0. Installazione](#installazione)
  - [1. Importare le API](#importare-le-api)
  - [2. Log in](#log-in)
    - [Attività della giornata](#attività-della-giornata)
    - [Assenze](#assenze)
    - [Note disciplinari](#note-disciplinari)
    - [Voti giornalieri](#voti-giornalieri)
    - [Voti scrutinio](#voti-scrutinio)
    - [Compiti](#compiti)
    - [Argomenti delle lezioni](#argomenti-delle-lezioni)
    - [Promemoria](#promemoria)
    - [Orario](#orario)
    - [Docenti](#docenti)
  - [3. Logout](#log-out)

## Installazione
Puoi installare facilmente questo client di ArgoScuolaNext utilizzando `go get`:
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
import "github.com/hearot/argoscuolanext-go/argoscuolanext"
```

## Log in
Per utilizzare le API dovrai prima definire lo strutto `Credentials` ed utilizzare il metodo `Login` per accedere.
```go
import (
    "github.com/hearot/argoscuolanext-go/argoscuolanext"
    "log"
)

func main() {
    credentials = argoscuolanext.Credentials{
        Username: "IL_TUO_USERNAME",
        Password: "LA_TUA_PASSWORD",
        SchoolCode: "IL_TUO_CODICE_SCUOLA",
    }

    _, err = credentials.Login()

    if err != nil {
        log.Fatal(err)
    }
}
```

### Attività della giornata
Puoi richiamare la query `oggi` usando il metodo `Session.Oggi()`. Devi obbligatoriamente impostare la data, puoi passare `time.Now()` se vuoi che richiami il metodo per oggi.
```go
import (
    "github.com/hearot/argoscuolanext-go/argoscuolanext"
    "log"
    "time"
)

func main() {
    credentials = argoscuolanext.Credentials{
        Username: "IL_TUO_USERNAME",
        Password: "LA_TUA_PASSWORD",
        SchoolCode: "IL_TUO_CODICE_SCUOLA",
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

Output d'esempio:
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
            "desCompiti":"Matematica: studiare le frazioni.",
            "prgScuola":"prgScuola",
            "docente":"(Prof. NOME DELL'INSEGNANTE)",
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
            "desArgomento":"Verifica d'Italiano.",
            "docente":"(Prof. NOME DELL'INSEGNANTE)",
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

### Assenze
Puoi richiamare la query `assenze` usando la funzione `Session.Assenze()`.
```go
import (
    "github.com/hearot/argoscuolanext-go/argoscuolanext"
    "log"
)

func main() {
    credentials = argoscuolanext.Credentials{
        Username: "IL_TUO_USERNAME",
        Password: "LA_TUA_PASSWORD",
        SchoolCode: "IL_TUO_CODICE_SCUOLA",
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

Output d'esempio:
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
         "giustificataDa":"(Prof. NOME DELL'INSEGNANTE)",
         "desAssenza":"",
         "registrataDa":"(Prof. NOME DELL'INSEGNANTE)"
      }
   ]
}
```

### Note disciplinari
Puoi richiamare la query `notedisciplinari` usando la funzione `Session.Notedisciplinari()`.
```go
import (
    "github.com/hearot/argoscuolanext-go/argoscuolanext"
    "log"
)

func main() {
    credentials = argoscuolanext.Credentials{
        Username: "IL_TUO_USERNAME",
        Password: "LA_TUA_PASSWORD",
        SchoolCode: "IL_TUO_CODICE_SCUOLA",
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

Output d'esempio:
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
         "desNota":"Lo studente non ha fatto i compiti.",
         "datNota":"2018-10-14",
         "docente":"(Prof. NOME DELL'INSEGNANTE)",
         "codMin":"schoolCode"
      }
   ]
}
```

### Voti giornalieri
Puoi richiamare la query `votigiornalieri` usando la funzione `Session.Votigiornalieri()`.
```go
import (
    "github.com/hearot/argoscuolanext-go/argoscuolanext"
    "log"
)

func main() {
    credentials = argoscuolanext.Credentials{
        Username: "IL_TUO_USERNAME",
        Password: "LA_TUA_PASSWORD",
        SchoolCode: "IL_TUO_CODICE_SCUOLA",
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

Output d'esempio:
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
         "docente":"(Prof NOME DELL'INSEGNANTE)\n)"
      }
   ]
}
```

### Voti scrutinio
Puoi richiamare la query `votiscrutinio` usando la funzione `Session.Votiscrutinio()`.
```go
import (
    "github.com/hearot/argoscuolanext-go/argoscuolanext"
    "log"
)

func main() {
    credentials = argoscuolanext.Credentials{
        Username: "IL_TUO_USERNAME",
        Password: "LA_TUA_PASSWORD",
        SchoolCode: "IL_TUO_CODICE_SCUOLA",
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

Output d'esempio:
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

### Compiti
Puoi richiamare la query `compiti` usando la funzione `Session.Compiti()`.
```go
import (
    "github.com/hearot/argoscuolanext-go/argoscuolanext"
    "log"
)

func main() {
    credentials = argoscuolanext.Credentials{
        Username: "IL_TUO_USERNAME",
        Password: "LA_TUA_PASSWORD",
        SchoolCode: "IL_TUO_CODICE_SCUOLA",
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

Output d'esempio:
```json
{
   "dati":[
      {
         "datGiorno":"2017-04-22",
         "desMateria":"S.I. BIOLOGIA",
         "numAnno":"2016",
         "prgMateria":"prgMateria",
         "prgClasse":"prgClasse",
         "desCompiti":"Fare esercizio numero 31 a pagina 2.",
         "prgScuola":"2",
         "docente":"(Prof. NOME DELL'INSEGNANTE)",
         "codMin":"schoolCode"
      }
   ]
}
```

### Argomenti delle lezioni
Puoi richiamare la query `argomenti` usando la funzione `Session.Argomenti()`.
```go
import (
    "github.com/hearot/argoscuolanext-go/argoscuolanext"
    "log"
)

func main() {
    credentials = argoscuolanext.Credentials{
        Username: "IL_TUO_USERNAME",
        Password: "LA_TUA_PASSWORD",
        SchoolCode: "IL_TUO_CODICE_SCUOLA",
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

Output d'esempio:
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
         "desArgomento":"Abbiamo visto un video.",
         "docente":"(Prof. NOME DELL'INSEGNANTE)",
         "codMin":"schoolCode\n)"
      }
   ]
}g
```

### Promemoria
Puoi richiamare la query `promemoria` usando la funzione `Session.Promemoria()`.
```go
import (
    "github.com/hearot/argoscuolanext-go/argoscuolanext"
    "log"
)

func main() {
    credentials = argoscuolanext.Credentials{
        Username: "IL_TUO_USERNAME",
        Password: "LA_TUA_PASSWORD",
        SchoolCode: "IL_TUO_CODICE_SCUOLA",
    }

    session, err = credentials.Login()

    if err != nil {
        log.Fatal(err)
    }

    _, err = session.Promemoria()

    if err != nil {
        log.Fatal(err)
    }
}
```

Output d'esempio:
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
         "desMittente":"NOME DELL'INSEGNANTE",
         "codMin":"schoolCode\n)"
      }
   ]
}
```

### Orario
Puoi richiamare la query `orario` usando la funzione `Session.Orario()`.
```go
import (
    "github.com/hearot/argoscuolanext-go/argoscuolanext"
    "log"
)

func main() {
    credentials = argoscuolanext.Credentials{
        Username: "IL_TUO_USERNAME",
        Password: "LA_TUA_PASSWORD",
        SchoolCode: "IL_TUO_CODICE_SCUOLA",
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

Output d'esempio:
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
               "docente":"(Prof. NOME DELL'INSEGNANTE)"
            }
         ],
         "numGiorno":"1",
         "codMin":"schoolCode"
      }
   ]
}
```

### Docenti
Puoi richiamare la query `docenticlasse` usando la funzione `Session.Docenticlasse()`.
```go
import (
    "github.com/hearot/argoscuolanext-go/argoscuolanext"
    "log"
)

func main() {
    credentials = argoscuolanext.Credentials{
        Username: "IL_TUO_USERNAME",
        Password: "LA_TUA_PASSWORD",
        SchoolCode: "IL_TUO_CODICE_SCUOLA",
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

Output d'esempio:
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
            "nome":"NOME",
            "cognome":"DELL'INSEGNANTE"
         },
         "codMin":"schoolCode"
      }
   ]
}
```

## Log out
Per fare il logout devi assegnare il valore `nil` all'oggetto.
```go
import (
    "github.com/hearot/argoscuolanext-go/argoscuolanext"
    "log"
)

func main() {
    credentials = argoscuolanext.Credentials{
        Username: "IL_TUO_USERNAME",
        Password: "LA_TUA_PASSWORD",
        SchoolCode: "IL_TUO_CODICE_SCUOLA",
    }

    session, err = credentials.Login()

    if err != nil {
        log.Fatal(err)
    }

    session = nil
}
```