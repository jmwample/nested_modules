package client

import (
	"crypto/sha256"
	"fmt"
)

type Client interface {
	Function() error
}

type BasicClient struct {
	hashfn func([]byte) [32]byte
}

func New() *BasicClient {
	return &BasicClient{hashfn: sha256.Sum256}
}

func (c *BasicClient) Function() error {

	sum := c.hashfn([]byte("hello world\n"))
	fmt.Printf("%x", sum)
	return nil
}
