FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY ./main.go .
COPY api .
COPY db .
COPY util .
COPY .env .
COPY app.env .

# RUN go get -d -v ./...
# RUN go install -v ./...

RUN CGO_ENABLED=0 GOOS=linux go build -o /weather_api ./main.go

FROM alpine
WORKDIR /app
COPY --from=builder /weather_api .
COPY app.env .
COPY .env .

EXPOSE 8080

CMD [ "/app/weather_api" ]