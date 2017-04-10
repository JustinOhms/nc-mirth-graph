package server

import (
	"fmt"
	"net/http"
	"strings"
)

func provideData(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "THIS IS A TEST")
}

func ServeDynamicContent() int {
	http.HandleFunc("/data", provideData)

	var bindport = 1380
	var finalport = 0
	for finalport == 0 {
		addr := fmt.Sprintf(":%d", bindport)
		err := http.ListenAndServe(addr, nil)
		if err != nil {
			if strings.Contains(err.Error(), "bind: permission denied") {
				bindport++
			} else {
				check(err)
			}
		} else {
			finalport = bindport
		}

	}
	return finalport
}
