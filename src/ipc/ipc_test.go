package ipc

import "testing"

type EchoServer struct{}

func (server *EchoServer) Handle(method, request string) *Response {
	return &Response{"返回码", request}
	//return "ECHO:" + request
}
func (server *EchoServer) Name() string {
	return "EchoServer"
}
func TestIpc(t *testing.T) {
	server := NewIpcServer(&EchoServer{})
	client1 := NewIpcClient(server)
	client2 := NewIpcClient(server)
	resp1, _ := client1.Call("method", "From Client1")
	resp2, _ := client2.Call("method", "From Client2")

	if resp1.Body != "From Client1" || resp2.Body != "From Client2" {
		t.Error("IpcClient.Call failed. resp1:", resp1, "resp2:", resp2)
	}

	//	resp3, _ := client1.Call("method", "From Client1 again")
	//	fmt.Println(resp3)
	client1.Close()
	client2.Close()
}
