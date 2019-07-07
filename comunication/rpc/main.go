package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	fmt.Println(args)
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

type TestMe struct {
	Count int
}

func (t *TestMe) GetCount(args *Args, reply *int) error {
	t.Count++
	*reply = t.Count
	return nil
}

type Client struct {
	proto       string
	addr        string
	handler     func(c *net.TCPConn) error
	concurrency int
	size        int
	nflight     int
	reqres      bool
	saddr       string
}

func (c *Client) NewConnection() (*net.TCPConn, error) {
	srcTcpAddr, _ := net.ResolveTCPAddr(c.proto, c.saddr)
	dstTcpAddr, _ := net.ResolveTCPAddr(c.proto, c.addr)
	return net.DialTCP(c.proto, srcTcpAddr, dstTcpAddr)
}

func main() {
	arith := new(Arith)
	rpc.Register(arith)

	testme := new(TestMe)
	rpc.RegisterName("testing", testme)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go jsonrpc.ServeConn(conn)
		// fmt.Println("new connection" + conn.RemoteAddr().String())
		// srv := jsonrpc.NewServerCodec(conn)
		// srv.rpc
		// fmt.Println(srv)
		// err := srv.ReadRequestHeader(srv)
		// checkError(err)
		// srv.ReadRequestBody()
		// srv.Close()
		// srv.ReadRequestHeader()
		// resp := rpc.Response()
		// jsonrpc.Res
		// srv.WriteResponse()
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
