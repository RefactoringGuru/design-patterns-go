package main

type button struct {
	command command
}

func (b *button) press() {
	b.command.execute()
}
