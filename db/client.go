package db

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

type Items map[string]string

type DbClient interface {
	Initialize() error
	SetValue(key string, value string) error
	GetValue(key string) ([]byte, error)
	GetAll() Items
}

type BoltDB struct {
	boltDB *bolt.DB
}

// Initialize the db
func (bc *BoltDB) Initialize() error {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	var err error
	bc.boltDB, err = bolt.Open("key-value.db", 0600, nil)
	// defer bc.boltDB.Close()
	if err != nil {
		log.Fatal(err)
		return err
	}

	// Create a bucket
	bc.boltDB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("StoreBucket"))
		if err != nil {
			return fmt.Errorf("[ERROR] Create bucket failed: %s", err)
		}
		return nil
	})

	return nil
}

// SetValue create a key value pair
func (bc *BoltDB) SetValue(key string, value string) error {
	err := bc.boltDB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("StoreBucket"))
		err := b.Put([]byte(key), []byte(value))
		return err
	})

	return err
}

// GetValue return the value of a specified key
func (bc *BoltDB) GetValue(key string) ([]byte, error) {
	var a []byte

	err := bc.boltDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("StoreBucket"))
		a = b.Get([]byte(key))
		return nil
	})

	if len(a) == 0 {
		a := []byte("No such key found")
		return a, err
	}

	return a, err
}

// Items is all the items in the db

// GetAll returns all  key value pairs in the db
func (bc *BoltDB) GetAll() Items {
	items := make(map[string]string)
	bc.boltDB.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte("StoreBucket"))

		b.ForEach(func(k, v []byte) error {
			items[string(k)] = string(v)
			return nil
		})
		return nil
	})

	return items
}
