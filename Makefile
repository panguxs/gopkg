mysql-restart:
	@docker rm -f mysql && ./scripts/dev/mysql_start.sh
migrate:
	@echo "Run migrations"
	@cd scripts/sql && go-bindata -pkg migrations -o ../../pkg/migrations/bindata.go . && cd ../../ \
	&& PG_CONF_FILE=$(PWD)/configs/app.example.yml \
	PG_TEST_DATA_FILE=$(PWD)/scripts/testdata/data.sql \
	go test -v ./pkg/migrations/.	
gopkg-tmpl:
	@echo "Run bindata gopkg tmpl"
	@cd tmpl && go-bindata -pkg tmpl -o ../../pkg/apiserver/gopkg/tmpl/bindata.go .
run:
	@echo "Run server"
	@PG_CONF_FILE=$(PWD)"/configs/app.example.yml" \
	PG_TEST_DATA_FILE=$(PWD)/scripts/testdata/data.sql \
	go run cmd/gopkg/main.go
test: 
	@PG_CONF_FILE=$(PWD)/configs/app.example.yml \
	go test -count=1 $(shell go list ./... | grep -v 'pgxs.io/gopkg/pkg/migrations')
test-all: migrate test
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags netgo -a -installsuffix cgo -o gopkg ./cmd/gopkg/...
docker-build:
	@docker build --no-cache -f docker/Dockerfile -t pgxs/gopkg:0.1.1 .