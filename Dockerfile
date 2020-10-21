FROM golang:1.15.2-alpine

WORKDIR /go/src/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# RUN go get -d -v ./...
# RUN go install -v ./...
# RUN go get -u github.com/labstack/echo/... \
#     && go get github.com/jinzhu/gorm \
#     && go get github.com/go-sql-driver/mysql

# RUN apt-get update -qq && apt-get install -y mariadb-client
# RUN chmod a+x ./docker-entrypoint.sh
# ENTRYPOINT ["./docker-entrypoint.sh"]

CMD ["go", "run", "main.go"]
