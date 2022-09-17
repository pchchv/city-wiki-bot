package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/biter777/countries"
)

var response string

func getEmoji(countryName string) (string, error) {
	var emoji string
	c := countries.ByName(countryName)
	if c.Info().Name == "Unknown" {
		return emoji, errors.New("Incorrect country name")
	}
	emoji = c.Emoji()
	return emoji, nil
}

func getCountryNames(name string) (string, string) {
	var country []string
	file, err := os.Open("_countries.csv")
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = ';'
	for {
		record, err := reader.Read()
		if err != nil && err != io.EOF {
			log.Panic(err)
			break
		}
		if record[1] == name || record[4] == name {
			country = record
			break
		}
	}
	if country == nil {
		log.Panic(errors.New("Incorrect country name"))
	}
	return country[1], country[4]
}

func responseBuilder(country string) error {
	ruCountryName, enCountryName := getCountryNames(country)
	emoji, err := getEmoji(enCountryName)
	if err != nil {
		return err
	}
	ruWikiUrl := "ru.wikipedia.org/wiki/" + ruCountryName
	enWikiUrl := "en.wikipedia.org/wiki/" + enCountryName
	response = fmt.Sprintf("%v\n\n%v\n\n%v\n", emoji, ruWikiUrl, enWikiUrl)
	return nil
}

func main() {
	bot()
}
