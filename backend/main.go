package main

import (
	"log"
	"net/http"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

func main() {
	// handler := http.HandlerFunc(PlayerServer)
	// log.Fatal(http.ListenAndServe(":5000", handler))
	log.Fatal(http.ListenAndServe(":8081", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, _, _, err := ws.UpgradeHTTP(r, w)
		log.Print("Client connection network ", conn.LocalAddr().Network())
		if err != nil {
			log.Println("err ws", err)
		}
		go func() {
			defer conn.Close()

			for {
				msg, op, err := wsutil.ReadClientData(conn)
				msgTxt := string(msg)
				log.Print("Received Message:", msgTxt)
				if err != nil {
					log.Fatal("read client data", err)
				}
				err = wsutil.WriteServerMessage(conn, op, msg)
				if err != nil {
					log.Fatal("write server message", err)
				}
			}
		}()
	})))

}
