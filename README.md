***Goland Api***

Api Making Crud Operations with static Data

**Instructions**

1. Run the following command to start the application:
go run main.go

1. Use the following command to make a GET request for all books:
curl http://localhost:8080/books

1. Use the following command to make a POST request to add a new book:
curl http://localhost:8080/books
--include
--header "Content-Type: application/json"
-d @body.json
--request "POST"

1. Use the following command to make a GET request for a specific book by its ID:
curl http://localhost:8080/books/2

1. Use the following command to checkout a book:
curl 'http://localhost:8080/checkout?id=2'
--request "PATCH"

1. Use the following command to return a book:
curl 'http://localhost:8080/return?id=2'
--request "PATCH"


