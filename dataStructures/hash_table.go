package dataStructures

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"hash/fnv"
	"reflect"
)

const defaultLoadFactor = 10

type HashTable[T comparable] interface {
	Add(key string, data T) error
	Get(key string) (T, bool, error)
	Size() int
}

type BucketItem[T comparable] interface {
	Data() T
	Key() string
}

func NewHashTable[T comparable]() HashTable[T] {
	table := &hashTable[T]{
		data: make([]LinkedList[BucketItem[T]], defaultLoadFactor),
	}

	for i, _ := range table.data {
		table.data[i] = NewLinkedList[BucketItem[T]]()
	}

	return table
}

type hashTable[T comparable] struct {
	data []LinkedList[BucketItem[T]]
}

type bucketItem[T comparable] struct {
	data T
	key  string
}

func (b *bucketItem[T]) Data() T {
	return b.data
}

func (b *bucketItem[T]) Key() string {
	return b.key
}

func (r *hashTable[T]) Size() int {
	size := 0
	for _, v := range r.data {
		size += v.Size()
	}
	return size
}

func (r *hashTable[T]) Add(key string, data T) error {
	index, err := r.hash(key)
	if err != nil {
		return err
	}
	var item BucketItem[T] = &bucketItem[T]{data: data, key: key}
	r.data[index].Add(item)
	return nil
}

func (r *hashTable[T]) Get(key string) (T, bool, error) {
	var emptyResult T
	index, err := r.hash(key)
	if err != nil {
		return emptyResult, false, err
	}

	current := r.data[index].Head()
	for !isNil(current) {
		value := current.Value()
		if !isNil(value) && value.Key() == key {
			return value.Data(), true, nil
		}
		current = current.Next()
	}

	return emptyResult, false, fmt.Errorf("key %v not found", key)
}

// to handle interfaces and other built in types
func isNil[T any](t T) bool {
	v := reflect.ValueOf(t)
	kind := v.Kind()
	// Must be one of these types to be nillable
	return (kind == reflect.Ptr ||
		kind == reflect.Interface ||
		kind == reflect.Slice ||
		kind == reflect.Map ||
		kind == reflect.Chan ||
		kind == reflect.Func) &&
		v.IsNil()
}

func (r *hashTable[T]) hash(data string) (uint32, error) {
	buffer := bytes.Buffer{}

	//we need a fixed size data here, so converting to by array slice
	slice := []byte(data)

	err := binary.Write(&buffer, binary.LittleEndian, slice)
	if err != nil {
		return 0, fmt.Errorf("failed to write binary of value %v - %s", data, err.Error())
	}
	h := fnv.New32a()
	_, err = h.Write(buffer.Bytes())
	if err != nil {
		return 0, fmt.Errorf("failed to hash value %v - %s", data, err.Error())
	}
	return h.Sum32() % uint32(len(r.data)), nil
}
