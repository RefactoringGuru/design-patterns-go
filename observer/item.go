package main

import "fmt"

type Item struct {
	observerList []Observer
	name         string
	inStock      bool
}

func newItem(name string) *Item {
	return &Item{
		name: name,
	}
}
func (i *Item) updateAvailability() {
	fmt.Printf("Item %s is now in stock\n", i.name)
	i.inStock = true
	i.notifyAll()
}
func (i *Item) register(o Observer) {
	i.observerList = append(i.observerList, o)
}

func (i *Item) deregister(o Observer) {
	i.observerList = removeFromSlice(i.observerList, o)
}

func (i *Item) notifyAll() {
	for _, observer := range i.observerList {
		observer.update(i.name)
	}
}

func removeFromSlice(observerList []Observer, observerToRemove Observer) []Observer {
	idx := findIndex(observerList, observerToRemove)
	if idx == -1 {
		return observerList
	}
	observerList = append(observerList[:idx], observerList[idx+1:]...)
	return observerList
}

func findIndex(observerList []Observer, observerToRemove Observer) int {
	for i, observer := range observerList {
		if observer.getID() == observerToRemove.getID() {
			return i
		}
	}
	return -1
}
