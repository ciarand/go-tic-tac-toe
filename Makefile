test:
	go test -v ./...

coverage:
	cd lib && go test -coverprofile=coverage.out && \
		go tool cover -html=coverage.out && rm coverage.out
