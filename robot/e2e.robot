*** Settings ***
Documentation     A test suite with a single test for Installation.
...

*** Test Cases ***
Installation
    Install via Helm
    [Teardown]    Clean up