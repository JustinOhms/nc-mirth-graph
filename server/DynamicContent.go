package server

import (
	"fmt"
	"io"

	"net"
	"net/http"
	"os"
	"strings"
)

var Content string

func provideData(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("data request in: %s \n", string(r.RequestURI))
	fmt.Fprint(w, Content)
	fmt.Printf("data request served content length %d\n", len(Content))
}

func provideUI(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	fmt.Printf(path)

	fr, err := os.Open(path)
	defer fr.Close()

	if err != nil {
		strings.Contains(err.Error(), "no such file or directory")
		w.Header().Set("404", "Not Found")
		fmt.Fprint(w, "NOT FOUND")
		return
	}

	io.Copy(w, fr)
	fr.Close()
}

func ServeDynamicContent(c chan int, d chan string) {
	http.HandleFunc("/data", provideData)
	http.HandleFunc("/ui/", provideUI)

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
	check(server.Serve(tcpKeepAliveListener{listener.(*net.TCPListener)}))
}

func updateDynamicContent(d chan string) {
	Content = <-d
}
