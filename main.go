package main

import (
	"fmt"

	"github.com/ahmdrz/goinsta"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

// Specification holds our env data
type Specification struct {
	Username string `required:"true" envconfig:"USERNAME"`
	Password string `required:"true" envconfig:"PASSWORD"`
}

var s Specification

func init() {
	godotenv.Load()
	envconfig.MustProcess("INSTA", &s)
}

func main() {
	insta := goinsta.New(s.Username, s.Password)

	err := insta.Login()
	must(err)

	defer insta.Logout()

	// Change these
	filename := "_example/test.jpg"
	caption := "An example"
	uploadID := insta.NewUploadID()
	quality := 87
	filter := goinsta.Filter_Normat

	res, err := insta.UploadPhoto(filename, caption, uploadID, quality, filter)
	must(err)

	fmt.Printf("Result: %v\n", res.Status)
}
