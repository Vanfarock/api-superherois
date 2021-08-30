FROM golang:1.17-alpine3.14
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o /app/api-superherois

EXPOSE 3333

CMD ["./api-superherois"]