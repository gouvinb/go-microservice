# Go-microservice

Basic microservice in Go.

This projet has been used in my school to build a microservice using Golang
without a framework and with ease. Inspired by
[gowebapp](https://github.com/josephspurrier/gowebapp)
writted by [Joseph Spurrier](http://www.josephspurrier.com/about/). My group
wanted this microservice to match docker container structure.

To download, run the following command:

```bash
# Prerequisites to use 'go generate' command
go get -u github.com/jteeuwen/go-bindata/...
go get github.com/gouvinb/go-microservice
```

## Prerequisites

### Go 1.6 and lasted version

This project is fully compatible.

### Go 1.5

You need to set `GOVENDOREXPERIMENT` to 1 to use vendor folder.

### Go 1.4 and earlier

This project isn't compatible as you do not update your imports.

## Features

-   Configuration for developpement and production
-   Compatible with lot of databases (MySQL, MariaDB, Bolt, Mongo...)

## Structure

```java
.
|-- main.go
|-- template
|   |-- base.tmpl
|   |-- footer.tmpl
|   |-- error
|   |   `-- [CODE_ERROR].tmpl
|   `-- [VIEW_NAME]
|       |-- [SUB_VIEW]
|       |   `-- [SUB_VIEW_NAME].tmpl
|       `-- [VIEW_NAME].tmpl
`-- vendor
    |-- config
    |   |-- config.go
    |   `-- config.json
    |-- controller
    |   |-- doc.go
    |   `-- [CONTROLLER].go
    |-- model
    |   |-- doc.go
    |   `-- [MODEL].go
    |-- plugin
    |   `-- [PLUGIN].go
    |-- route
    |   |-- middleware
    |   |   |--doc.go
    |   |   `-- [MIDDLEWARE].go
    |   |-- routewrapper
    |   |   |-- doc.go
    |   |   `-- [ROUTERWRAPPER].go
    |   `-- route.go
    |-- shared
    |   |-- doc.go
    |   |-- cors.go
    |   |-- database.go
    |   |-- server.go
    |   |-- session.go
    |   |-- view.go
    |   `-- [SHARED]-utils.go
    `-- utils
        `-- [UTILS].go
```

### Main package

It will initialize the logs and load `config.json` by means of `go-bindata`
before launch microservice.

### Config

#### Go

Config package will load a default config file (`config.json`) by means of
generating go-bindata.

#### JSON

`Config.json` contains all environement, it needs to convert Go
source code before build into managable (_see Main.go at line 17_).

**The default configuration is based on the default values from shared package.
You should edit at first  the `config.json` file instead of values from shared
package.**

### Controller

Controller package can send commands to update the model's state and to his
associated view to change the view's presentation of the model.

### Model

Model package stores data which is retrieved according to commands from the
controller.

### Route

Route package will load router for web server.

#### Middleware

Middleware package allows the use of http.HandlerFunc compatible funcs with
`julienschmidt/httprouter`.

#### RouteWrapper

Routewrapper package is a wrapper for better implementation of the routes.

### Shared

Package shared contain all microservice config :

-   Cors file provides authorization for cross origin requests.
    (Only if have no session)
-   Database file provides an interface to manage database.
-   Server file is a wrapper around the net/http package that starts listeners
    for HTTP and HTTPS.
-   Session provides a wrapper for gorilla/sessions package. (Only if cors
    disabled)
-   View provides an interface to display view at an user.

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
