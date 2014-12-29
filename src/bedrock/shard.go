package bedrock

import (
	"errors"
	"bytes"
	"sync"
)

var (
	NotFound = errors.New("Not Found")
)

// shard index entry
// implements comparable
type IndexEntry struct {
	Key      [32]byte
	Refcount uint64
	Size     uint64
	Offset   uint64
}

// a shard index
type ShardIndex struct {
	// sorted entry set
	Version uint64
	Entries []IndexEntry
}

// shard
type Shard struct {
	sync.Mutex
	Index ShardIndex
}

func (self *Shard) Get(key [32]byte) ([]byte, error) {
	self.RLock()
	defer self.Runlock()
	kslice := []byte(key)
	for i := range self.Index.Entries {
		if bytes.Equal(kslice, []byte(self.Index.Entries[i].Key)) {

		}
	}
	return nil, NotFound
}

func (self *Shard) Put(key [32]byte, data []byte) error {
}

func (self *Shard) Delete(key [32]byte) error {
}

func (self *Shard) Inc(key [32]byte) error {
}

func (self *Sahrd) Dec(key [32]byte) error {
}

// remove any tombstoned entries from index and data file
func (self *Shard) Compact() error {
	return errors.New("Shard.Compact NOT IMPLEMENTED!")
}
