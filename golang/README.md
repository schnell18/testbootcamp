# Introduction

Unit test with fibonacci sequence. This project explores the golang standard
testing framework as well as [Ginkgo][1], which is a BDD test framework.

## Run unit test

    go test ./...

## Run benchmark

    go test -bench=.

## Run coverage

    go test –coverprofile=coverage.out
    go tool cover -html coverage.out

[1]: https://onsi.github.io/ginkgo/
