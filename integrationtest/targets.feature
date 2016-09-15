#noinspection CucumberUndefinedStep
Feature: Targets

  Background:
    Given the application is running
    And I have a valid user

  Scenario: Create a darts game
    Given I make a request POST /darts/api/targets
    """
    {
      "Username": "{username}",
      "TargetName": "T20",
      "Attempts": 20,
      "Successes": 20
    }
    """
    Then I get a 201 response

  Scenario: Missing Username
    Given I make a request POST /darts/api/targets
    """
    {
      "TargetName": "T20",
      "Attempts": 20,
      "Successes": 20
    }
    """
    Then I get a 400 response

  Scenario: Missing TargetName
    Given I make a request POST /darts/api/targets
    """
    {
      "Username": "Username",
      "Attempts": 20,
      "Successes": 20
    }
    """
    Then I get a 400 response

  Scenario: Missing Attempts
    Given I make a request POST /darts/api/targets
    """
    {
      "Username": "{username}",
      "TargetName": "T20",
      "Successes": 20
    }
    """
    Then I get a 400 response

  Scenario: Missing Success
    Given I make a request POST /darts/api/targets
    """
    {
      "Username": "{username}",
      "TargetName": "T20",
      "Attempts": 20
    }
    """
    Then I get a 400 response

  Scenario: Username is not the same as logged in user
    Given I make a request POST /darts/api/targets
    """
    {
      "Username": "madeUpUser",
      "TargetName": "T20",
      "Attempts": 20,
      "Successes": 20
    }
    """
    Then I get a 403 response