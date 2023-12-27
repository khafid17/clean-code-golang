External system perform request (HTTP, Messaging, etc)
The Handler creates various Model from request data
The Handler calls Controller, and execute it using Model data
The Usecase create Entity data for the business logic
The Usecase calls Repository, and execute it using Entity data
The Repository use Entity data to perform database operation
The Repository perform database operation to the database
The Gateway using Model data to construct request to external system
The Gateway perform request to external system (HTTP, Messaging, etc)

Invoice Rest API can 3 endpoit
1. Form Input
2. Create Data
3. Get All Data

Include collection postman and database.sql, program done implement Dependency Injection and Unit Test.

Build program with Clean Architecture in Golang
step by step installation and running program
1. clone project or extract project
2. running "go mod tidy"
3. import Mysql or running query in folder migration "001_esb_up.sql"
4. import collection json in postman
5. running project "go run main.go"
6. hit endpoit by postman
7. done 
