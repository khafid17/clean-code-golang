1. External system perform request (HTTP, Messaging, etc)
2. The Handler creates various Model from request data
3. The Handler calls Controller, and execute it using Model data
4. The Usecase create Entity data for the business logic
5. The Usecase calls Repository, and execute it using Entity data
6. The Repository use Entity data to perform database operation
7. The Repository perform database operation to the database
8. The Gateway using Model data to construct request to external system
9. The Gateway perform request to external system (HTTP, Messaging, etc)

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
