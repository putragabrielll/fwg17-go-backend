FROM golang

# bebas nama folder di docker
WORKDIR /go-coffee-shop 

COPY . .

RUN cp .env.example .env
RUN go mod tidy

EXPOSE 8080

CMD go run .