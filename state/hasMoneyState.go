package main

import "fmt"

type hasMoneyState struct {
	vendingMachine *vendingMachine
}

func (i *hasMoneyState) requestItem() error {
	return fmt.Errorf("Item dispense in progress")
}

func (i *hasMoneyState) addItem(count int) error {
	return fmt.Errorf("Item dispense in progress")
}

func (i *hasMoneyState) insertMoney(money int) error {
	return fmt.Errorf("Item out of stock")
}
func (i *hasMoneyState) dispenseItem() error {
	fmt.Println("Dispensing Item")
	i.vendingMachine.itemCount = i.vendingMachine.itemCount - 1
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
	} else {
		i.vendingMachine.setState(i.vendingMachine.hasItem)
	}
	return nil
}
