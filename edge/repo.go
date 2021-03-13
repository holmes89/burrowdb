package edge

import (
	"github.com/boltdb/bolt"
	"github.com/holmes89/burrowdb"
)

type edgeRepo struct {
	db *bolt.DB
}

func (r *edgeRepo) Get(id string) (*burrowdb.Edge, error) {
	return nil, nil
}

func (r *edgeRepo) Put(edge *burrowdb.Edge) error {
	return nil
}

func (r *edgeRepo) Delete(id string) error {
	return nil
}
