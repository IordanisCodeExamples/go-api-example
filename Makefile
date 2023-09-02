build:
	go build app.go
run:
	go run app.go
compile:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=arm go build -o bin/main-linux-arm app.go
	GOOS=linux GOARCH=arm64 go build -o bin/main-linux-arm64 app.go
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 app.go
test:
	go test ./...
test-conver:
	go test -cover ./...
generate:
	go generate -x ./...
get-generator:
	go install github.com/golang/mock/mockgen
integration:
	docker-compose down --rmi local
	docker-compose build --parallel
	docker-compose up -d --force-recreate
	go test -v -vet all -tags=integration ./internal/persistence/mongo
lint:
	golangci-lint run --print-issued-lines=false --out-format code-climate:gl-code-quality-report.json,line-number