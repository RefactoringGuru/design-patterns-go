package main

import "fmt"

type Client struct {
}

func (c *Client) insertLightningConnectorIntoComputer(com Computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.InsertIntoLightningPort()
}
