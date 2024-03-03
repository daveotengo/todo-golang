1. Run the following command to start the application:
go run main.go

css
Copy code

2. Use the following command to make a GET request for all books:
curl http://localhost:8080/books

css
Copy code

3. Use the following command to make a POST request to add a new book:
curl http://localhost:8080/books
--include
--header "Content-Type: application/json"
-d @body.json
--request "POST"

css
Copy code

4. Use the following command to make a GET request for a specific book by its ID:
curl http://localhost:8080/books/2

css
Copy code

5. Use the following command to checkout a book:
curl 'http://localhost:8080/checkout?id=2'
--request "PATCH"

bash
Copy code

6. Use the following command to return a book:
curl 'http://localhost:8080/return?id=2'
--request "PATCH"

Copy code
