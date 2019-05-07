package db

import (
	"fmt"

	"github.com/boltdb/bolt"
)

// DataStore defines the database functions
type DataStore interface {
	CreateBuckets(buckets string) error
	AddItem(key, value, bucket string) error
	RemoveItem(key, bucket string) error
	ListItems(bucket string) ([]string, error)
}

// Store represents data storage and functions.
type Store struct {
	DB *bolt.DB
}

// Connect provides a store with a bolt database instance
func Connect(name string) (Store, error) {
	botDB, err := bolt.Open(name, 0600, nil)

	if err != nil {
		return Store{}, nil
	}
	return Store{DB: botDB}, nil
}

// Close the database connection
func (s Store) Close() error {
	if err := s.DB.Close(); err != nil {
		return err
	}
	return nil
}

// CreateBuckets should be used at the start of a connection to create all
// buckets necessary for db operations.
func (s Store) CreateBuckets(buckets []string) error {
	return s.DB.Update(func(tx *bolt.Tx) error {
		for _, bucket := range buckets {
			_, err := tx.CreateBucketIfNotExists([]byte(bucket))
			if err != nil {
				return fmt.Errorf("unable to create bucket %s: %s", bucket, err)
			}
		}
		return nil
	})
}

// AddItem takes a key, value, and bucket and adds it to the db.
func (s Store) AddItem(key, value, bucket string) error {
	return s.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		err := b.Put([]byte(key), []byte("42"))
		return err
	})
}

// RemoveItem takes a key and bucket and removes it from the db.
func (s Store) RemoveItem(key, bucket string) error {
	return s.DB.Update(func(tx *bolt.Tx) error {
		return tx.Bucket([]byte(bucket)).Delete([]byte(key))
	})
}

// ListItems returns all items from a bucket.
func (s Store) ListItems(bucket string) ([]string, error) {
	var keyList []string
	if err := s.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if err := b.ForEach(func(k, v []byte) error {
			keyList = append(keyList, string(k))
			return nil
		}); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return keyList, err
	}
	return keyList, nil
}
