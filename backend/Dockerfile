FROM golang:1.16.4
WORKDIR /app
COPY ./ ./
RUN go mod download
CMD [ "go", "run", "main.go" ]