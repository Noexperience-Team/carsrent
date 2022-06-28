FROM golang:1.16-alpine as backend
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 go build -o ./carRent ./src
EXPOSE 8888
entrypoint [ "./carRent" ]
