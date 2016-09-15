# Go-microservice

Basic microservice in Go.

This projet has been used in my school to build a microservice using the Go
language without a framework and easily in a project with an microservice
struct. Inspired by [gowebapp](https://github.com/josephspurrier/gowebapp)
writted by [Joseph Spurrier](http://www.josephspurrier.com/about/). it was
written again with a goal to be adapted to docker container.

To download, run the following command:

```bash
go get -u github.com/jteeuwen/go-bindata/... # Needed for go generate
go get github.com/gouvinb/go-microservice
```

## Features

-   Configuration for developper and production
-   Adapted for lot of database (MySQL, MariaDB, Bolt, Mongo)

## Structure

```java
.
├── main.go
├── config
│   ├── config.go
│   └── config.json
├── controller
│   ├── doc.go
│   └── [CONTROLLER].go
├── model
│   ├── doc.go
│   └── [MODEL].go
├── route
│   ├── middleware
│   │   ├── doc.go
│   │   └── [MIDDLEWARE].go
│   └── route.go
└── shared
    ├── doc.go
    ├── database.go
    ├── server.go
    ├── session.go
    └── [SHARED]-utils.go
```

### Main package

It handles initializing the logs and load config.json throught go-bindata befor
launch microservice.

### Config

#### Go

Config load a default config file (_config.json_) throught generate go-bindata.

#### JSON

Config.json contains all environement, it needs converts into managable Go
source code before build (_see Main.go at line 17_).

**The default configuration is based on the default values from shared package.
Please edit in primary the config.json file instead of values from shared
package.**

### Controller

Package controller can send commands to the model to update the model's
state. It can also send commands to its associated view to change the view's
presentation of the model.

#### Example

<!--
* POST   http://localhost/value      - Create a new value
* GET    http://localhost/value      - Retrieve a list of all values
* PUT    http://localhost/value/{id} - Update a value by ID
* DELETE http://localhost/value      - Delete all values by ID
-->

### Model

Package model stores data that is retrieved according to commands from the
controller.

#### Example

<!--
* CREATE
* GET LIST
* GET ONE
* EDIT
* DELETE ALL
* DELETE ONE
-->

### Route

#### Go

#### Middleware

### Shared

#### Database

#### Server

#### Session

#### Utils

## Feedback

All feedback is welcome. Let me know if you have any suggestions, questions, or
criticisms.
If something is not idiomatic to Go, please let me know know so we can make it
better.

## Contributing

1.  Fork it!
2.  Create your feature branch: git checkout -b my-new-feature
3.  Commit your changes: git commit -am 'Add some feature'
4.  Push to the branch: git push origin my-new-feature
5.  Submit a pull request :D

## LICENSE

Copyright 2016 gouvinb. All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE.md file.
