package main

type client struct{}

func (client *client) insertSquareUsbInComputer(com computer) {
	com.insertInSquarePort()
}
