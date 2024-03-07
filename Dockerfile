FROM golang:alpine3.19 as build

WORKDIR /server

COPY . .

RUN go mod download

RUN go build -o ./investment_service cmd/main.go

FROM alpine:3.19

COPY --from=build /server/investment_service ./investment_service

COPY --from=build /server/.env ./.env

ENTRYPOINT ["./investment_service"]