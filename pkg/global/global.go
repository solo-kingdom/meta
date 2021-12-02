package global

import "github.com/dgraph-io/badger/v3"

type Vars struct {
	KV *badger.DB
}

var GV = &Vars{}
