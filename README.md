# rest-banking

Learning objectives

- Mechanism of HTTP web server
- Handler Functions and Router (Request Multiplexer)
- Request and Response Headers
- Converting (Marshal) data structures to JSON and XML
- Connecting and Working with MySQL DB
- Introduce structured logger by zap
- Introduce Sqlx for dealing with MySQL DB
- Login functionality

## Hexagonal Architecture

[Blogpost from Netflix](https://netflixtechblog.com/ready-for-changes-with-hexagonal-architecture-b315ec967749)

![Architecture at a high level](https://miro.medium.com/max/1400/1*NfFzI7Z-E3ypn8ahESbDzw.png)

## Routes

GET /customers -> get all customers in JSON from local MySQL DB
GET /customers/{numeric customer id} -> get single customer as JSON by id
GET /customers?status=active | inactive -> get all customers by status.

## Mock Routes

GET /mock/customers -> get all mock customers in JSON from mock slice of customers
GET /mock/customers/{numeric customer id} -> get single mock customer as JSON by id
GET /mock/customers/{numeric customer id} -> get single mock customer as JSON by id
GET /mock/customers?status=active | inactive -> get all customers by status.
