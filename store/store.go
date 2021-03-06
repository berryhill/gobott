package store

import (
	"fmt"
	"log"
	"encoding/binary"

	"github.com/boltdb/bolt"
)

func openDb() (*bolt.DB, error) {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	return db, err
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))

	return b
}

func InitDb() error {
	db, err := openDb()
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("machine"))

		return err
	})
	if err != nil {
		fmt.Println(err)
	}

	return err
}

func AddToDb(bucket []byte, key []byte, value []byte) error {
	db, err := openDb()
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(bucket)
		if err != nil {
			return err
		}

		err = bucket.Put(key, value)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
	if err = db.Close(); err != nil {
		log.Fatal(err)
	}

	return err
}

func RetrieveFromDb(bucket []byte, key []byte) ([]byte, error) {
	db, err := openDb()
	var val []byte
	defer db.Close()

	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucket)
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found!", key)
		}

		val := bucket.Get(key)
		fmt.Println(string(val))

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
	if err = db.Close(); err != nil {
		log.Fatal(err)
	}

	return val, err
}

func RetrieveAllFromDb(model interface{}, bucket []byte) /*(map[string][]byte, error)*/ error {
	db, err := openDb()
	defer db.Close()

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		c := b.Cursor()

		//var list []interface{}
		for k, v := c.First(); k != nil; k, v = c.Next() {
			//retrieved := new(model)
			//list = append(list, )

			fmt.Printf("key=%s, value=%s\n", k, v)
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
	if err = db.Close(); err != nil {
		log.Fatal(err)
	}

	return nil
}

func UpdateDb(bucket []byte, key []byte, data []byte) error {
	db, err := openDb()
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		err = b.Put(key, data)
		return err
	})

	if err != nil {
		log.Fatal(err)
	}
	if err = db.Close(); err != nil {
		log.Fatal(err)
	}

	return nil
}

func DeleteBucket(bucket []byte) error {
	//TODO implement
	return nil
}
