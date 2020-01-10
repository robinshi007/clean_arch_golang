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

## Todo List

* [done] input model validate and transform
* [done] output model transform and presenter with json/message_pack
* [done] error handle with golang 1.13 unwrap
* [done] response format standard
* [done] add GraphQL endpoint
* [done] error handler(404, 5xx)
* [done] SQL format library sqlf
* [done] dbm with SQL expression log
* [doing] add context for handler, usecase, repository and respond
* [doing] unit testing, integration testing, e2e testing
* [doing] acceptance testing
* [done] add graphql error to middleware
* [doing] full exmaple to redirect service(rest, graphql, grpc, testing, doc, ui)
* [doing] adpting sqlx and squirrel
* [doing] add middleware limiter, cors, crsf?, profiler and appdash
* dev with ticktock

## Dev Circle

* phase 1: requirement, customer and competition product analysis
* phase 2: prototype with html/css/js, min run with tech stack.
* phase 3: coding and testing with feature complete
* phase 4: documentation, safety and integrition enhancement

## Refers

* <https://github.com/go-chi/chi>
* <https://github.com/hatajoe/8am>
* <https://github.com/bxcodec/go-clean-arch>
* <https://github.com/bradford-hamilton/go-graphql-api>
* <https://github.com/tensor-programming/hex-microservice>
* <https://github.com/sourcegraph/sourcegraph>
