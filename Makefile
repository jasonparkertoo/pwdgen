
BINARY_NAME=pwdgen

build:
	GOARCH=amd64 GOOS=linux go build -o ./$(OUTDIR)/$(BINARY_NAME) .

clean:
	go clean
	test -e ./$(OUTDIR)/$(BINARY_NAME) && rm ./$(OUTDIR)/$(BINARY_NAME)

test:
	go test