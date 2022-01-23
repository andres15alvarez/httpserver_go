# API Rest Server
### This server use only the standard library of go (http, fmt and json). Use routers, handlers and middlewares.
### To start the server:
`go run .`
### or
`go build -o ./build/server && ./build/server`
### The server will be running when the following message print in the terminal:
`Server running in port :8000`

## Endpoints
#### The following examples use httpie.
### Home:
`http :8000`
#### Response:
`{"data": "Hello, World"}`

### Retrieve users:
`http :8000/user Authorization:hola`
#### Response:
```
TTP/1.1 200 OK
Content-Length: 138
Content-Type: application/json
Date: Sun, 23 Jan 2022 21:01:49 GMT

[
    {
        "email": "johndoe@example.com",
        "name": "John Doe"
    },
    {
        "email": "bob@example.com",
        "name": "Bob"
    },
    {
        "email": "alice@example.com",
        "name": "Alice"
    }
]

```

### Create an user:
`http :8000/user Authorization:hola name=Andres email="andres@example.com"`
#### Response:
```
HTTP/1.1 201 Created
Content-Length: 47
Content-Type: application/json
Date: Sun, 23 Jan 2022 20:45:54 GMT

{
    "email": "andres@example.com",
    "name": "Andres"
}
```

#### If you do not provide the Authorization header, then you got a 401 response:
```
HTTP/1.1 401 Unauthorized
Content-Length: 0
Date: Sun, 23 Jan 2022 20:47:24 GMT
```