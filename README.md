Deep Sea
========

Small project demonstrating an issue with Golang's code coverage, and JetBrains GoLand.

Reproducing
-----------

Running `go test deepsea`, with default template within Goland 2019.2, with `tests/_placeholder.go`:

```
GOROOT=/usr/local/go #gosetup
GOPATH=/Users/geert/go #gosetup
/usr/local/go/bin/go test -json ./... #gosetup
?   	github.com/geertjanvdk/deepsea	[no test files]
=== RUN   TestSubmarine_Dive
--- PASS: TestSubmarine_Dive (0.00s)
PASS
ok  	github.com/geertjanvdk/deepsea/submarine	0.005s
=== RUN   TestSomething
--- FAIL: TestSomething (0.00s)
    my_test.go:11: I simply fail
FAIL
FAIL	github.com/geertjanvdk/deepsea/tests	0.005s
```

All OK.

Running `go test deepsea`, but now with coverage (clicking on the '... With Coverage' button), coverage is not
showing in JetBrains:

```
GOROOT=/usr/local/go #gosetup
GOPATH=/Users/geert/go #gosetup
/usr/local/go/bin/go test -json -covermode=atomic -coverpkg=./... -coverprofile /Users/geert/Library/Caches/GoLand2019.2/coverage/deepsea$go_test_deepsea.out ./... #gosetup
go build github.com/geertjanvdk/deepsea/tests: no non-test Go files in /Users/geert/go/src/github.com/geertjanvdk/deepsea/tests
?   	github.com/geertjanvdk/deepsea	[no test files]
FAIL	github.com/geertjanvdk/deepsea/submarine [build failed]
=== RUN   TestSomething
--- FAIL: TestSomething (0.00s)
    my_test.go:11: I simply fail
FAIL
coverage: 33.3% of statements in ./...
FAIL	github.com/geertjanvdk/deepsea/tests	0.005s
```

With the above run, you will see result in file [deepsea_no_placeholder.png](_support/deepsea_no_placeholder.png).
Notice the `no non-test Go Files` message.

When renaming the `tests/_placeholder.go` as `tests/placeholder.go`, we will get the result we want [deepsea_with_placeholder.png](_support/deepsea_with_placeholder.png)

```
GOROOT=/usr/local/go #gosetup
GOPATH=/Users/geert/go #gosetup
/usr/local/go/bin/go test -json -covermode=atomic -coverpkg=./... -coverprofile /Users/geert/Library/Caches/GoLand2019.2/coverage/deepsea$go_test_deepsea.out ./... #gosetup
?   	github.com/geertjanvdk/deepsea	[no test files]
=== RUN   TestSubmarine_Dive
--- PASS: TestSubmarine_Dive (0.00s)
PASS
coverage: 33.3% of statements in ./...
ok  	github.com/geertjanvdk/deepsea/submarine	0.005s	coverage: 33.3% of statements in ./...
=== RUN   TestSomething
--- FAIL: TestSomething (0.00s)
    my_test.go:11: I simply fail
FAIL
coverage: 33.3% of statements in ./...
FAIL	github.com/geertjanvdk/deepsea/tests	0.005s
```
