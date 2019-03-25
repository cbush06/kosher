@test2
Feature: Filling and Valdiating File Fields
    Test steps that interact with file fields. This step must be
    ran with the "test2" environment.

    Scenario: Fill a File Field
        Test filling a file field.

        Given I am on the "file-upload" page
        And I fill "file-upload" with "${RESOURCESDIR}/Kosher.png"
        And I wait 5 seconds
        Then "file-upload" should contain "C:\fakepath\Kosher.png"