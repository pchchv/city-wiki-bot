package main

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
	"os"

	"github.com/biter777/countries"
)

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
		if err != nil {
			if err == io.EOF {
				log.Println("_______EOF_______")
				break
			}
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

func main() {
}
