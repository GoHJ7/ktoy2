FROM golang:alpine as builder
WORKDIR /app
COPY . .
RUN go build -o scrapper .

FROM alpine
COPY --from=builder /app/scrapper /scrapper
CMD ["/scrapper"]

