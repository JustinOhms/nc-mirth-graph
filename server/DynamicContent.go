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
	fmt.Fprint(w, "graphdata=")
	fmt.Fprint(w, Content)
	fmt.Fprint(w, ";")
	fmt.Printf("data request served content length %d\n", len(Content))
}

func provideUI(w http.ResponseWriter, r *http.Request, c chan int) {
	path := r.URL.Path[1:]
	fmt.Println("provide ui path:", path)
	performCopy(w, path)
	c <- 1

}

func performCopy(w http.ResponseWriter, path string) {
	fr, err := os.Open(path)
	defer fr.Close()
	fmt.Println("perfcopy ", path)
	if err != nil {
		strings.Contains(err.Error(), "no such file or directory")
		w.Header().Set("404", "Not Found")
		fmt.Fprint(w, "NOT FOUND")
		return
	}
	fmt.Println("iocopy ", path)
	b, err := io.Copy(w, fr)
	check(err)
	fmt.Println("Bytes sent:", path, b)
	fr.Close()

}

func provideUIHandler(w http.ResponseWriter, r *http.Request) {
	c := make(chan int)
	go provideUI(w, r, c)

	<-c
}

func ServeDynamicContent(c chan int, d chan string) {
	http.HandleFunc("/data.json", provideData)
	http.HandleFunc("/ui/", provideUIHandler)

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
