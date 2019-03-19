Feature: Form Filling Macros
    Macros related to filling forms.

    Scenario: I fill the input form
        Fill the Selenium Easy input form demo.

        Then I fill in "first_name" with "John"
        And I fill in "last_name" with "Doe"

    Scenario: I fill the rest of the input form
        Fill the other fields in the Selenium Easy input form demo.

        Then I fill in "email" with "jdoe@gmail.com"
        And I fill in "phone" with "(444)333-2222"
        And I fill in "address" with "222 Ausley Road"
        And I fill in "city" with "Albany"
        And I select "New York" from "state"
        And I fill in "zip" with "333333"
        And I choose the "Yes" radio