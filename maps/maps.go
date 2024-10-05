package maps

import (
	"errors"
)

type Dictionary map[string]string

var (
	errorKeyExists = errors.New("Key already exists in the list!.")
	errorNotExists = errors.New("Key doesn't exist")
)

func (d Dictionary) Search(key string) (string, error) {
	defination, ok := d[key]
	if !ok {
		return "", errorNotExists
	}
	return defination, nil
}

func (d Dictionary) Add(key, value string) error {
	got, ok := d.Search(key)
	if ok == errorNotExists {
		d[key] = value
	}
	if got != "" {
		return errorKeyExists
	}
	return nil
}

func (d Dictionary) Remove(key string) error {
	delete(d, key)
	return nil
}
