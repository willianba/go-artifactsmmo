tests:
	go test ./actions

coverage:
	go test ./actions -coverprofile coverage.out .
	go tool cover -html coverage.out -o coverage.html
	rm coverage.out

