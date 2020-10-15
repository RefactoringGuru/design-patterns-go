package main

type evictionAlgo interface {
	evict(c *cache)
}
