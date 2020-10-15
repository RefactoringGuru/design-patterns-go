package main

import "fmt"

type hasItemState struct {
	vendingMachine *vendingMachine
}

func (i *hasItemState) requestItem() error {
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
		return fmt.Errorf("No item present")
	}
	fmt.Printf("Item requestd\n")
	i.vendingMachine.setState(i.vendingMachine.itemRequested)
	return nil
}

func (i *hasItemState) addItem(count int) error {
	fmt.Printf("%d items added\n", count)
	i.vendingMachine.incrementItemCount(count)
	return nil
}

func (i *hasItemState) insertMoney(money int) error {
	return fmt.Errorf("Please select item first")
}
func (i *hasItemState) dispenseItem() error {
	return fmt.Errorf("Please select item first")
}
