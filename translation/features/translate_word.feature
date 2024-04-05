Feature: Translate Word

  Scenario: Translate specific word "hello" to supported languages
    Given the word "hello"
    When I translate the word to "english"
    Then the translation should be "hello"
    When I translate the word to "german"
    Then the translation should be "hallo"
    When I translate the word to "finnish"
    Then the translation should be "hei"


  Scenario: Case-insensitive word
    Given a word "HELLO"
    When I translate the word to "english"
    Then the translation should be "hello"

  Scenario: Case-insensitive language
    Given a word "hello"
    When I translate the word to "ENGLISH
    Then the translation should be "hello"

  Scenario: Unsupported word
    Given a word "xyzzy"
    When I translate the word to "english"
    Then the translated word should be an empty string

  Scenario: Unsupported language
    Given a word "hello"
    When I translate the word to "dutch"
    Then the translated word should be an empty string
