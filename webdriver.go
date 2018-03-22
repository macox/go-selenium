package main

import (
	"fmt"

	"os"

	"github.com/tebeka/selenium"
)

func getWebDriver() selenium.WebDriver {
	const (
		// These paths will be different on your system.
		seleniumPath    = "vendor/selenium-server-standalone-3.4.0.jar"
		geckoDriverPath = "vendor/geckodriver-v0.18.0-linux64"
		port            = 8080
	)
	opts := []selenium.ServiceOption{
		selenium.StartFrameBuffer(),           // Start an X frame buffer for the browser to run in.
		selenium.GeckoDriver(geckoDriverPath), // Specify the path to GeckoDriver in order to use Firefox.
		selenium.Output(os.Stderr),            // Output debug information to STDERR.
	}
	selenium.SetDebug(true)
	selenium.NewSeleniumService(seleniumPath, port, opts...)
//	service, err := selenium.NewSeleniumService(seleniumPath, port, opts...)
//	if err != nil {
//		panic(err) // panic is used only as an example and is not otherwise recommended.
//	}
//	defer service.Stop()

	// Connect to the WebDriver instance running locally.
	caps := selenium.Capabilities{"browserName": "firefox"}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		panic(err)
	}
//	defer wd.Quit()
	return wd
}

