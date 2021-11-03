package main

func main() {
	client := &client{}
	macMachine := &mac{}
	client.insertSquareUsbInComputer(macMachine)

	windowsMachine := &windows{}
	windowsAdapter := &windowsAdapter{
		windowsMachine: windowsMachine,
	}

	client.insertSquareUsbInComputer(windowsAdapter)

}
