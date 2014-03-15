package dpkgdb

import (
	"sync"
)

type DB struct {
	sync.RWMutex

	archs      []string
	archsIndex map[string]struct{}

	packages            []*packageT
	packagesIndexByName map[string]*packageT
}

func newDB() *DB {
	db := new(DB)
	db.archs = make([]string, 0, 2)
	db.archsIndex = make(map[string]struct{})

	db.packages = make([]*packageT, 0, 64)
	db.packagesIndexByName = make(map[string]*packageT)

	return db
}
