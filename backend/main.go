package main

import (
	"fmt"
	"net/http"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

func main() {
	err := http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			fmt.Println("err ws", err)
		}
		go func() {
			defer conn.Close()
			for {
				msg, op, err := wsutil.ReadClientData(conn)
				if err != nil {
					fmt.Println("read client data", err)
				}
				err = wsutil.WriteServerMessage(conn, op, msg)
				if err != nil {
					fmt.Println("write server message", err)
				}
			}
		}()
	}))
	if err != nil {
		fmt.Println("http listen and server", err)
	}

}
