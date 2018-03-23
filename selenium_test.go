package main

import (
	"fmt"

	"github.com/DATA-DOG/godog"
)

var wd = getWebDriver()
var pageTitle string

func thereIsAWebPage() error{

	wd.Get("http://google.com")

	return nil
}

func iGetTheTitle() error {

    actualTitle, err := wd.Title()
    if err != nil {
        panic(err)
    }

    pageTitle = actualTitle
	return nil
}

func theTitleShouldBe(expectedTitle string) error {
	if pageTitle != expectedTitle {
		return fmt.Errorf("expected title to be %d, but is %d", expectedTitle, pageTitle)
	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^there is a web page$`, thereIsAWebPage)
	s.Step(`^I get the title$`, iGetTheTitle)
	s.Step(`^the title should be "([^"]*)"$`, theTitleShouldBe)
}
