// 场景描述：
// 我们可通过一下方式打开电视机：
// 1. 按下遥控器上的 ON 开关；
// 2. 按下电视机上的 ON 开关。
// 我们可以从实现 ON 命令对象并以电视机作为接收者入手。当在此命令上调用execute执行方法时，方法会调用TV.on打开电视函数。最后的工作是定义请求者：这里实际上有两个请求者：遥控器和电视机。两者都将嵌入ON命令对象。
package command

func main() {
	tv := &tv{}

	onCommand := &onCommand{
		device: tv,
	}

	offCommand := &offCommand{
		device: tv,
	}

	onButton := &button{
		command: onCommand,
	}
	onButton.press()

	offButton := &button{
		command: offCommand,
	}
	offButton.press()
}
