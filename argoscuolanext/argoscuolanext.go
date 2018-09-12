// This file is a part of argoscuolanext-go
//
// Copyright (c) 2018 The argoscuolanext-go Authors (see AUTHORS)
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice
// shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// ArgoScuolaNext APIs for Go.
//
// ArgoScuolaNext APIs for Go creates a platform that can
// be used to check a student's statistics using just his
// Credentials. For example:
//
//     package main
//
//     import (
//         "github.com/hearot/argoscuolanext-go/argoscuolanext"
//         "log"
//     )
//
//     func main() {
//         Credentials = argoscuolanext.Credentials(
//             Username: "USERNAME",
//             Password: "PASSWORD",
//             SchoolCode: "SCHOOLCODE",
//         )
//
//         session, err = Credentials.Login()
//
//         if err != nil {
//             log.Fatal(err)
//         }
//
//         log.Print(session.Assenze())
//     }
//
// See the documentation for more details.
package argoscuolanext

import (
	"encoding/json"
	"errors"
	"github.com/parnurzeal/gorequest"
	"strconv"
	"time"
)

// restApiUrl is the REST API Endpoint.
var restApiUrl = "https://www.portaleargo.it/famiglia/api/rest/"

// argoKey is the application key for the API.
var argoKey = "ax6542sdru3217t4eesd9"

// argoSession is the version of the API.
var argoVersion = "2.0.2"

// Struct used to define all abilitations (what the users can do using
// the APIs).
type Abilitations struct {
	OrarioScolastico             bool `json:"ORARIO_SCOLASTICO"`              // If the student can view his timetable
	ValutazioniPeriodiche        bool `json:"VALUTAZIONI_PERIODICHE"`         // If the student can view his periodic marks
	CompitiAssegnati             bool `json:"COMPITI_ASSEGNATI"`              // If the student can view his homeworks
	TabelloneScrutinioFinale     bool `json:"TABELLONE_SCRUTINIO_FINALE"`     // If the student can view his final marks
	CurriculumVisualizzaFamiglia bool `json:"CURRICULUM_VISUALIZZA_FAMIGLIA"` // If the student can view the curriculum
	ConsiglioDiIstituto          bool `json:"CONSIGLIO_DI_ISTITUTO"`          // If the student can view his school council
	NoteDisciplinari             bool `json:"NOTE_DISCIPLINARI"`              // If the student can view his annotations
	AccessoConControlloScheda    bool `json:"ACCESSO_CON_CONTROLLO_SCHEDA"`   // If the student can access his manage his data
	VotiGiudizi                  bool `json:"VOTI_GIUDIZI"`                   // If the student can view his final marks with opinions
	ValutazioniGiornaliere       bool `json:"VALUTAZIONI_GIORNALIERE"`        // If the student can view his marks
	IgnoraOpzioneVotiDocenti     bool `json:"IGNORA_OPZIONE_VOTI_DOCENTI"`    // If the student can't manage the marks
	ArgomentiLezione             bool `json:"ARGOMENTI_LEZIONE"`              // If the student can view his classes arguments
	ConsiglioDiClasse            bool `json:"CONSIGLIO_DI_CLASSE"`            // If the student can view his class council
	ValutazioniSospesePeriodiche bool `json:"VALUTAZIONI_SOSPESE_PERIODICHE"` // If the student can view the middle marks ("pagellino")
	PinVoti                      bool `json:"PIN_VOTI"`                       // The pin to access the marks
	PagelleOnline                bool `json:"PAGELLE_ONLINE"`                 // If the student can view his school reports online
	RecuperoDebitoInt            bool `json:"RECUPERO_DEBITO_INT"`            // If the student has got the school debt
	RecuperoDebitoSf             bool `json:"RECUPERO_DEBITO_SF"`             // It the student has got the school debt during winter
	PromemoriaClasse             bool `json:"PROMEMORIA_CLASSE"`              // If the student can view his schedules & tasks
	VisualizzaBachecaPubblica    bool `json:"VISUALIZZA_BACHECA_PUBBLICA"`    // If the student can view his showcase
	CurriculumModificaFamiglia   bool `json:"CURRICULUM_MODIFICA_FAMIGLIA"`   // If the student can edit his curriculum
	TabellonePeriodiIntermedi    bool `json:"TABELLONE_PERIODI_INTERMEDI"`    // If the student can view his winter marks ("pagella del primo periodo")
	TasseScolastiche             bool `json:"TASSE_SCOLASTICHE"`              // If the student has to pay school taxes
	DocentiClasse                bool `json:"DOCENTI_CLASSE"`                 // If the student can view his teachers
	VisualizzaAssenzeRegProf     bool `json:"VISUALIZZA_ASSENZE_REG_PROF"`    // If the student can manage absences of his class
	VisualizzaCurriculum         bool `json:"VISUALIZZA_CURRICULUM"`          // If the student can view the curriculum
	AssenzePerData               bool `json:"ASSENZE_PER_DATA"`               // If the student can view his absences filtered by date
	RichiestaCertificati         bool `json:"RICHIESTA_CERTIFICATI"`          // If the student can request the certificates
	AccessoSenzaControllo        bool `json:"ACCESSO_SENZA_CONTROLLO"`        // If the student can access without surveillance
	PrenotazioneAlunni           bool `json:"PRENOTAZIONE_ALUNNI"`            // If the student can book a date with his teachers
	ModificaRecapiti             bool `json:"MODIFICA_RECAPITI"`              // If the student can edit his shipping information
	PagellinoOnline              bool `json:"PAGELLINO_ONLINE"`               // If the student can manage his middle marks ("pagellino")
	MediaPesata                  bool `json:"MEDIA_PESATA"`                   // If the school is using the weighted average
	GiustificazioniAssenze       bool `json:"GIUSTIFICAZIONI_ASSENZE"`        // If the student can justify his absences
}

// Struct that represents an absence done by the student.
type Absence struct {
	CodEvento          string `json:"codEvento"`          // The event code
	NumOra             string `json:"numOra"`             // The absence hours
	DatGiustificazione string `json:"datGiustificazione"` // When the student has to justify the absence
	PrgScuola          string `json:"prgScuola"`          // The student's school ID
	PrgScheda          string `json:"prgScheda"`          // The student's ID
	BinUID             string `json:"binUid"`             // The BinUID
	CodMin             string `json:"codMin"`             // The ministerial code
	DatAssenza         string `json:"datAssenza"`         // The absence date
	NumAnno            string `json:"numAnno"`            // The year
	PrgAlunno          string `json:"prgAlunno"`          // The student's ID in his classroom
	FlgDaGiustificare  string `json:"flgDaGiustificare"`  // Justification flag
	GiustificataDa     string `json:"giustificataDa"`     // Who justified the absence
	DesAssenza         string `json:"desAssenza"`         // Description of the absence
	RegistrataDa       string `json:"registrataDa"`       // Who registered the absence
}

// Struct that represents the response of the "assenze" method.
type Absences struct {
	Dati         []Absence    `json:"dati"`         // The absences done by the student
	Abilitazioni Abilitations `json:"abilitazioni"` // The student's abilitations
}

// Struct used for the Authentication. It contains the type of the user and
// the authentication token.
type Authentication struct {
	Token      string `json:"token"`      // The authentication token
	TipoUtente string `json:"tipoUtente"` // The type of the user
}

// Credentials is the fundamental struct of the
// entire API, it stores the user Credentials.
type Credentials struct {
	SchoolCode string // The School code of your School on ArgoScuolaNext (not the ministerial code!)
	Username   string // Your username on ArgoScuolaNext
	Password   string // Your password on ArgoScuolaNext
}

// Struct used for the "oggi" method. It represents
// a day.
type Day struct {
	Dati          []Event      `json:"dati"`          // An array of events happened during that day
	Abilitazioni  Abilitations `json:"abilitazioni"`  // The student's abilitations
	NuoviElementi int          `json:"nuoviElementi"` // If there are new elements
}

// Representation of an event.
// It could be homeworks, arguments,
// marks and other types of events.
type Event struct {
	Dati struct {
		DatGiorno    string `json:"datGiorno"`    // The day when the event happened
		DesMateria   string `json:"desMateria"`   // The subject
		NumAnno      int    `json:"numAnno"`      // The year
		PrgMateria   string `json:"prgMateria"`   // The subject's ID
		PrgClasse    string `json:"prgClasse"`    // The student's class ID
		DesCompiti   string `json:"desCompiti"`   // The assigned homeworks
		DesArgomento string `json:"desArgomento"` // The arguments of the day
		PrgScuola    string `json:"prgScuola"`    // The student's school ID
		Docente      string `json:"docente"`      // The teacher who registered the event
		CodMin       string `json:"codMin"`       // The ministerial code
	} `json:"dati"` // The useful informations of the event
	Giorno    string `json:"giorno"`    // The day when the event happened or will happen
	NumAnno   int    `json:"numAnno"`   // The year when the event happened or will happen
	PrgAlunno string `json:"prgAlunno"` // The student's ID in his classroom
	PrgScheda string `json:"prgScheda"` // The student's ID
	PrgScuola string `json:"prgScuola"` // The student's school ID
	Tipo      string `json:"tipo"`      // The type of the event
	Titolo    string `json:"titolo"`    // The event title
	Ordine    int    `json:"ordine"`    // The event ID
	CodMin    string `json:"codMin"`    // The ministerial code
}

// Struct used by the Cambiopassword method to
// change the password. It will be converted to JSON.
type PasswordStruct struct {
	OldPassword string `json:"vecchiaPassword"` // The old password
	NewPassword string `json:"nuovaPassword"`   // The new password
}

// Struct that represents the start and the end of the school.
type SchoolTime struct {
	DatInizio string `json:"datInizio"` // The start (format: YYYY-MM-DD)
	DatFine   string `json:"datFine"`   // The end (format: YYYY-MM-DD)
}

// Session represents the current connection
// to the API. It stores the Credentials, Keys and
// tokens.
type Session struct {
	Credentials *Credentials   // An instance of Credentials that stores your credentials
	LoggedIn    bool           // If the user logged in
	Auth        Authentication // The representation of the Authentication tokens
	Settings    []Settings     // An array of informations about the user
}

// Struct that contains all informations about an user.
type Settings struct {
	SchedaSelezionata bool         `json:"schedaSelezionata"` // The chosen student
	DesScuola         string       `json:"desScuola"`         // The student's school
	PrgScuola         int          `json:"prgScuola"`         // The student's school ID
	PrgScheda         int          `json:"prgScheda"`         // The student's ID
	DesSede           string       `json:"desSede"`           // The student's school venue
	AuthToken         string       `json:"authToken"`         // The student's auth token
	Alunno            Student      `json:"alunno"`            // The student
	CodMin            string       `json:"codMin"`            // The ministerial code
	NumAnno           int          `json:"numAnno"`           // The year
	PrgAlunno         int          `json:"prgAlunno"`         // The student's ID in his classroom
	PrgClasse         int          `json:"prgClasse"`         // The student's classroom ID
	DesDenominazione  string       `json:"desDenominazione"`  // The student's denomination
	DesCorso          string       `json:"desCorso"`          // The student's classroom letter (in Italy, all classes have got a letter)
	Abilitazioni      Abilitations `json:"abilitazioni"`      // What the student can do using the APIs
	AnnoScolastico    SchoolTime   `json:"annoScolastico"`    // The representation of the year, start & end dates
}

// Struct that represents the student.
type Student struct {
	DesCf                string `json:"desCf"`                // The student's "codice fiscale"
	DesCognome           string `json:"desCognome"`           // The student's surname
	DesVia               string `json:"desVia"`               // The student's house street
	DesCap               string `json:"desCap"`               // The student's house cap
	DesNome              string `json:"desNome"`              // The student's name
	DesCellulare         string `json:"desCellulare"`         // The student's mobile number
	DesComuneNascita     string `json:"desComuneNascita"`     // The student's birthplace
	FlgSesso             string `json:"flgSesso"`             // The student's gender
	DatNascita           string `json:"datNascita"`           // The student's birth
	DesIndirizzoRecapito string `json:"desIndirizzoRecapito"` // The student's house for shipping
	DesComuneRecapito    string `json:"desComuneRecapito"`    // The student's city for shipping
	DesCapResidenza      string `json:"desCapResidenza"`      // The student's city cap
	DesComuneResidenza   string `json:"desComuneResidenza"`   // The student's city
	DesTelefono          string `json:"desTelefono"`          // The student's house phone
	DesCittadinanza      string `json:"desCittadinanza"`      // The student's citizenship
}

// Struct that represents a Teacher.
type Teacher struct {
	PrgClasse   int    `json:"prgClasse"`   // The student's classroom ID
	PrgAnagrafe int    `json:"prgAnagrafe"` // The teacher's birth ID
	PrgScuola   int    `json:"prgScuola"`   // The student's school's ID
	Materie     string `json:"materie"`     // The teacher's subject
	Docente     struct {
		Email   string `json:"email"`   // The teacher's e-mail
		Nome    string `json:"nome"`    // The teacher's name
		Cognome string `json:"cognome"` // The teacher's surname
	} `json:"docente"` // A struct that represents some useful informations about the Teacher
	CodMin string `json:"codMin"` // The ministerial code
}

// It represents an array of Teacher objects.
type Teachers []Teacher

// Login() is a method of Credentials struct
// that is used to log in to the API. It will
// return a Session instance.
func (c *Credentials) Login() (Session, error) {
	request := gorequest.New()

	session := Session{
		Credentials: c,
	}

	_, bodyResp, errs := request.Get(restApiUrl+"login").
		Set("Content-Type", "application/json").
		Set("x-key-app", argoKey).
		Set("x-version", argoVersion).
		Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.133 Safari/537.36").
		Set("x-cod-min", session.Credentials.SchoolCode).
		Set("x-user-id", session.Credentials.Username).
		Set("x-pwd", session.Credentials.Password).
		Query("_dc=" + time.Now().Format("20060102150405")).
		Query("datGiorno=" + time.Now().Format("2006-01-02")).
		End()

	if len(errs) > 0 {
		return Session{}, errors.New("authentication failed, check your Credentials")
	}

	err := json.Unmarshal([]byte(bodyResp), &session.Auth)

	if err != nil {
		return Session{}, errors.New("authentication failed, check your Credentials")
	}

	session.LoggedIn = true

	_, bodyResp, errs = request.Get(restApiUrl+"schede").
		Set("Content-Type", "application/json").
		Set("x-key-app", argoKey).
		Set("x-version", argoVersion).
		Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.133 Safari/537.36").
		Set("x-cod-min", session.Credentials.SchoolCode).
		Set("x-auth-token", session.Auth.Token).
		Query("_dc=" + time.Now().Format("20060102150405")).
		Query("datGiorno=" + time.Now().Format("2006-01-02")).
		End()

	if len(errs) > 0 {
		return Session{}, errors.New("authentication failed, check your Credentials")
	}

	err = json.Unmarshal([]byte(bodyResp), &session.Settings)

	if err != nil {
		return Session{}, errors.New("authentication failed, check your Credentials")
	}

	return session, nil
}

// Request is the method used by Session struct
// to do a request to the API. It will return
// the JSON.
func (s *Session) request(method string, date time.Time) (string, error) {
	request := gorequest.New()

	_, bodyResp, errs := request.Get(restApiUrl+method).
		Set("Content-Type", "application/json").
		Set("x-key-app", argoKey).
		Set("x-version", argoVersion).
		Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.133 Safari/537.36").
		Set("x-cod-min", s.Credentials.SchoolCode).
		Set("x-auth-token", s.Auth.Token).
		Set("x-prg-alunno", strconv.Itoa(s.Settings[0].PrgAlunno)).
		Set("x-prg-scheda", strconv.Itoa(s.Settings[0].PrgScheda)).
		Set("x-prg-scuola", strconv.Itoa(s.Settings[0].PrgScuola)).
		Query("_dc=" + time.Now().Format("20060102150405")).
		Query("datGiorno=" + date.Format("2006-01-02")).
		End()

	if len(errs) > 0 {
		return "{}", errors.New("authentication failed, check your Credentials")
	}

	return bodyResp, nil
}

// Post request is the method used by Session struct
// to do a request to the API using a JSON body.
// It will return the converted JSON.
func (s *Session) postRequest(method string, body string, date time.Time) (interface{}, error) {
	var res interface{}

	request := gorequest.New()

	_, bodyResp, errs := request.Post(restApiUrl+method).
		Set("Content-Type", "application/json").
		Set("x-key-app", argoKey).
		Set("x-version", argoVersion).
		Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.133 Safari/537.36").
		Set("x-cod-min", s.Credentials.SchoolCode).
		Set("x-auth-token", s.Auth.Token).
		Set("x-prg-alunno", strconv.Itoa(s.Settings[0].PrgAlunno)).
		Set("x-prg-scheda", strconv.Itoa(s.Settings[0].PrgScheda)).
		Set("x-prg-scuola", strconv.Itoa(s.Settings[0].PrgScuola)).
		Query("_dc=" + time.Now().Format("20060102150405")).
		Query("datGiorno=" + date.Format("2006-01-02")).
		Send(body).
		End()

	if len(errs) > 0 {
		return nil, errs[0]
	}

	err := json.Unmarshal([]byte(bodyResp), &res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

// Returns the student's absences.
func (s *Session) Assenze() (Absences, error) {
	absences := Absences{}

	response, err := s.request("assenze", time.Now())

	if err != nil {
		return absences, err
	}

	json.Unmarshal([]byte(response), &absences)

	return absences, nil
}

// Returns what the student's done.
func (s *Session) Argomenti() (interface{}, error) {
	return s.request("argomenti", time.Now())
}

// Returns the student's homeworks.
func (s *Session) Compiti() (interface{}, error) {
	return s.request("compiti", time.Now())
}

// Returns the student's teachers.
func (s *Session) Docenticlasse() (Teachers, error) {
	teachers := Teachers{}

	response, err := s.request("docenticlasse", time.Now())

	if err != nil {
		return teachers, err
	}

	json.Unmarshal([]byte(response), &teachers)

	return teachers, nil
}

// Change the password of the user.
func (s *Session) Cambiopassword(newPassword string) (interface{}, error) {
	m := PasswordStruct{
		OldPassword: s.Credentials.Password,
		NewPassword: newPassword,
	}

	query, err := json.MarshalIndent(m, "", "  ")

	if err != nil {
		return nil, err
	}

	return s.postRequest("cambiopassword", string(query), time.Now())
}

// Returns the student's plan.
//
// You can view what's happening today or on another day just
// by passing a time.Time object as parameter. If you want
// to get statistics about today, pass time.Now().
func (s *Session) Oggi(date time.Time) (Day, error) {
	day := Day{}

	response, err := s.request("oggi", date)

	if err != nil {
		return day, err
	}

	json.Unmarshal([]byte(response), &day)

	return day, nil
}

// Returns the student's timetable.
func (s *Session) Orario() (interface{}, error) {
	return s.request("orario", time.Now())
}

// Returns the student's annotations.
func (s *Session) Notedisciplinari() (interface{}, error) {
	return s.request("notedisciplinari", time.Now())
}

// Returns the student's notes.
func (s *Session) Promemoria() (interface{}, error) {
	return s.request("promemoria", time.Now())
}

// Returns the student's marks.
func (s *Session) Votigiornalieri() (interface{}, error) {
	return s.request("votigiornalieri", time.Now())
}

// Returns the student's final marks.
func (s *Session) Votiscrutinio() (interface{}, error) {
	return s.request("votiscrutinio", time.Now())
}
