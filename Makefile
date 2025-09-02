vet:
	go vet ./...

staticcheck:
	staticcheck ./...

test: vet staticcheck
