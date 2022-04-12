package db

import (
	"encoding/binary"
	"encoding/json"
	"os"
	"time"

	"github.com/boltdb/bolt"
)

var bucketName = []byte("tasks")

type Task struct {
	ID          int
	TimeAdded   time.Time
	Value       string
	Completed   bool
	CompletedAt time.Time
}

type DB struct {
	db *bolt.DB
}

func InitDatabase(dbPath string) *DB {
	db, err := bolt.Open(dbPath, os.ModePerm, nil)
	if err != nil {
		panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bucketName)
		return err
	})
	if err != nil {
		panic(err)
	}

	return &DB{
		db: db,
	}
}

func (d *DB) All() ([]Task, error) {
	var tasks []Task
	err := d.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			var task Task
			if err := json.Unmarshal(v, &task); err != nil {
				return err
			}
			tasks = append(tasks, task)
		}
		return nil
	})
	return tasks, err
}

func (d *DB) Save(task string) error {
	return d.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		id, _ := b.NextSequence()
		t := Task{
			ID:        int(id),
			TimeAdded: time.Now(),
			Value:     task,
			Completed: false,
		}
		marshalled, err := json.Marshal(t)
		if err != nil {
			return err
		}
		return b.Put(itob(t.ID), marshalled)
	})
}

func (d *DB) MarkComplete(id int) error {
	return d.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)
		var task Task
		unmarshalled := b.Get(itob(id))
		if err := json.Unmarshal(unmarshalled, &task); err != nil {
			return err
		}
		task.Completed = true
		task.CompletedAt = time.Now()
		marshalled, err := json.Marshal(&task)
		if err != nil {
			return err
		}
		return b.Put(itob(task.ID), marshalled)
	})
}

func (d *DB) Remove(id int) error {
	return d.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)
		return b.Delete(itob(id))
	})
}

func (d *DB) Last24HTasks() ([]Task, error) {
	var tasks []Task
	err := d.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)
		c := b.Cursor()

		t := time.Now().Add(-24 * time.Hour)
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var task Task
			if err := json.Unmarshal(v, &task); err != nil {
				return err
			}
			if task.TimeAdded.After(t) {
				tasks = append(tasks, task)
			}
		}
		return nil
	})
	return tasks, err
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
