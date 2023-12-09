package main

import (
	"github.com/gorilla/websocket"
	"hotel_train_antonyGG/types"
	"log"
	"math"
	"math/rand"
	"time"
)

const (
	sendInterval = time.Second
	wsEndpoint   = "ws://127.0.0.1:30000/ws"
)

func genLocation() (float64, float64) {
	return genCoard(), genCoard()
}

func genCoard() float64 {
	n := float64(rand.Intn(100) + 1)
	f := rand.Float64()
	return n + f
}

func main() {
	obuIds := generateOBUIDS(20)

	conn, _, err := websocket.DefaultDialer.Dial(wsEndpoint, nil)

	if err != nil {
		log.Fatal(err)
	}

	for {
		for i := 0; i < len(obuIds); i++ {

			lat, long := genLocation()

			data := types.OBUData{
				OBUID: obuIds[i],
				Lat:   lat,
				Long:  long,
			}

			if err = conn.WriteJSON(data); err != nil {
				log.Fatal(err)
			}

		}

		time.Sleep(sendInterval)
	}

}

func generateOBUIDS(n int) []int {
	ids := make([]int, n)
	for i := 0; i < n; i++ {
		ids[i] = rand.Intn(math.MaxInt)
	}
	return ids
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
