/*
 * This file is a part of ArgoScuolaNext-Go
 *
 * Copyright (c) 2019 The ArgoScuolaNext-Go Authors (see AUTHORS)
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice
 * shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package argoscuolanext

import (
	"log"
	"os"
	"testing"
	"time"
)

func TestLogin(t *testing.T) {
	credentials := Credentials{
		Username:   os.Getenv("USERNAME_ARGOSCUOLANEXT"),
		Password:   os.Getenv("PASSWORD_ARGOSCUOLANEXT"),
		SchoolCode: os.Getenv("SCHOOLCODE_ARGOSCUOLANEXT"),
	}

	session, err := credentials.Login()

	if err != nil {
		log.Fatal(err)
	}

	log.Print(session)
}

func TestSession_Assenze(t *testing.T) {
	credentials := Credentials{
		Username:   os.Getenv("USERNAME_ARGOSCUOLANEXT"),
		Password:   os.Getenv("PASSWORD_ARGOSCUOLANEXT"),
		SchoolCode: os.Getenv("SCHOOLCODE_ARGOSCUOLANEXT"),
	}

	session, err := credentials.Login()

	if err != nil {
		log.Fatal(err)
	}

	log.Print(session.Assenze())
}

func TestSession_Argomenti(t *testing.T) {
	credentials := Credentials{
		Username:   os.Getenv("USERNAME_ARGOSCUOLANEXT"),
		Password:   os.Getenv("PASSWORD_ARGOSCUOLANEXT"),
		SchoolCode: os.Getenv("SCHOOLCODE_ARGOSCUOLANEXT"),
	}

	session, err := credentials.Login()

	if err != nil {
		log.Fatal(err)
	}

	log.Print(session.Argomenti())
}

func TestSession_Cambiopassword(t *testing.T) {
	credentials := Credentials{
		Username:   os.Getenv("USERNAME_ARGOSCUOLANEXT"),
		Password:   os.Getenv("PASSWORD_ARGOSCUOLANEXT"),
		SchoolCode: os.Getenv("SCHOOLCODE_ARGOSCUOLANEXT"),
	}

	session, err := credentials.Login()

	if err != nil {
		log.Fatal(err)
	}

	log.Print(session.Cambiopassword(os.Getenv("PASSWORD_ARGOSCUOLANEXT")))
}

func TestSession_Compiti(t *testing.T) {
	credentials := Credentials{
		Username:   os.Getenv("USERNAME_ARGOSCUOLANEXT"),
		Password:   os.Getenv("PASSWORD_ARGOSCUOLANEXT"),
		SchoolCode: os.Getenv("SCHOOLCODE_ARGOSCUOLANEXT"),
	}

	session, err := credentials.Login()

	if err != nil {
		log.Fatal(err)
	}

	log.Print(session.Compiti())
}

func TestSession_Docenticlasse(t *testing.T) {
	credentials := Credentials{
		Username:   os.Getenv("USERNAME_ARGOSCUOLANEXT"),
		Password:   os.Getenv("PASSWORD_ARGOSCUOLANEXT"),
		SchoolCode: os.Getenv("SCHOOLCODE_ARGOSCUOLANEXT"),
	}

	session, err := credentials.Login()

	if err != nil {
		log.Fatal(err)
	}

	log.Print(session.Docenticlasse())
}

func TestSettings_GetSession(t *testing.T) {
	credentials := Credentials{
		Username:   os.Getenv("USERNAME_ARGOSCUOLANEXT"),
		Password:   os.Getenv("PASSWORD_ARGOSCUOLANEXT"),
		SchoolCode: os.Getenv("SCHOOLCODE_ARGOSCUOLANEXT"),
	}

	session, err := credentials.Login()

	if err != nil {
		log.Fatal(err)
	}

	sessionFork := session.Settings[0].GetSession()

	log.Print(sessionFork)
}

func TestSession_Oggi(t *testing.T) {
	credentials := Credentials{
		Username:   os.Getenv("USERNAME_ARGOSCUOLANEXT"),
		Password:   os.Getenv("PASSWORD_ARGOSCUOLANEXT"),
		SchoolCode: os.Getenv("SCHOOLCODE_ARGOSCUOLANEXT"),
	}

	session, err := credentials.Login()

	if err != nil {
		log.Fatal(err)
	}

	log.Print(session.Oggi(time.Now()))
}

func TestSession_Orario(t *testing.T) {
	credentials := Credentials{
		Username:   os.Getenv("USERNAME_ARGOSCUOLANEXT"),
		Password:   os.Getenv("PASSWORD_ARGOSCUOLANEXT"),
		SchoolCode: os.Getenv("SCHOOLCODE_ARGOSCUOLANEXT"),
	}

	session, err := credentials.Login()

	if err != nil {
		log.Fatal(err)
	}

	log.Print(session.Orario())
}

func TestSession_Notedisciplinari(t *testing.T) {
	credentials := Credentials{
		Username:   os.Getenv("USERNAME_ARGOSCUOLANEXT"),
		Password:   os.Getenv("PASSWORD_ARGOSCUOLANEXT"),
		SchoolCode: os.Getenv("SCHOOLCODE_ARGOSCUOLANEXT"),
	}

	session, err := credentials.Login()

	if err != nil {
		log.Fatal(err)
	}

	log.Print(session.Notedisciplinari())
}

func TestSession_Promemoria(t *testing.T) {
	credentials := Credentials{
		Username:   os.Getenv("USERNAME_ARGOSCUOLANEXT"),
		Password:   os.Getenv("PASSWORD_ARGOSCUOLANEXT"),
		SchoolCode: os.Getenv("SCHOOLCODE_ARGOSCUOLANEXT"),
	}

	session, err := credentials.Login()

	if err != nil {
		log.Fatal(err)
	}

	log.Print(session.Promemoria())
}

func TestSession_Votigiornalieri(t *testing.T) {
	credentials := Credentials{
		Username:   os.Getenv("USERNAME_ARGOSCUOLANEXT"),
		Password:   os.Getenv("PASSWORD_ARGOSCUOLANEXT"),
		SchoolCode: os.Getenv("SCHOOLCODE_ARGOSCUOLANEXT"),
	}

	session, err := credentials.Login()

	if err != nil {
		log.Fatal(err)
	}

	log.Print(session.Votigiornalieri())
}

func TestSession_Votiscrutinio(t *testing.T) {
	credentials := Credentials{
		Username:   os.Getenv("USERNAME_ARGOSCUOLANEXT"),
		Password:   os.Getenv("PASSWORD_ARGOSCUOLANEXT"),
		SchoolCode: os.Getenv("SCHOOLCODE_ARGOSCUOLANEXT"),
	}

	session, err := credentials.Login()

	if err != nil {
		log.Fatal(err)
	}

	log.Print(session.Votiscrutinio())
}
