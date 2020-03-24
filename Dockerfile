FROM golang:1.13-alpine
RUN mkdir /app
WORKDIR /app
COPY . .
RUN go build -o authservice
EXPOSE 1323
CMD ["./authservice"]

