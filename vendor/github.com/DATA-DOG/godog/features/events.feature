Feature: suite events
  In order to run tasks before and after important events
  As a test suite
  I need to provide a way to hook into these events

  Background:
    Given I'm listening to suite events

  Scenario: triggers before scenario event
    Given a feature path "features/load.feature:6"
    When I run feature suite
    Then there was event triggered before scenario "load features within path"

  Scenario: triggers appropriate events for a single scenario
    Given a feature path "features/load.feature:6"
    When I run feature suite
    Then these events had to be fired for a number of times:
      | BeforeSuite    | 1 |
      | BeforeFeature  | 1 |
      | BeforeScenario | 1 |
      | BeforeStep     | 3 |
      | AfterStep      | 3 |
      | AfterScenario  | 1 |
      | AfterFeature   | 1 |
      | AfterSuite     | 1 |

  Scenario: triggers appropriate events whole feature
    Given a feature path "features/load.feature"
    When I run feature suite
    Then these events had to be fired for a number of times:
      | BeforeSuite    | 1  |
      | BeforeFeature  | 1  |
      | BeforeScenario | 6  |
      | BeforeStep     | 19 |
      | AfterStep      | 19 |
      | AfterScenario  | 6  |
      | AfterFeature   | 1  |
      | AfterSuite     | 1  |

  Scenario: triggers appropriate events for two feature files
    Given a feature path "features/load.feature:6"
    And a feature path "features/multistep.feature:6"
    When I run feature suite
    Then these events had to be fired for a number of times:
      | BeforeSuite    | 1 |
      | BeforeFeature  | 2 |
      | BeforeScenario | 2 |
      | BeforeStep     | 7 |
      | AfterStep      | 7 |
      | AfterScenario  | 2 |
      | AfterFeature   | 2 |
      | AfterSuite     | 1 |

  Scenario: should not trigger events on empty feature
    Given a feature "normal.feature" file:
      """
      Feature: empty

        Scenario: one

        Scenario: two
      """
    When I run feature suite
    Then these events had to be fired for a number of times:
      | BeforeSuite    | 1 |
      | BeforeFeature  | 0 |
      | BeforeScenario | 0 |
      | BeforeStep     | 0 |
      | AfterStep      | 0 |
      | AfterScenario  | 0 |
      | AfterFeature   | 0 |
      | AfterSuite     | 1 |

  Scenario: should not trigger events on empty scenarios
    Given a feature "normal.feature" file:
      """
      Feature: half empty

        Scenario: one

        Scenario: two
          Then passing step

        Scenario Outline: three
          Then passing step

          Examples:
            | a |
            | 1 |
      """
    When I run feature suite
    Then these events had to be fired for a number of times:
      | BeforeSuite    | 1 |
      | BeforeFeature  | 1 |
      | BeforeScenario | 2 |
      | BeforeStep     | 2 |
      | AfterStep      | 2 |
      | AfterScenario  | 2 |
      | AfterFeature   | 1 |
      | AfterSuite     | 1 |
