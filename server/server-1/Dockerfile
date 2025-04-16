FROM golang:1.21-alpine
WORKDIR /app
COPY . .
RUN go build -o server .
EXPOSE 8000
CMD ["./server"]
