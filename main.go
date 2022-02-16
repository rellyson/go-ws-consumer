package main

import (
	"log"

	"github.com/rellyson/go-ws-consumer/app/web_services"
)

func main() {

	log.Println("calling countries ws...")
	web_services.ListCountries()

	log.Println("calling currencies ws...")
	web_services.ListCurrenciesByName()

	log.Println("calling languages ws...")
	web_services.LanguagesByName()
}
