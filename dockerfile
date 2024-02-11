FROM golang:alpine
WORKDIR /app
ADD . .
RUN go install
CMD go run .
