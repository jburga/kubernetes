FROM golang:1.16-alpine AS BUILDER
WORKDIR /app

COPY main.go ./

RUN CGO_ENABLE=0 GO111MODULE=off GOOS=linux go build -o /app/main main.go

FROM alpine
WORKDIR /app
COPY --from=BUILDER /app/main .
ENV PORT=80
CMD [ "/app/main" ]