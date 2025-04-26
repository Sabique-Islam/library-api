## Features
- **Get all books**
- **Get a book by ID** 
- **Create a new book**
- **Checkout a book**
- **Return a book**
- **Delete a book**

### Cloning

```
git clone https://github.com/Sabique-Islam/library-api
cd library-api
```

```
go mod init example/library-api
```

```
go get github.com/gin-gonic/gin
```

```
go run main.go
```

## Example Usage

### Get all books

```
curl http://localhost:8080/api/books
```

### Get a book by ID

```
curl http://localhost:8080/api/books/1
```

### Create a new book

```
curl -X POST http://localhost:8080/api/books -H "Content-Type: application/json" -d '{"id": "10", "title": "title_name", "author": "author_name", "quantity": 2}'
```

### Checkout a book

```
curl -X PATCH "http://localhost:8080/api/checkout?id=10"
```

### Return a book

```
curl -X PATCH "http://localhost:8080/api/return?id=10"

```

### Delete a book

```
curl -X DELETE "http://localhost:8080/api/books?id=10"
"

```
