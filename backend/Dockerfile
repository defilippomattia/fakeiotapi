FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .
#look in db.go for clarification
ENV FRA_DB_HOST=db-container
ENV FRA_DB_PORT=5432
ENV FRA_DB_USER=fakeiot
ENV FRA_DB_PASSWORD=fakeiot
ENV FRA_DB_NAME=fakeiot

CMD ["./main"]
