package proxy

// 服务端接口
type server interface {
	handleRequest(string, string) (int, string)
}
