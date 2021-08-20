package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

var (
	errNotFound    = errors.New("not found")
	errKeyExist    = errors.New("key already exists")
	errKeyNotExist = errors.New("key not exists")
)

func (d Dictionary) Search(key string) (string, error) {
	value, exists := d[key]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

func (d Dictionary) Add(key string, value string) error {
	_, e := d.Search(key)

	if e == errNotFound {
		d[key] = value
		return nil
	} else {
		return errKeyExist
	}
}

func (d Dictionary) Update(key string, value string) error {
	_, e := d.Search(key)

	if e == errNotFound {
		return errKeyNotExist
	} else {
		d[key] = value
		return nil
	}
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}

/*

package main

import (
	"fmt"
	"go_study/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"First": "A"}

	// value, err := dictionary.Search("First")
	// fmt.Println(value, err)

	// value, err = dictionary.Search("Second")
	// fmt.Println(value, err)

	// err = dictionary.Add("Second", "B")
	// value, err = dictionary.Search("Second")
	// fmt.Println(value, err)

	err := dictionary.Update("First", "B")
	fmt.Println(err)

	err = dictionary.Update("Second", "B")
	fmt.Println(err)

	fmt.Println(dictionary)
}
*/
