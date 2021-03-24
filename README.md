![Gollang workflow](https://github.com/maxwellgithinji/customer_orders/actions/workflows/ci.yml/badge.svg) [![Maintainability](https://api.codeclimate.com/v1/badges/eacafa30d30ccb157fc1/maintainability)](https://codeclimate.com/github/maxwellgithinji/customer_orders/maintainability) [![Test Coverage](https://api.codeclimate.com/v1/badges/eacafa30d30ccb157fc1/test_coverage)](https://codeclimate.com/github/maxwellgithinji/customer_orders/test_coverage)

# customer_orders

A simple go REST service to show implementation of PostgreSQL, CI/CD, Open ID Connect, testing, deployment

- [Link to Live app](https://customer-orders-api.herokuapp.com/swagger/index.html)

## Setup

## Database

- create a postgresql db locally
- Use `DATABASE_URL` env var to reference it
- run `make migrateup` to create tables
- run `make migratedowm` to remove tables

### Server

- run `$ make server` to start the application

## Open Api

- Set up open api by following [this guide](https://auth0.com/docs/quickstart/backend/golang/01-authorization)
- test `/login` and `/logout` routes on this app to ensure everything is working
- **Note** a `/callback` route on this app checks whether a user has an account in this system before saving them

## Africas Talking

- go to [Africa's Talking](https://account.africastalking.com/apps/sandbox/sms/bulk/outbox) sandbox and create an API key for the app
- you can launch a simulator with the target phone numbers you will use in your app

### Testing and coverage

- use the following commands to run tests

- `$ make test`
- `$ make testcover`
- `$ make testview`

## Flow

### Authentication

- use `GET /login` to get the redirect url for login
- open `redirect` url from login response and you will be prompted to log in by OAuth

### User Management

- go to `GET /auth/profile` to get the profile information of current user, for a new user, the `status` is `inactive` until they have onboarded
- use `POST /auth/onboard` to update phone number and username. If you don't update your phone number, or if your `status` is not `active` you wont be able to make orders

### Item management

- use `POST /auth/item` (will require admin account in future) to add items to order
- use `GET /auth/item` (will require pagination in future) to view all items
- use `DELETE /auth/item/{id}` (will require admin account in future) to delete an item

### Orders

- use `POST /auth/orders` and enter the item id in the body to order an item
- go to Africa's talking and ensure you received an SMS
- go to `GET /auth/currentuser/orders` to view orders of currently logged in user
- go to `GET /auth/orders` to get all orders ever made (Pagination and admin o=implementation required in future)

## Log Out

- go to `POST /logout` to log out of the application
- all routes with `*/auth/*` should give an error 401 when you try accessing them when you are not logged in
