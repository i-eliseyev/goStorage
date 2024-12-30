package tests

import (
	"errors"
	"goStorage/internal/storage"
	"testing"
)

func TestPut(t *testing.T) {
	const key = "create-key"
	const value = "create-value"

	var val interface{}
	var contains bool

	defer delete(storage.Store.M, key)

	// Sanity check
	_, contains = storage.Store.M[key]
	if contains {
		t.Error("key/value already exists")
	}

	// err should be nil
	err := storage.Put(key, value)
	if err != nil {
		t.Error(err)
	}

	val, contains = storage.Store.M[key]
	if !contains {
		t.Error("create failed")
	}

	if val != value {
		t.Error("val/value mismatch")
	}
}

func TestGet(t *testing.T) {
	const key = "read-key"
	const value = "read-value"

	var val interface{}
	var err error

	defer delete(storage.Store.M, key)

	// Read a non-thing
	val, err = storage.Get(key)
	if err == nil {
		t.Error("expected an error")
	}
	if !errors.Is(err, storage.ErrorNoSuchKey) {
		t.Error("unexpected error:", err)
	}

	storage.Store.M[key] = value

	val, err = storage.Get(key)
	if err != nil {
		t.Error("unexpected error:", err)
	}

	if val != value {
		t.Error("val/value mismatch")
	}
}

func TestDelete(t *testing.T) {
	const key = "delete-key"
	const value = "delete-value"

	var contains bool

	defer delete(storage.Store.M, key)

	storage.Store.M[key] = value

	_, contains = storage.Store.M[key]
	if !contains {
		t.Error("key/value doesn't exist")
	}

	storage.Delete(key)

	_, contains = storage.Store.M[key]
	if contains {
		t.Error("Delete failed")
	}
}
