package web_services

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

func ListCountries() {
	url := "http://webservices.oorsprong.org/websamples.countryinfo/CountryInfoService.wso"

	body := `<?xml version="1.0" encoding="utf-8"?>
	<soap12:Envelope xmlns:soap12="http://www.w3.org/2003/05/soap-envelope">
	  <soap12:Body>
		<ListOfCountryNamesByName xmlns="http://www.oorsprong.org/websamples.countryinfo">
		</ListOfCountryNamesByName>
	  </soap12:Body>
	</soap12:Envelope>`

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
