# REST API example written in Go 

### Requirements

To be able to show the desired features of curl this REST API must match a few
requirements:

* [x] `GET /employees` returns list of employees as JSON
* [x] `GET /employees/{id}` returns details of specific employee as JSON
* [x] `POST /employees` accepts a new employee to be added
* [x] `POST /employees` returns status 415 if content is not `application/json`
* [x] `GET /employees/random` redirects (Status 302) to a random employee

### Data Types

A employee object should look like this:
```json
  {
    "id": 1,
    "email_address": "aronprenovostmktg@gmail.com",
    "user_name": "aronp123@gmail.com",
    "first_name": "Aron",
    "last_mame": "Prenovost",
    "city": "Seattle",
  }
```

### Persistence

Data is stored is MySQL 

### API endpoints

## employees

* `GET /employees` returns list of employees as JSON
```
$ curl localhost:8080/employees
```

* `POST /employees` accepts a new employee to be added
```
$ curl localhost:8080/employees -X POST -d '{"name": "User121", "city": "Seattle"}' -H "Content-Type: application/json"
```

* `PUT /employees` update employee record
```
$ curl localhost:8080/employees -X PUT -d '{"name": "Aronias", "city": "Kent", "id": 54}' -H "Content-Type: application/json"
```

* `DELETE /employees` deletes all employee records 
```
$ curl localhost:8080/employees -X DELETE
```

### Long-Term Goals 
### Build a real world "production" REST API: 

* [ ] Scalable, must be able to run more than one instance.

* [ ] Dockerized, runnable on minikube.

* [ ] Unit tested, must be able to run "go test ./..." directly from clone.

* [ ] Integration tested, recommend docker-compose.

* [ ] OpenAPI/Swagger (or similar for gRPC) documented.

* [ ] dep vendored, but using the standard library often, instead of piling on dependencies.

* [ ] Authenticated and Authorized via apikeys and/or user accounts.

* [ ] Built and tested via CI: Travis, CircleCi, etc. Recommend Makefile for task documentation.

* [ ] Flag & ENV config, API keys, ports, dev mode, etc.

* [ ] "why" comments, not "what" or "how" which should be clear through func/variable names and godoc comments.

* [ ] Use of Context to limit request time.

* [ ] Leveled JSON logging, logrus or similar.

* [x] Postgres/MySQL, sqlx or an ORM.

* [ ] Redis/memcache for scalable caching.

* [ ] Datadog, New Relic, AppDynamics, etc for monitoring and statistics.

* [ ] Well documented README.md with separate sections for API user and service developer audiences. Maybe even include graphviz or mermaidJS UML diagrams.

* [ ] Clean git history with structured commits and useful messages. No merge master commits.

* [ ] Passing go fmt, go lint, or better, go-metalinter in the CI.


### Running Project 

### generate .env and assign environment + db connection variables
```
$ cp .env-sample .env
```

### Build go module
```
$ go build
```

### build Docker image 
```
$ docker build -t {image name} .
```

### run project 
```
$ docker run -p 80:80 {image name}
```
 