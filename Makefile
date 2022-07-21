GOBUILD=go build
GOTEST=go test

all:clean stop build_app
	./token_app
build_app: 
	$(GOBUILD) -v .
clean:
	rm -f token_app
stop:
	pkill token_app || true
