# go-selenium

# Uses

[DATA-DOG Godog]

[Tebeka Selenium]

[DATA-DOG Godog]: https://github.com/DATA-DOG/godog
[Tebeka Selenium]: https://github.com/tebeka/selenium

# Prerequisites

selenium-server-standalone-3.4.0: (referenced in webdriver.go)

    % wget "http://selenium-release.storage.googleapis.com/3.4/selenium-server-standalone-3.4.0.jar"

geckodriver-v0.18.0-linux64.tar.gz: (referenced in webdriver.go)

    % wget "https://github.com/mozilla/geckodriver/releases/download/v0.18.0/geckodriver-v0.18.0-linux64.tar.gz"
    % tar -xvzf geckodriver-v0.18.0-linux64.tar.gz geckodriver-v0.18.0-linux64
    % chmod +x geckodriver-v0.18.0-linux64

firefox:

    % apt-get install firefox
    
# TODO

Kill the webDriver following test execution
