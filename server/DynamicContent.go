package server

import (
	"fmt"

	"net"
	"net/http"
	"strings"
)

var Content string

var finishedchannel chan bool

func provideData(w http.ResponseWriter, r *http.Request) {
	//fmt.Printf("data request in: %s \n", string(r.RequestURI))
	fmt.Fprint(w, "graphdata=")
	fmt.Fprint(w, Content)
	fmt.Fprint(w, ";")
	fmt.Printf("serving data: %d\n", len(Content))
	finishedchannel <- true
}

func provideUI(w http.ResponseWriter, r *http.Request, c chan int) {
	path := r.URL.Path //r.URL.Path[1:]
	fmt.Println("serving interface:", path)

	uselocal := true
	FSIoCopy(uselocal, path, w)

	c <- 1
}

func provideUIHandler(w http.ResponseWriter, r *http.Request) {
	c := make(chan int, 10)
	go provideUI(w, r, c)

	for {
		<-c
	}
}

func ServeDynamicContent(c chan int, d chan string, f chan bool) {
	http.HandleFunc("/data.json", provideData)
	http.HandleFunc("/ui/", provideUIHandler)

	finishedchannel = f

	go updateDynamicContent(d)

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
	c <- finalport
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
