package argoscuolanext

import (
	"log"
	"os"
	"testing"
)

func TestLogin(t *testing.T) {
	credentials := Credentials{
		username:   os.Getenv("USERNAME_ARGOSCUOLANEXT"),
		password:   os.Getenv("PASSWORD_ARGOSCUOLANEXT"),
		schoolCode: os.Getenv("SCHOOLCODE_ARGOSCUOLANEXT"),
	}

	_, err := credentials.Login()

	if err != nil {
		log.Fatal(err)
	}
}

func TestMethod(t *testing.T) {
	credentials := Credentials{
		username:   os.Getenv("USERNAME_ARGOSCUOLANEXT"),
		password:   os.Getenv("PASSWORD_ARGOSCUOLANEXT"),
		schoolCode: os.Getenv("SCHOOLCODE_ARGOSCUOLANEXT"),
	}

	session, err := credentials.Login()

	if err != nil {
		log.Fatal(err)
	}

	log.Print(session.Orario())
}
