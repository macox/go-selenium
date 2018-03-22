package main

import (
	"fmt"

	"github.com/DATA-DOG/godog"
)

var wd = getWebDriver()
var title = ""

func thereIsAWebPage() error{

	wd.Get("http://google.com")


	return nil
}

func iGetTheTitle() error {

	title = wd.Title().Value()
	return nil
}

func theTitleShouldBe(pageTitle string) error {
	fmt.Printf("**********: %s", title)
	if title != pageTitle {
		return fmt.Errorf("expected title to be %d, but is %d", pageTitle, title)
	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^there is a web page$`, thereIsAWebPage)
	s.Step(`^I get the title$`, iGetTheTitle)
	s.Step(`^the title should be "([^"]*)"$`, theTitleShouldBe)
}
