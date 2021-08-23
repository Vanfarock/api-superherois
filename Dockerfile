FROM golang:alpine3.9
WORKDIR /app
COPY . .

RUN ./download-dependencies.sh

CMD ["go", "run", "main.go"]