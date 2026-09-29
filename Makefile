cover-tests:
	@. ./.test.env && go clean -testcache && go test -cover -race ./...

tests:
	@. ./.test.env && go clean -testcache && go test -v -race ./...

%::
	@true