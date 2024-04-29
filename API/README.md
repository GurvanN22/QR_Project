# API Documentation 

This API serve the QrCode web server by handling the users and image store by the users

## Configuration

In the .env file we can change :
- The port of the API (int)
- The path of the database (string)
- If we fill the database when empty for tests applications (bool)

## Start the API

### Golang mod init

We start by create the local package to handle the imports in the project. In the "API" folder type 

```sh-session
go mod init api
go mod tidy
```

### Golang imports 

The next step is to import packages

```sh-session
go get github.com/joho/godotenv
go get github.com/mattn/go-sqlite3
go get -u github.com/swaggo/swag/cmd/swag
go mod tidy
```

## Project architecture

```
API
├── data_functions
│   └── creation.go
├── db
│   ├── data.sqlite3
│   ├── db.db
│   ├── images
│   └── query
│       ├── creation.sql
│       └── exemple.sql
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
├── handlers
│   ├── connect-user.go
│   ├── create_user.go
│   ├── delete-image.go
│   ├── documentation.go
│   ├── image.go
│   ├── info-image.go
│   ├── info-user.go
│   ├── new-image.go
│   ├── root.go
│   └── tools
│       ├── chiffrement.go
│       └── method.go
├── main.go
├── README.md
├── server
│   └── server.go
└── test
    ├── esteban.png
    └── test.html
```
## Endpoints 

Here the differents endpoints :

- /api
- /create-user
    - parameters
        - email 
        - password
    - code 
        - 400 (missing fields)
        - 404 (user not found)
        - 200 (user authenticated)
- /connect-user
    - parameters
        - name
        - email
        - password
    - code 
        - 400 (missing fields)
        - 200 (user created)
- /info-user
    - parameters
        - 400 (wrong field : id)
        - 200 (user found)
        - 404 (no data found)
- /new-image
    - parameters
        - link 
        - user_id
        - file (file of the qr)
    - code 
        - 400 (Bad request)
        - 200 (Image added successfully)
- /info-image
    - parameters
        - id 
    - code 
        - 400 (wrong field : id)
        - 404 (no data found)
        - 200 (Image)
- /delete-image
    - parameters
        - id 
    - code 
        - 404 (Image not found)
        - 200 (Image deleted)
- /image/(id)

This endpoint is particular. It work like on youtube with video url with the id in the url 

exemple : 
```
http:localhost:4000/image/FHGUEJDLSJEUFLR5
```

It will return the image with the "FHGUEJDLSJEUFLR5" id 

by EXPOSE  

@ll rights reserved 2024

