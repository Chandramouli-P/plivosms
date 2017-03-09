.PHONY: build watch

NAME = testproject
LDFLAGS = -ldflags "-X main.Version=${APP_VERSION}"

build:
	go build -o ${NAME} main.go

run: build
	./${NAME}

watch:
	reflex -r '\.(go)$$' -R '^vendor/' -s -- sh -c 'go build -o ${NAME} main.go && ./${NAME}'

cover:
	go test -coverprofile=coverage.out testproject/handler
	go test -coverprofile=coverage.out testproject/service

cover_show:
	go tool cover -html=coverage.out -o=coverage.html
