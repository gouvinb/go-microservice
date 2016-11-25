#!/usr/bin/groovy
node ('linux') {
  stage 'checkout code'
  checkout scm
  env
  sh 'pwd'
  sh 'ls -lisahF'

  stage 'building libraries'
  sh 'docker run --rm -v $(pwd):/go/src/github.com/gouvinb/go-microservice golang:latest go build github.com/gouvinb/go-microservice'

  stage 'test'
  sh 'docker run --rm -v $(pwd):/go/src/github.com/gouvinb/go-microservice golang:latest go test github.com/gouvinb/go-microservice'
}