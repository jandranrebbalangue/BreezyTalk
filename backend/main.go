package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

type Response struct {
	Message string `json:"message"`
	Headers struct {
		HXRequest     string `json:"HX-Request"`
		HXTrigger     string `json:"HX-Trigger"`
		HXTriggerName string `json:"HX-Trigger-Name"`
		HXTarget      string `json:"HX-Target"`
		HXCurrentUrl  string `json:"HX-Current-URL"`
	} `json:"HEADERS"`
	Time time.Time `json:"time"`
}

var clients = make(map[net.Conn]bool)

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, _, _, err := ws.UpgradeHTTP(r, w)
	if err != nil {
		log.Println("err ws", err)
	}
	go func() {
		clients[conn] = true
		defer conn.Close()
		for {
			response := Response{}
			response.Time = time.Now()
			updatedJsonData, _ := json.Marshal(response)
			var updatedResponse Response
			err = json.Unmarshal(updatedJsonData, &updatedResponse)
			if err != nil {
				log.Println("Error:", err)
				return
			}
			timeString := updatedResponse.Time.Format("2006/01/02 3:04PM")
			msg, op, err := wsutil.ReadClientData(conn)
			if err != nil {
				log.Fatal("read client data", err)
				delete(clients, conn)
				return
			}
			msgTxt := string(msg)
			errs := json.Unmarshal([]byte(msgTxt), &response)
			if errs != nil {
				log.Print("Error json:", errs)
				return
			}
			for client := range clients {
				message := `<div id="idMessage" hx-swap-oob="beforeend:#idMessage"> <p>` + timeString + " " + response.Message + `</p></div>`
				err = wsutil.WriteServerMessage(client, op, []byte(message))
			}

			if err != nil {
				log.Fatal("write server message", err)
				return
			}
		}
	}()

}

func main() {
	http.HandleFunc("/chat", handleWebSocket)
	log.Fatal(http.ListenAndServe(":8081", nil))

}
