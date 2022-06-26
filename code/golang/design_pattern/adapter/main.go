package adapter

// 客户端可以直连mac的Lightning 接口
// windows没有提供Lightning 接口，需要适配器将它的接口转成Lightning 接口。
func main() {
	client := &client{}
	mac := &mac{}

	client.insertLightningConnectorIntoComputer(mac)

	windowsMachine := &windows{}
	windowsMachineAdapter := &windowsAdapter{
		windowMachine: windowsMachine,
	}

	client.insertLightningConnectorIntoComputer(windowsMachineAdapter)
}
