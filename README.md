# token
As part of the recruitment process we want to know how you think, code and structure your work. In order to do that, we're going to ask you to complete this coding challenge.  


There is also a "hello world" go application in cmd/ and a docker-compose.yml for running Amazon postgres locally.

We need you to:

Provide a Go implementation of the http service in the cmd/ directory of this repo.
Implement a postgres based store for this HTTP service
Provide adequate test coverage for this simple service

## urls
- post localhost:8080/v1/admin/login
- post localhost:8080/v1/admin/user-login
- post localhost:8080/v1/auth/create-token