## Golang Clean Arch

For golang web apps template, with REST, GraphQL and GRPC as access endpoint

## Feature

* Stack (Golang, Redis[cache], Postgres[sqldb], Nats[message bus])
* Library(chi, migrate, pq, sqlf, sqlhooks, error1.13, graphql)
* Endpoint (REST, GraphQL and GRPC)
* Serializer (JSON, MessagePack)
* Database (Postgres, SQLite3)
* Logging, tracing and metrics(zerolog, opentracing and prometheus)
* Security(JWT, UUID, CRSF, DDOS)
* Test support(Unit test, Integration test and e2e test)
* Clean arch with loosely coupled, easy to extand
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
* [done] add graphql error to middleware
* [done] add context for handler, usecase, repository and respond
* [doing] full exmaple to redirect service(rest, graphql, grpc, testing, doc, ui)
* [ ] unit testing, integration testing, e2e testing
* [ ] adpting sqlx and squirrel
* add middleware cors, crsf?, profiler
* opentracing & metrics
* version with migration support
* docker deploy support
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
