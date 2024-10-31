# Syndio Backend App

Using the `employees.db` sqlite database in this repository with the following table/data:

```
sqlite> .open employees.db
sqlite> .schema employees
CREATE TABLE employees (id INTEGER PRIMARY KEY, gender TEXT not null);
sqlite> SELECT * FROM employees;
1|male
2|male
3|male
4|female
5|female
6|female
7|non-binary
```

Create an API endpoint that saves job data for the corresponding employees.

Example job data:

```json
[
  {
    "employee_id": 1,
    "department": "Engineering",
    "job_title": "Senior Enginer"
  },
  {
    "employee_id": 2,
    "department": "Engineering",
    "job_title": "Super Senior Enginer"
  },
  { "employee_id": 3, "department": "Sales", "job_title": "Head of Sales" },
  { "employee_id": 4, "department": "Support", "job_title": "Tech Support" },
  {
    "employee_id": 5,
    "department": "Engineering",
    "job_title": "Junior Enginer"
  },
  { "employee_id": 6, "department": "Sales", "job_title": "Sales Rep" },
  {
    "employee_id": 7,
    "department": "Marketing",
    "job_title": "Senior Marketer"
  }
]
```

## Requirements

- The API must take an environment variable `PORT` and respond to requests on that port.
- You provide:
  - Basic setup instructions required to run the API
  - Guide on how to ingest the data through the endpoint
  - A way to update the existing database given to you

## Success

- We can run the API and ingest database on your setup instructions
- The API is written in Python or Go

## Not Required

- Tests
- Logging, monitoring, or anything more than basic error handling

## Submission

- Respond to the email you received giving you this with:
  - a zip file, or link to a git repo
  - instructions on how to setup and run the code (could be included w/ zip/git)
- We'll follow the instructions to test it on a local machine, then we'll get back to you

## Notes

- Keep it simple
- If the API does what we requested, then it's a success
- Anything extra (tests, other endpoints, ...) is not worth bonus/etc
- We expect this to take less than two hours, please try and limit your effort to that window
- We truly value your time and just want a basic benchmark and common piece of code to use in future interviews
- If we bring you in for in-person interviews, your submission might be revisited and built on during the interview process

--

## Instructions

At the project root, do the following in a terminal:

1. Install project dependencies

```
go mod download
```

2. Export environment variable for your desired port

```
export PORT=<DESIRED_PORT>
```

3. Run the API

```
go run .
```

4. Make a request to the API on your desired port. If desired port is not provided, the API will respond to requests on port `8080` by default.

```
curl --location 'http://localhost:<DESIRED_PORT>/api/v1/jobs' \
--header 'Content-Type: application/json' \
--data '[
  { "employee_id": 1, "department": "Engineering", "job_title": "Senior Enginer" },
  { "employee_id": 2, "department": "Engineering", "job_title": "Super Senior Enginer" },
  { "employee_id": 3, "department": "Sales", "job_title": "Head of Sales"},
  { "employee_id": 4, "department": "Support", "job_title": "Tech Support" },
  { "employee_id": 5, "department": "Engineering", "job_title": "Junior Enginer" },
  { "employee_id": 6, "department": "Sales", "job_title": "Sales Rep" },
  { "employee_id": 7, "department": "Marketing", "job_title": "Senior Marketer" }
]'
```
