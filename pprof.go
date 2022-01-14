package redis

import (
	"net"
	"net/http"
	_ "net/http/pprof"
)

var Port int

func StartPProf() {
	go func() {
		listener, err := net.Listen("tcp", "127.0.0.1:0")
		if err != nil {
			println("startPProf failed: ", err.Error())
		}
		Port = listener.Addr().(*net.TCPAddr).Port
		println("PProf using port:", listener.Addr().(*net.TCPAddr).Port)
		err = http.Serve(listener, nil)
		if err != nil {
			println("startPProf failed: ", err.Error())
		}
	}()
}
