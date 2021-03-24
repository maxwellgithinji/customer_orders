![Gollang workflow](https://github.com/maxwellgithinji/customer_orders/actions/workflows/ci.yml/badge.svg)  [![Maintainability](https://api.codeclimate.com/v1/badges/eacafa30d30ccb157fc1/maintainability)](https://codeclimate.com/github/maxwellgithinji/customer_orders/maintainability)  [![Test Coverage](https://api.codeclimate.com/v1/badges/eacafa30d30ccb157fc1/test_coverage)](https://codeclimate.com/github/maxwellgithinji/customer_orders/test_coverage)

# customer_orders

A simple go REST service to show implementation of PostgreSQL, CI/CD, Open ID Connect, testing, deployment

- [Link to Live app](https://customer-orders-api.herokuapp.com/swagger/index.html)

## Setup

### Init

`$ go mod  tidy`

### Swagger

`$ make swaginstall`
`$ make swag`

### Server

`$ make make server`

### Testing and coverage

`$ make test`
`$ make testcover`
`$ make testview`


