FROM golang:1.19-alpine3.16 as build

RUN apk update 

WORKDIR /app

COPY go.mod /app/
COPY go.sum /app/

RUN go mod download
RUN go mod tidy

COPY . /app/

RUN go build -o /app/main

# --------

FROM alpine:3.16

WORKDIR /app

COPY --from=build /app/main /app/main

CMD ["./main"]