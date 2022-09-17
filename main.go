package main

import (
	"errors"

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

func main() {
}
