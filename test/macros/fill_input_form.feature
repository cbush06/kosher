Feature: Fill Input Form Macros
    Macros related to filling an input form.

    Scenario: I fill the input form
        Fill the Selenium Easy input form demo.

        Then I fill in "first_name" with "Clinton"
        And I fill in "last_name" with "Bush"

    Scenario: I fill the rest of the input form
        Fill the other fields in the Selenium Easy input form demo.

        Then I fill in "email" with "cbush06@gmail.com"
        And I fill in "phone" with "(229)308-2222"
        And I fill in "address" with "222 Ausley Road"
        And I fill in "city" with "Albany"
        And I select "Georgia" from "state"
        And I fill in "zip" with "31707"
        And I choose the "Yes" radio