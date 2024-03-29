# ShakeSearch Challenge

Welcome to the Pulley Shakesearch Challenge! This repository contains a simple web app for searching text in the complete works of Shakespeare.

## Prerequisites

To run the tests, you need to have [Go](https://go.dev/doc/install) and [Docker](https://docs.docker.com/engine/install/) installed on your system.

## Your Task

Your task is to fix the underlying code to make the failing tests in the app pass. There are 3 frontend tests and 3 backend tests, with 2 of each currently failing. You should not modify the tests themselves, but rather improve the code to meet the test requirements. You can use the provided Dockerfile to run the tests or the app locally. The success criteria are to have all 6 tests passing.

## Instructions

<img width="404" alt="image" src="https://github.com/ProlificLabs/shakesearch/assets/98766735/9a5b96b5-0e44-42e1-8d6e-b7a9e08df9a1">

*** 

**Do not open a pull request or fork the repo**. Use these steps to create a hard copy.

1. Create a repository from this one using the "Use this template" button.
2. Fix the underlying code to make the tests pass
3. Include a short explanation of your changes in the readme or changelog file
4. Email us back with a link to your copy of the repo

## Running the App Locally


This command runs the app on your machine and will be available in browser at localhost:3001.

```bash
make run
```

## Running the Tests

This command runs backend and frontend tests.

Backend testing directly runs all Go tests.

Frontend testing run the app and mochajs tests inside docker, using internal port 3002.

```bash
make test
```

Good luck!

## Changelog

**Changes for Backend Tests**

To fix TestSearchCaseSensitive: Store the complete works as lowercase bytes in SuffixArray (in the Load function) & also lowercase the query string before performing the lookup against lowercased suffix array (in the Search function)All code changes were made in main.go.

To fix TestSearchDrunk: Update the Search function to return results in pages of max 20 results each. For this I added support for a "page" param to the /search request, so frontend can increment this param and pass through to /search for the next page of results. handleSearch then passes this param through to the Search function eventually. All code changes were made in main.go.

**Changes for Frontend Tests**

should return search results for "romeo, wherefore art thou": I think this test was fixed by the fix for TestSearchCaseSensitive.

should load more results for "horse" when clicking "Load More": Added an event listener to the "Load More" button to actually pass through the incremented page number and append the next page of results to the existing results from initial "Search". Updated app.js to pass through the incremented page number correctly to the /search request on clicking "Load More".

P.S: Hacky but I also had to bump the timeout for mocha in package.json to resolve some timeouts when running frontend tests locally: https://stackoverflow.com/a/51586973
