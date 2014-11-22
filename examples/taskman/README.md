TaskMan
================

Install Dependencies
========
```go
go get github.com/syndtr/goleveldb/leveldb
go get github.com/jteeuwen/go-bindata/...
go get github.com/elazarl/go-bindata-assetfs
go get github.com/go-martini/martini
```

Build assets 
======
```go
go-bindata data/...
```

run main
======
```go
go run main.go bindata.go
```