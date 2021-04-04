package burrowdb

import (
	"encoding/json"
	"errors"

	"github.com/boltdb/bolt"
	"github.com/google/uuid"
)

var (
	nodeBucket  = []byte("nodes")
	indexBucket = []byte("index")
)

type Repository struct {
	db *bolt.DB
}

func NewRepository(db *bolt.DB) *Repository {
	if err := createNodeBucket(db); err != nil {
		panic("unable to create bucket")
	}
	return &Repository{
		db: db,
	}
}

func createNodeBucket(db *bolt.DB) error {
	tx, err := db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Use the transaction...
	if _, err := tx.CreateBucketIfNotExists(nodeBucket); err != nil {
		return err
	}

	if _, err := tx.CreateBucketIfNotExists(indexBucket); err != nil {
		return err
	}

	// Commit the transaction and check for error.
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (s *Repository) Create(nodes ...*Node) error {
	return s.db.Batch(func(tx *bolt.Tx) error {
		for _, node := range nodes {
			node.ID = uuid.New().String()
			n, err := json.Marshal(node)
			if err != nil {
				return errors.New("unable to marshal node")
			}
			if err := tx.Bucket(nodeBucket).Put([]byte(node.ID), n); err != nil {
				return errors.New("unable to insert node")
			}
			for _, l := range node.Labels {
				ib, err := tx.Bucket(indexBucket).CreateBucketIfNotExists([]byte(l))
				if err != nil {
					return errors.New("unable create label container")
				}
				if err := ib.Put([]byte(node.ID), n); err != nil {
					return errors.New("unable store label")
				}
			}
		}
		return nil
	})
}
