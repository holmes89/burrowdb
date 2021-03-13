package node

import (
	"encoding/json"

	"github.com/boltdb/bolt"
	"github.com/holmes89/burrowdb"
)

var nodeBucket = []byte("nodes")

type nodeRepo struct {
	db *bolt.DB
}

func (r *nodeRepo) Get(id string) (*burrowdb.Node, error) {
	var node burrowdb.Node
	r.db.View(func(tx *bolt.Tx) error {
		res := tx.Bucket(nodeBucket).Get([]byte(id))
		return json.Unmarshal(res, &node)
	})
	return &node, nil
}

func (r *nodeRepo) Put(node *burrowdb.Node) error {
	return r.db.Update(func(tx *bolt.Tx) error {
		n, _ := json.Marshal(node)
		return tx.Bucket(nodeBucket).Put([]byte(node.ID), n)
	})
}

func (r *nodeRepo) Delete(id string) error {
	return r.db.Update(func(tx *bolt.Tx) error {
		return tx.Bucket(nodeBucket).Delete([]byte(id))
	})
}
