package web_services

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

func LanguagesByName() {
	url := "http://webservices.oorsprong.org/websamples.countryinfo/CountryInfoService.wso"

	body := `<?xml version="1.0" encoding="utf-8"?>
	<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
	  <soap:Body>
		<ListOfLanguagesByName xmlns="http://www.oorsprong.org/websamples.countryinfo">
		</ListOfLanguagesByName>
	  </soap:Body>
	</soap:Envelope>`

	resp, err := http.Post(url, "text/xml; charset=utf-8", bytes.NewBuffer([]byte(body)))

	if err != nil {
		log.Fatal("could not consume webservice. ", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)

		bs := string(bodyBytes)
		log.Print(bs)
	}
}
