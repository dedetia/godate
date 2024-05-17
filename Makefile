BINARY=godate

init:
	go mod tidy
	go mod vendor

test:
	go test -v -cover -count=1 -failfast ./... -coverprofile="coverage.out"

test-html:
	 go tool cover -html=coverage.out -o coverage.html

test-html-output: test test-html

build:
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY} -mod=vendor -a -installsuffix cgo -ldflags '-w'

run:
	docker compose up --build -d

run-lint:
	docker compose up --build -d lint

stop:
	docker compose down

dependency:
	@echo "> Installing the server dependencies ..."
	@go mod vendor

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

mock: mock-repository

mock-repository:
	mockgen -source=./internal/core/port/registry/repository.go -destination=./shared/mock/repository/repository_registry_mock.go -package repository
	mockgen -source=./internal/core/port/repository/swipe.go -destination=./shared/mock/repository/swipe_mock.go -package repository
	mockgen -source=./internal/core/port/repository/user.go -destination=./shared/mock/repository/user_mock.go -package repository

.PHONY: clean build test mock