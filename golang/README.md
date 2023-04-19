# Introduction

Unit test with fibonacci sequence. This project explores the golang standard
testing framework as well as [Ginkgo][1], which is a BDD test framework.

## Run coverage

    go test –coverprofile=coverage.out
    go tool cover -html coverage.out

## Run benchmark

    go test -bench=BenchmarkFibonacci


[1]: https://onsi.github.io/ginkgo/
