APP_EXECUTABLE?=./bin/crawler
RELEASE?=1.0
MIGRATIONS_RELEASE?=1.0
IMAGENAME?=mikhailmi/crawler-service:v$(RELEASE)

.PHONY: clean
clean:
	rm -f ${APP_EXECUTABLE}

.PHONY: docker-build
docker-build: clean
	docker build -t $(MIGRATIONS_IMAGENAME) -f DockerfileMigrations .
	docker build -t $(IMAGENAME) .

.PHONY: build
build:
	go build -o bin\crawler.exe cmd\main.go

.PHONY: stop
stop:
	taskkill /F /IM crawler.exe

.PHONY: release
release:
	git tag v$(RELEASE)
	git push origin v$(RELEASE)