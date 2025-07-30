# App binary name
APP_NAME=app
PORT=8080

build:
	go build -o $(APP_NAME) ./cmd/app

run: build
	./$(APP_NAME)

test:
	go test ./...

clean:
	rm -f $(APP_NAME)

docker-build:
	docker build -t go-template .

docker-run:
	docker run -p $(PORT):$(PORT) go-template

deploy:
	gcloud run deploy go-template \
		--source=. \
		--region=us-central1 \
		--platform=managed \
		--allow-unauthenticated \
		--set-env-vars=GO_ENV=production,PORT=$(PORT)
