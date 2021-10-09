.DEFAULT_TARGET=build
URL := http://localhost:8080

run:
	go run cmd/*.go

build:
	go build cmd/*.go

test:
	curl -X POST ${URL}/add -d @utils/test.json \
	--header "Content-Type: application/json"
