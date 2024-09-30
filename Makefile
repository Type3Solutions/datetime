test:
	go test -v ./... -race -shuffle=on

bench:
	go test -bench=. -benchmem ./...

test-cover:
	go test -race -shuffle=on -coverprofile=coverage.txt ./... && go tool cover -html=coverage.txt