generate-mocks:
	go install github.com/vektra/mockery/v2@v2.53.3
	mockery --dir src/core/_interfaces/ --all --output src/core/_mocks
