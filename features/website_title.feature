Feature: web page title
  In order to have website
  As a website
  I need to be able to have a title

  Scenario: Website has a title
    Given there is a web page
    When I get the title
    Then the title should be "Google"
