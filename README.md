# Technical assessment

Technical assessment for a Loan Startup

## Requirements (Only for testing in localhost)

- Go 1.x

## How to run this API

### On localhost

Run the API with:

```sh
go run main.go
```

### On public API (AWS)

This API is also deployed on AWS using API Gateway service. Just hit this URL for your API requests:

[https://ak694wjtf2.execute-api.us-east-1.amazonaws.com/Dev/credit-assignment](https://ak694wjtf2.execute-api.us-east-1.amazonaws.com/Dev/credit-assignment)

_NOTE: The above public API URL will be available just for a few days to be evaluated by the technical assessment reviewers._

## How to use it

### Request

Send a **POST** request to the `/credit-assignment` endpoint including the investment key and value to the body request:

```
{
    "investment": <amount>
}
```

- _**\<amount\>**_: Any positive integer multiple of 100

### Response

- **HTTP 200 OK**

    If the investment given can be assigned in valid credits

- **HTTP 400 Bad Request**

    If the investment given is not valid

### Examples

- **Sending a valid investment**

Request body

```json
{
    "investment": 1100
}
```

Response: _HTTP 200 OK_

```json
{
    "credit_type_300": 2,
    "credit_type_500": 1,
    "credit_type_700": 0
}
```

- **Sending an invalid investment**

Request body

```json
{
    "investment": 400
}
```

Response: _HTTP 400 Bad Request_

```json
{
    "credit_type_300": 0,
    "credit_type_500": 0,
    "credit_type_700": 0
}
```
