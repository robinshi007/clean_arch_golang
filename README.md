## Golang Clean Arch

For golang web apps template, with REST, GraphQL and GRPC as access endpoint

## Feature

* Stack (Golang, Redis, Postgres/SQLite3, Nats)
* Library(chi, migrate, pq, sqlf, sqlhooks, error1.13, graphql)
* Endpoint (REST, GraphQL and GRPC)
* Serializer (JSON, MessagePack)
* Database (Postgres, SQLite3)
* Logging, tracing and metrics(zerolog, opentracing and prometheus)
* Security(JWT, UUID)
* Test support(Unit test and Integration test)
* Clean arch with loosely coupled
* Docker deploy support

## Todo

* [done] input model validate and transform
* [done] output model transform and presenter json/xml
* [done] error handle with golang 1.13 unwrap
* [done] response format standard
* [doing] add GraphQL endpoint
* [doing] error handler(404, 5xx)
* [done] SQL format library sqlf
* [done] dbm with SQL expression log
* [doing] add context for handler, usecase, repository and respond
* unit test
* [done] integrite test
* e2e test
* dev with ticktock

## Refers

* <https://github.com/go-chi/chi>
* <https://github.com/hatajoe/8am>
* <https://github.com/bxcodec/go-clean-arch>
* <https://github.com/bradford-hamilton/go-graphql-api>
* <https://github.com/tensor-programming/hex-microservice>
* <https://github.com/sourcegraph/sourcegraph>
