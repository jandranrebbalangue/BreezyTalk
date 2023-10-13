package main

import (
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(PlayerServer)
	log.Fatal(http.ListenAndServe(":5000", handler))
	// err := http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// u := ws.Upgrader{
	// 	OnHeader: func(key, value []byte) (err error) {
	// 		log.Printf("non-websocket header: %q=%q", key, value)
	// 		return
	// 	},
	// }
	// conn, _, _, err := ws.UpgradeHTTP(r, w)
	// log.Print("Client connection network ", conn.LocalAddr().Network())
	// n, errs := conn.Write([]byte("Test"))
	// log.Println("conn write errs", errs)
	// log.Println("conn write ", n)
	// if err != nil {
	// 	log.Println("err ws", err)
	// }
	// go func() {
	// 	defer conn.Close()
	// 	for {
	// 		msg, op, err := wsutil.ReadClientData(conn)
	// 		if err != nil {
	// 			log.Println("read client data", err)
	// 		}
	// 		err = wsutil.WriteServerMessage(conn, op, msg)
	// 		if err != nil {
	// 			log.Println("write server message", err)
	// 		}
	// 	}
	// }()
	// }))
	// if err != nil {
	// 	log.Println("http listen and server", err)
	// }

}
