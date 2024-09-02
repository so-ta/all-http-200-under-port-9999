package main

import (
	"fmt"
	"net"
	"net/http"
	"sync"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK) // 200 OK
	fmt.Fprintln(w, "OK")
}

func startServer(port int, wg *sync.WaitGroup) {
	defer wg.Done()

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: http.HandlerFunc(handler),
	}

	// ポートが使用中かどうかを事前にチェック
	ln, err := net.Listen("tcp", server.Addr)
	if err != nil {
		fmt.Printf("port %d is already in use\n", port)
		return
	}
	ln.Close() // チェック後はリスナーを閉じる

	err = server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		fmt.Printf("Error starting server on port %d: %v\n", port, err)
	}
}

func main() {
	var wg sync.WaitGroup

	for port := 2; port <= 9999; port++ {
		wg.Add(1)
		go func(port int) {
			startServer(port, &wg)
		}(port)
	}

	fmt.Println("start all servers")
	wg.Wait()
}
