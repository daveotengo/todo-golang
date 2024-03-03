Instructions

1. go run main.go //to start application

2. curl http://localhost:8080/books //to make a get request

3. curl http://localhost:8080/books --include --header "Content-Type:application/json" -d @body.json --request "POST"  // to make post request
4. curl http://localhost:8080/books/2 //to make a get request by Id

   
5. curl 'http://localhost:8080/checkout?id=2' --request "PATCH" //to checkout a book

6. curl 'http://localhost:8080/return?id=2' --request "PATCH" //to return a book

