FROM golang:1.22 as builder

WORKDIR /app

COPY ./app .

COPY ./app/go.mod ./app/go.sum ./
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o butte

FROM alpine:3.7 as runner

# binary
COPY --from=builder /app/butte /butte

EXPOSE 80
ENV ENV=production

CMD ["/butte"]
