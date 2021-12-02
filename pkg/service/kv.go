package service

import (
	"errors"
	"github.com/dgraph-io/badger/v3"
	"github.com/solo-kingdom/meta/pkg/global"
)

func Get(key string) (bool, string, error) {
	var res *[]byte
	err := global.GV.KV.View(func(txn *badger.Txn) error {
		var err error
		if item, err := txn.Get([]byte(key)); err == nil {
			return item.Value(func(val []byte) error {
				res = &val
				return nil
			})
		}
		return err
	})
	if err != nil {
		return false, "", err
	}
	if res == nil {
		return false, "", errors.New("")
	}

	return true, string((*res)[:]), err
}

func Set(key string, value string) error {
	return global.GV.KV.Update(func(txn *badger.Txn) error {
		e := badger.NewEntry([]byte(key), []byte(value))
		err := txn.SetEntry(e)
		return err
	})
}
