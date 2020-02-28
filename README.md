# Mercari-API
Mercari-API is an unofficial API built using GO to fetch information about product listings and sellers on Mercari.com

## Installation & Run
```bash
# Download this project
go get github.com/nicholasc861/mercari-api
```
To run the API server, ensure that you run the following commands:

```bash
# Build and Run
cd mercari-api
go build

# API Endpoint : http://127.0.0.1:8080
```

## Structure
```
├── app
│   ├── app.go
│   └── handler          // Our API core handlers
│       ├── handler.go   // Response functions
│       ├── helper.go    // Helpers functions for handlers
│       └── types.go     // Type structures
└── main.go
```

## API
#### /products/{keyword}
* `GET` : Get all products matching the query keyword

#### /product/{id}
* `GET` : Get information about the product matching the specific id

#### /user/{id}
* `GET` : Get information about the user matching the specific id