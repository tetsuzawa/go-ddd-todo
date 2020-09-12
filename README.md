# go-ddd-todo

GoのDDDの実装を試すリポジトリ

## 構造

```
$ tree
.
├── README.md
├── application
│   └── usecase
│       └── todo_usecase.go
├── cmd
│   └── api
│       └── main.go
├── domain
│   ├── model
│   │   └── todo_model.go
│   ├── repository
│   │   └── todo_repository.go
│   └── service
│       └── todo_service.go
├── go.mod
├── go.sum
├── infrastructure
│   └── persistence
│       └── datastore
│           └── todo_repository.go
├── interfaces
│   └── api
│       └── server
│           ├── auth
│           ├── dbutil
│           │   └── db.go
│           ├── handler
│           │   ├── app_handler.go
│           │   └── todo_handler.go
│           ├── httputil
│           │   ├── error.go
│           │   └── response_json.go
│           ├── middleware
│           │   ├── log.go
│           │   └── recover.go
│           ├── router
│           │   └── router.go
│           └── server.go
└── registry
    └── registry.go

21 directories, 19 files
```

## 参考

- https://engineer.recruit-lifestyle.co.jp/techblog/2018-03-16-go-ddd/
- https://eng-blog.iij.ad.jp/archives/2442


