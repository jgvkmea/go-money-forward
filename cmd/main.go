package main

import (
	"flag"

	"github.com/jgvkmea/go-money-forward/service"
	"github.com/sirupsen/logrus"
)

var (
	email    = flag.String("email", "", "moneyforwarm login email")
	password = flag.String("password", "", "moneyforwarm login password")
)

func main() {
	log := logrus.New()

	flag.Parse()
	if *email == "" || *password == "" {
		log.Errorln("require flag argument email and password")
		return
	}

	if err := service.UpdateBankData(*email, *password); err != nil {
		log.Errorln("failed to update: ", err)
	}
}
