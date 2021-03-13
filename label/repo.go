package label

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/holmes89/burrowdb"
)

var labelBucket = []byte("labels")

type labelRepo struct {
	db *bolt.DB
}

func (r *labelRepo) Search(label string, opts map[string]interface{}) ([]string, error) {
	// find all buckets with label prefix
	return nil, nil
}

func (r *labelRepo) Put(label burrowdb.Label, node burrowdb.Node) error {
	return r.db.Update(func(tx *bolt.Tx) error {
		k := key(label, node) // should this be another nested bucket? Then all sub values would be there... just iterate.
		b, err := tx.Bucket(labelBucket).CreateBucketIfNotExists(k)
		if err != nil {
			return err
		}
		for key, v := range node.Properties {
			var buf bytes.Buffer
			enc := gob.NewEncoder(&buf)
			err := enc.Encode(v)
			if err != nil {
				return err
			}
			b.Put([]byte(key), buf.Bytes())
		}
		return nil
	})
}

func key(label burrowdb.Label, node burrowdb.Node) []byte {
	l := label
	if l == "" {
		l = "nil"
	}
	s := fmt.Sprintf("%s:%s", l, node.ID)
	return []byte(s)
}
