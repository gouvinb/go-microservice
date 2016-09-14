# Go-microservice

Basic microservice in Go.

This project used at school for build a microservice using the Go language without a framework and easily in a project with an microservice struct. Inspired by [gowebapp](https://github.com/josephspurrier/gowebapp) writted by [Joseph Spurrier](http://www.josephspurrier.com/about/). It was rewritten with the objective of being adapted a docker container.

To download, run the following command:

```bash
go get github.com/gouvinb/go-microservice
```

## Features

## Structure

```txt
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

### Main.go

### Config

#### Go

#### JSON

### Controller

### Model

### Route

#### Go

#### Middleware

### Shared

#### database

#### server

#### session

#### Utils

## Feedback

All feedback is welcome. Let me know if you have any suggestions, questions, or criticisms.
If something is not idiomatic to Go, please let me know know so we can make it better.

## LICENSE

Copyright 2016 gouvinb. All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE.md file.
