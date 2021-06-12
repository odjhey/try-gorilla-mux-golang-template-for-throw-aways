- try golang for http apis
- also been wanting to try `ninja`, been using `make` for a long time and it is serving me well, but doesn't hurt to try new things

- use this as template for throw away api endpoints used during manual testing

- sample [postman collection here](docs/golang-try-gorilla-mux.postman_collection.json)

# TODO

- [x] use sqlite3 for some persistence
- [ ] create swagger doc

# Refs

- https://sqlite.org/cli.html
- https://gorm.io/docs/query.html

```
sqlite3
.open .data.db
.mode line
select * from products;
```
