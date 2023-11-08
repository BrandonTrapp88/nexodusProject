Feature: Organization API

  Background:
    Given a user named "Oliver" with password "testpass"
    Given a user named "Oscar" with password "testpass"
    Given a user named "EvilBob" with password "testpass"

  Scenario: Show basic organization api in action

    #
    # Get the user and default org ids for two users...
    #
    Given I am logged in as "Oliver"
    When I GET path "/api/users/me"
    Then the response code should be 200
    Given I store the ".id" selection from the response as ${oliver_user_id}

    When I GET path "/api/organizations"
    Then the response code should be 200
    Given I store the ${response[0].id} as ${oliver_organization_id}
    Given I store the ${response[0].security_group_id} as ${security_group_id}

    Given I am logged in as "Oscar"
    When I GET path "/api/users/me"
    Then the response code should be 200
    Given I store the ".id" selection from the response as ${oscar_user_id}
    Given I store the ".username" selection from the response as ${oscar_username}

    When I GET path "/api/organizations"
    Then the response code should be 200
    Given I store the ${response[0].id} as ${oscar_organization_id}
    # validate the default org id is the same as the user id
    Then "${oscar_organization_id}" should match "${oscar_user_id}"

    #
    # Oscar should only be able to see the orgs that he is a part of.
    When I GET path "/api/organizations"
    Then the response code should be 200
    And the response should match json:
      """
      [
        {
          "description": "${oscar_username}'s organization",
          "id": "${oscar_organization_id}",
          "name": "${oscar_username}",
          "owner_id": "${oscar_user_id}"
        }
      ]
      """

  Scenario: The user creates an organization with the same name twice to display the 409 error on second creation

    Given I am logged in as "Oliver"
    When I GET path "/api/users/me"
    Then the response code should be 200
    Given I store the ".id" selection from the response as ${oliver_user_id}

    # Create an organization for the first time to simulate happy path
    When I POST path "/api/organizations" with json body:
      """
      {
        "description": "The Blue Zone",
        "name": "zone-blue"
      }
      """

    # The above request will succeed the first time the e2e test is run,
    # but then should fail with an error, so ignore the result of this request.

    # Recreate the same organization to simulate unhappy path
    When I POST path "/api/organizations" with json body:
      """
      {
        "description": "The Blue Zone",
        "name": "zone-blue"
      }
      """
    Given I store the ".id" selection from the response as ${organization_id}
    Then the response code should be 409
    And the response should match json:
      """
      {
        "id": "${organization_id}",
        "error" :"resource already exists"
      }
      """
