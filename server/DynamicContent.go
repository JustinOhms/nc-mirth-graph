package server

import (
	"fmt"

	"github.com/caimeo/console"

	"net"
	"net/http"
	"strings"
)

var Content string
var UseLocal bool = false

var finishedchannel chan bool

func provideData(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "graphdata=")
	fmt.Fprint(w, Content)
	fmt.Fprint(w, ";")
	console.Verbose(fmt.Sprintf("serving data: %d", len(Content)))
	finishedchannel <- true
}

func provideUI(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path //r.URL.Path[1:]
	console.Verbose("serving interface: ", path)
	FSIoCopy(UseLocal, path, w)
}

func ServeDynamicContent(p chan int, d chan string, f chan bool) {
	http.HandleFunc("/data.json", provideData)
	http.HandleFunc("/ui/", provideUI)

	finishedchannel = f

	go updateDynamicContent(d)

	var server *http.Server
	var listener net.Listener
	var bindport = 1380
	var finalport = 0
	for finalport == 0 {
		addr := fmt.Sprintf(":%d", bindport)
		server = &http.Server{Addr: addr, Handler: nil}
		ln, err := net.Listen("tcp", addr)
		if err != nil {
			if strings.Contains(err.Error(), "bind: permission denied") || strings.Contains(err.Error(), "bind: address already in use") {
				bindport++
			} else {
				check(err)
			}
		} else {
			finalport = bindport
		}
		listener = ln
	}
	//fmt.Println("p", finalport)
	p <- finalport
	check(server.Serve(tcpKeepAliveListener{listener.(*net.TCPListener)}))
	for {

	}
}

func updateDynamicContent(d chan string) {
	//just wait for content to be updated
	for {
		Content = <-d
	}
}
