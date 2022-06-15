FROM golang:1.18.3

# copy the source code to the container
COPY /cmd/main.go /app/main.go

# runnning golang app in the container "go run /app/main.go"
CMD [ "go", "run", "/app/main.go" ]