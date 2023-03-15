FROM golang:1.19-buster as builder

WORKDIR /go/src/app

COPY . .

RUN go build -v -o migrations main.go

FROM builder as prod

WORKDIR /bin

COPY --from=builder /go/src/app/migrations .

CMD migrations

