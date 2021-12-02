package service

import (
	"github.com/dgraph-io/badger/v3"
	"github.com/solo-kingdom/meta/pkg/global"
)

func Get(key string) (string, error) {
	var res []byte
	err := global.GV.KV.View(func(txn *badger.Txn) error {
		item, _ := txn.Get([]byte(key))
		_ = item.Value(func(val []byte) error {
			res = val
			return nil
		})
		return nil
	})
	return string(res[:]), err
}

func Set(key string, value string) error {
	return global.GV.KV.Update(func(txn *badger.Txn) error {
		e := badger.NewEntry([]byte(key), []byte(value))
		err := txn.SetEntry(e)
		return err
	})
}
