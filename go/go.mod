module github.com/Khmer495/social-creators-teamG-server/go

go 1.15

require (
	github.com/jinzhu/gorm v1.9.16
	github.com/labstack/echo/v4 v4.1.17
	github.com/mattn/go-sqlite3 v2.0.1+incompatible // indirect
)

replace gopkg.in/urfave/cli.v2 => github.com/urfave/cli/v2 v2.2.0 // Need to use github.com/oxequa/realize, used in ./Dockerfile
