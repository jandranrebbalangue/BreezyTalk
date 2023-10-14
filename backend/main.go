package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

func main() {
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
	log.Fatal(http.ListenAndServe(":8081", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			log.Println("err ws", err)
		}
		go func() {
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
				timeString := updatedResponse.Time.Format("2006/01/02 15:04:05")
				msg, op, err := wsutil.ReadClientData(conn)
				msgTxt := string(msg)
				errs := json.Unmarshal([]byte(msgTxt), &response)
				if errs != nil {
					log.Print("Error json:", errs)
					return
				}
				message := `<div id="idMessage" hx-swap-oob="true"> ` + timeString + " " + response.Message + `</div>`
				log.Print("Received Message:", message)
				if err != nil {
					log.Fatal("read client data", err)
					return
				}
				err = wsutil.WriteServerMessage(conn, op, []byte(message))
				if err != nil {
					log.Fatal("write server message", err)
					return
				}
			}
		}()
	})))

}
