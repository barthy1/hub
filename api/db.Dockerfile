FROM golang:1.15-alpine3.12 AS builder

WORKDIR /go/src/github.com/tektoncd/hub/api
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=s390x go build -o db-migration ./cmd/db/...

FROM alpine:3.12

RUN apk --no-cache add ca-certificates && addgroup -S hub && adduser -S hub -G hub
USER hub

WORKDIR /app
COPY --from=builder /go/src/github.com/tektoncd/hub/api/db-migration /app/db-migration
EXPOSE 8000

CMD [ "/app/db-migration" ]
