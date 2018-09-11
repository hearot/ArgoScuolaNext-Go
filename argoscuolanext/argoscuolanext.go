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
//         _, err = Credentials.Login()
//
//         if err != nil {
//             log.Fatal(err)
//         }
//     }
//
// See the documentation for Credentials for more details.
package argoscuolanext

import (
	"encoding/json"
	"errors"
	"github.com/asmcos/requests"
	"github.com/parnurzeal/gorequest"
	"strconv"
	"time"
)

// Credentials is the fundamental struct of the
// entire API, it stores the user Credentials.
type Credentials struct {
	SchoolCode string // The School code of your School on ArgoScuolaNext (not the ministerial code!)
	Username   string // Your username on ArgoScuolaNext
	Password   string // Your password on ArgoScuolaNext
}

// restApiUrl is the REST API Endpoint.
var restApiUrl = "https://www.portaleargo.it/famiglia/api/rest/"

// argoKey is the application key for the API.
var argoKey = "ax6542sdru3217t4eesd9"

// argoSession is the version of the API.
var argoVersion = "2.0.2"

// Session represents the current connection
// to the API. It stores the Credentials, Keys and
// tokens.
type Session struct {
	Credentials *Credentials           // An instance of Credentials that stores your credentials
	LoggedIn    bool                   // If the user logged in
	Auth        map[string]string      // A map of tokens used for the Authentication
	Keys        map[string]interface{} // A map of keys used to do the requests
}

// Struct used by the Cambiopassword method to
// change the password. It will be converted to JSON.
type passwordStruct struct {
	OldPassword string `json:"vecchiaPassword"` // The old password
	NewPassword string `json:"nuovaPassword"`   // The new password
}

// Login() is a method of Credentials struct
// that is used to log in to the API. It will
// return a Session instance.
func (c *Credentials) Login() (Session, error) {
	req := requests.Requests()

	var resAuth interface{}
	var resKeys interface{}

	session := Session{
		Credentials: c,
		Auth:        make(map[string]string),
		Keys:        make(map[string]interface{}),
	}

	r, err := req.Get(
		restApiUrl+"login",
		requests.Header{
			"Content-Type": "application/json",
			"x-key-app":    argoKey,
			"x-version":    argoVersion,
			"user-agent":   "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.133 Safari/537.36",
			"x-cod-min":    session.Credentials.SchoolCode,
			"x-user-id":    session.Credentials.Username,
			"x-pwd":        session.Credentials.Password,
		},
		requests.Params{
			"_dc": time.Now().Format("20060102150405"),
		},
	)

	if err != nil {
		return Session{}, errors.New("authentication failed, check your Credentials")
	}

	err = r.Json(&resAuth)

	if err != nil {
		return Session{}, errors.New("authentication failed, check your Credentials")
	}

	for k, v := range resAuth.(map[string]interface{}) {
		session.Auth[string(k)] = v.(string)
	}

	session.LoggedIn = true

	r, err = req.Get(
		restApiUrl+"schede",
		requests.Header{
			"Content-Type": "application/json",
			"x-key-app":    argoKey,
			"x-version":    argoVersion,
			"user-agent":   "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.133 Safari/537.36",
			"x-cod-min":    session.Credentials.SchoolCode,
			"x-auth-token": session.Auth["token"],
		},
		requests.Params{
			"_dc": time.Now().Format("20060102150405"),
		},
	)

	if err != nil {
		return Session{Credentials: c}, err
	}

	err = r.Json(&resKeys)

	if err != nil {
		return Session{}, errors.New("authentication failed, check your Credentials")
	}

	session.Keys = resKeys.([]interface{})[0].(map[string]interface{})

	return session, nil
}

// Request is the method used by Session struct
// to do a request to the API. It will return
// the converted JSON.
func (s *Session) request(method string, date time.Time) (interface{}, error) {
	req := requests.Requests()

	var res interface{}

	r, err := req.Get(
		restApiUrl+method,
		requests.Header{
			"Content-Type": "application/json",
			"x-key-app":    argoKey,
			"x-version":    argoVersion,
			"user-agent":   "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.133 Safari/537.36",
			"x-cod-min":    s.Credentials.SchoolCode,
			"x-auth-token": string(s.Keys["authToken"].(string)),
			"x-prg-alunno": strconv.Itoa(int(s.Keys["prgAlunno"].(float64))),
			"x-prg-scheda": strconv.Itoa(int(s.Keys["prgScheda"].(float64))),
			"x-prg-scuola": strconv.Itoa(int(s.Keys["prgScuola"].(float64))),
		},
		requests.Params{
			"_dc":       time.Now().Format("20060102150405"),
			"datGiorno": date.Format("2006-01-02"),
		},
	)

	if err != nil {
		return nil, err
	}

	err = r.Json(&res)

	if err != nil {
		return nil, err
	}

	return res, nil
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
		Set("x-auth-token", string(s.Keys["authToken"].(string))).
		Set("x-prg-alunno", strconv.Itoa(int(s.Keys["prgAlunno"].(float64)))).
		Set("x-prg-scheda", strconv.Itoa(int(s.Keys["prgScheda"].(float64)))).
		Set("x-prg-scuola", strconv.Itoa(int(s.Keys["prgScuola"].(float64)))).
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
func (s *Session) Assenze() (interface{}, error) {
	return s.request("assenze", time.Now())
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
func (s *Session) Docenticlasse() (interface{}, error) {
	return s.request("docenticlasse", time.Now())
}

// Returns the student's plan.
//
// You can view what's happening today or on another day just
// by passing a time.Time object as parameter. If you want
// to get statistics about today, pass time.Now().
func (s *Session) Oggi(date time.Time) (interface{}, error) {
	return s.request("assenze", date)
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

// Change the password of the user.
func (s *Session) Cambiopassword(newPassword string) (interface{}, error) {
	m := passwordStruct{
		OldPassword: s.Credentials.Password,
		NewPassword: newPassword,
	}

	query, err := json.MarshalIndent(m, "", "  ")

	if err != nil {
		return nil, err
	}

	return s.postRequest("cambiopassword", string(query), time.Now())
}
