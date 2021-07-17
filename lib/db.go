package lib

import "github.com/syndtr/goleveldb/leveldb"

const (
	DBSuffix = "_akakun"
)

type DB struct {
	*leveldb.DB
}
