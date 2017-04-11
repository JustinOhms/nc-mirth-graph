package server

import (
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"
)

func provideData(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "THIS IS A TEST")
}

func ServeDynamicContent(c chan int) error {
	http.HandleFunc("/data", provideData)

	var server *http.Server
	var listener net.Listener
	var bindport = 1380
	var finalport = 0
	for finalport == 0 {
		addr := fmt.Sprintf(":%d", bindport)
		//err := http.ListenAndServe(addr, nil)
		server = &http.Server{Addr: addr, Handler: nil}
		//err := server.ListenAndServe()

		ln, err := net.Listen("tcp", addr)
		if err != nil {
			if strings.Contains(err.Error(), "bind: permission denied") {
				bindport++
			} else {
				check(err)
			}
		} else {
			finalport = bindport
		}
		listener = ln
	}
	fmt.Println("p", finalport)
	c <- finalport
	return server.Serve(tcpKeepAliveListener{listener.(*net.TCPListener)})
}

// Unfortunately because tcpKeepAliveListener in net/html is private I have to reimplement it
// tcpKeepAliveListener sets TCP keep-alive timeouts on accepted
// connections. It's used by ListenAndServe and ListenAndServeTLS so
// dead TCP connections (e.g. closing laptop mid-download) eventually
// go away.
type tcpKeepAliveListener struct {
	*net.TCPListener
}

func (ln tcpKeepAliveListener) Accept() (c net.Conn, err error) {
	tc, err := ln.AcceptTCP()
	if err != nil {
		return
	}
	tc.SetKeepAlive(true)
	tc.SetKeepAlivePeriod(3 * time.Minute)
	return tc, nil
}
