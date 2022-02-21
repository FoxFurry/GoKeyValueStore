package command

import (
	"errors"
	"github.com/google/btree"
)

type get struct {
	pair
}

func (c *get) Execute(storage *btree.BTree) (string, error) {
	result := storage.Get(c.pair)

	if result == nil {
		return "", nil
	}

	resPair, ok := result.(pair)
	if !ok {
		return "", errors.New("failed to cast result")
	}

	return resPair.value, nil
}

func (cmd *get) Key() string {
	return cmd.key
}

func (cmd *get) Type() CommandType {
	return Write
}

func Get(key string) Command {
	return &get{
		pair: pair{
			key: key,
		},
	}
}
