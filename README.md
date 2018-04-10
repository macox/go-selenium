# go-selenium

# Uses

[DATA-DOG Godog]

[Tebeka Selenium]

[DATA-DOG Godog]: https://github.com/DATA-DOG/godog
[Tebeka Selenium]: https://github.com/tebeka/selenium

# Prerequisites

godog:
 
    %go get github.com/DATA-DOG/godog/cmd/godog

selenium-server-standalone-3.4.0: (referenced in webdriver.go)

    % wget "http://selenium-release.storage.googleapis.com/3.4/selenium-server-standalone-3.4.0.jar" 
    
tebeka/selenium
    
    %go get github.com/tebeka/selenium
    
Firefox    

geckodriver-v0.18.0-linux64.tar.gz: (referenced in webdriver.go)

    % wget "https://github.com/mozilla/geckodriver/releases/download/v0.18.0/geckodriver-v0.18.0-linux64.tar.gz"
    % tar -xvzf geckodriver-v0.18.0-linux64.tar.gz geckodriver-v0.18.0-linux64
    % chmod +x geckodriver-v0.18.0-linux64
    % rm geckodriver-v0.18.0-linux64.tar.gz

firefox:

    % apt-get install firefox
    
xvfb:    
    
    % apt-get install xvfb
    
Chrome

chromeDriver:
    
    % wget https://chromedriver.storage.googleapis.com/2.37/chromedriver_linux64.zip
    % unzip chromedriver_*.zip
    % mv chromedriver /usr/bin/chromedriver
    % chmod a+x /usr/bin/chromedriver
    % rm -rf chromedriver_*.zip
    


    
