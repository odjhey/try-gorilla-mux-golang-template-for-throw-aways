- try golang for http apis
- also try `ninja`, been using `make` for a long time and it is serving me well, but doesn't hurt to try new things

- use this as template for throw away api endpoints used during manual testing

- [docs, swagger files, postman](docs/)

# TODO

- [x] use sqlite3 for some persistence
- [x] create swagger doc
- [ ] add endpoint test reporter (see pactum.js, frisbyjs, pact.io)

# Refs

- https://sqlite.org/cli.html
- https://gorm.io/docs/query.html
- https://github.com/gorilla/mux

```
sqlite3
.open .data.db
.mode line
select * from products;
```
