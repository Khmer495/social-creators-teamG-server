# social-creators-teamG-server

## 使用方法
Makefileとdocker-composeを使用  
```
Makfile
up:
	docker-compose up
down:
	docker-compose down
build:
	docker-compose build
```

## 使用言語・ライブラリ
- Go
  - [echo](https://github.com/labstack/echo)
  - [gorm](https://github.com/go-gorm/gorm)

## ディレクトリ構成
./go配下はクリーンアーキテクチャを採用  
```
.
├── Makefile
├── README.md
├── api
│   └── openapi.yaml
├── database
│   ├── data
│   └── my.cnf
├── docker-compose.yml
├── docs
│   └── openapi_viewer
└── go
    ├── Dockerfile
    ├── entity
    │   └── model
    ├── go.mod
    ├── go.sum
    ├── infrastracture
    ├── interface
    │   ├── controller
    │   └── database
    ├── main.go
    └── usecase
```
参考：
- [Clean ArchitectureでAPI Serverを構築してみる](https://qiita.com/hirotakan/items/698c1f5773a3cca6193e)
- [Golang - EchoとGORMでClean Architecture APIを構築する](https://qiita.com/so-heee/items/0cca93008eae635c642a)
