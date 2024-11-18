FROM golang:1.23-alpine

WORKDIR /app
COPY . ./
RUN go build main.go

CMD [ "/app/main" ]