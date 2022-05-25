# rest-banking

Learning objectives

- Mechanism of HTTP web server
- Handler Functions and Router (Request Multiplexer)
- Request and Response Headers
- Converting (Marshal) data structures to JSON and XML

Routes:
GET /greeting -> Hello world
GET /customers -> list of customers in json by default of in xml if header "Content-Type: application/xml"
GET /customers/{numeric value}
POST /customers -> Post request received
