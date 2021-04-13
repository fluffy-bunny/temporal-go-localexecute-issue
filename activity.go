package app

import (
	"fmt"
)

type BeautifulActivities struct {
}

var MyBeautifulActivities *BeautifulActivities

func NewBeautifulActivities() *BeautifulActivities {
	return &BeautifulActivities{}
}
func (a *BeautifulActivities) ComposeGreeting(name string) (string, error) {
	if a == nil {
		// This happens when
		fmt.Println("Ermaghd me pointer is nil!")
		panic("Ermaghd me pointer is nil!")
	} else {
		fmt.Println("Nothing to see here, everything is A-OK!")
	}
	greeting := fmt.Sprintf("Hello %s!", name)
	return greeting, nil
}
