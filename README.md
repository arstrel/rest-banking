# Banking project

This project consist of the following services:

- [REST api](https://github.com/arstrel/rest-banking)
- [Auth](https://github.com/arstrel/auth-banking)

# rest-banking

Main REST service for banking app.

Learning objectives:

- Mechanism of HTTP web server
- Handler Functions and Router (Request Multiplexer)
- Request and Response Headers
- Converting (Marshal) data structures to JSON and XML
- Connecting and Working with MySQL DB
- Introduce structured logger by zap
- Introduce Sqlx for dealing with MySQL DB
- Introduce Data Transfer Object (DTO) to interact between user side <-> business side <-> server side

## Environment variables

Following env vars are required to run the app:
SERVER_ADDRESS\
SERVER_PORT\
DB_USER\
DB_PASSWD\
DB_PORT\
DB_ADDR\
DB_NAME\

Currently set in `start.sh` for local environment

## Hexagonal Architecture

[Blogpost from Netflix](https://netflixtechblog.com/ready-for-changes-with-hexagonal-architecture-b315ec967749)

![Architecture at a high level](https://miro.medium.com/max/1400/1*NfFzI7Z-E3ypn8ahESbDzw.png)

## Routes

GET /customers -> get all customers in JSON from local MySQL DB\
GET /customers/{numeric customer id} -> get single customer as JSON by id\
GET /customers?status=active | inactive -> get all customers by status.\
POST /customers/{customer id}/account/{account id} -> make a transaction that can be "Deposit" or "Withdraw".\
POST /customers/{customer id}/account -> Create new account that can be "Saving" or "Checking"

## Mock Routes

GET /mock/customers -> get all mock customers in JSON from mock slice of customers\
GET /mock/customers/{numeric customer id} -> get single mock customer as JSON by id\
GET /mock/customers/{numeric customer id} -> get single mock customer as JSON by id\
GET /mock/customers?status=active | inactive -> get all customers by status.
