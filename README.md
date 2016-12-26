# Example GraphQL Query

For example on GrahpQL Query object with Golang.

## Usage

For using this example you may need to installed Golang compiler on your machine,
and running following these commands.

```
$ git clone https://github.com/nomkhonwaan/example-graphql-query.git
$ cd /path/to/example-graphql-query
```

For running directly use this command.

```
$ go run main.go
```

And if you want to compile before running it, just using these commands.

```
$ go build -o graphql-server main.go
$ ./graphql-server
```

## Query via cURL

For using the GraphQL server via cURL like this.

**Retrieve people**
```
$ curl -H "Content-Type: application/graphql" \
       -d "{ \
             people { \
               id \
               email \
             } \
           }" \
       http://localhost:8080/graphql
```

**Retrieve people by "gender"**
```
$ curl -H "Content-Type: application/graphql" \
       -d "{ \
             people(gender: \"Male\") { \
               id \
               email \
               gender \
             } \
           }" \
       http://localhost:8080/graphql
```

## Query via Postman

Import this collection to your Postman application https://www.getpostman.com/collections/6c5f187d94244a71e955