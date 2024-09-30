test:
	go test -v ./... -race -shuffle=on

# Earlier versions of Go do not support the -shuffle flag
test-no-shuffle:
	go test ./... -race

bench:
	go test -bench=. -benchmem ./...

test-cover:
	go test -race -shuffle=on -coverprofile=coverage.txt ./... && go tool cover -html=coverage.txt