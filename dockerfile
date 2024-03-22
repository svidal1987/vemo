FROM golang:1.22.1 as builder
WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -o server
EXPOSE 8080:8080
CMD [ "./server" ]

