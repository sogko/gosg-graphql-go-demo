# gosg-graphql-go-demo
Demo and slides for GoSG Meetup #15 - May 2016

### Slides
Slides can be found in the [`slides`](./slides) folder.

[Download the PDF](./slides/gosg-may-2016-presentation.pdf)


### Demo

To run the GraphQL server
```bash
$ go run server/main.go
```

- Client V1 - [http://localhost:3000/](http://localhost:3000/)
- Client V2 - [http://localhost:3000/v2/](http://localhost:3000/v2/)
- GraphiQL - [http://localhost:3000/graphiql](http://localhost:3000/graphiql)
- GraphQL endpoint - [http://localhost:3000/graphql](http://localhost:3000/graphql)


To re-build the clients (not necessary unless you make changes to the client JS code)

```bash
$ cd client
$ npm run build-v1  # build client version 1
$ npm run build-v2  # build client version 2
```

### License
MIT.
