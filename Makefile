test:
	go test -v ./... -race -shuffle=on

bench:
	go test -bench=. -benchmem ./...