package session

import (
	"github.com/asmcos/requests"
	"time"
)

type Credentials struct {
	schoolCode string
	username   string
	password   string
}

var restApiUrl = "https://www.portaleargo.it/famiglia/api/rest/"
var argoKey = "ax6542sdru3217t4eesd9"
var argoVersion = "2.0.2"

type Session struct {
	credentials *Credentials
	loggedIn    bool
	auth        map[string]string
	keys        map[string]interface{}
}

func (c *Credentials) Login() (Session, error) {
	req := requests.Requests()

	var resAuth interface{}
	var resKeys interface{}

	session := Session{
		credentials: c,
		auth:        make(map[string]string),
		keys:        make(map[string]interface{}),
	}

	r, err := req.Get(
		restApiUrl+"login",
		requests.Header{
			"Content-Type": "application/json",
			"x-key-app":    argoKey,
			"x-version":    argoVersion,
			"user-agent":   "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.133 Safari/537.36",
			"x-cod-min":    session.credentials.schoolCode,
			"x-user-id":    session.credentials.username,
			"x-pwd":        session.credentials.password,
		},
		requests.Params{
			"_dc": time.Now().Format("20060102150405"),
		},
	)

	if err != nil {
		return Session{}, err
	}

	err = r.Json(&resAuth)

	if err != nil {
		return Session{}, err
	}

	for k, v := range resAuth.(map[string]interface{}) {
		session.auth[string(k)] = v.(string)
	}

	session.loggedIn = true

	r, err = req.Get(
		restApiUrl+"schede",
		requests.Header{
			"Content-Type": "application/json",
			"x-key-app":    argoKey,
			"x-version":    argoVersion,
			"user-agent":   "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.133 Safari/537.36",
			"x-cod-min":    session.credentials.schoolCode,
			"x-auth-token": session.auth["token"],
		},
		requests.Params{
			"_dc": time.Now().Format("20060102150405"),
		},
	)

	if err != nil {
		return Session{credentials: c}, err
	}

	err = r.Json(&resKeys)

	if err != nil {
		return Session{}, err
	}

	session.keys = resKeys.([]interface{})[0].(map[string]interface{})

	return session, nil
}

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
			"x-cod-min":    s.credentials.schoolCode,
			"x-prg-alunno": s.keys["prgAlunno"].(string),
			"x-prg-scheda": s.keys["prgScheda"].(string),
			"x-prg-scuola": s.keys["prgScuola"].(string),
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
