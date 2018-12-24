fmt:
	goimports -w $$(find . -type f -name '*.go' -not -path "./vendor/*")

test:
	go test -v ./...

run:
	go run main.go
