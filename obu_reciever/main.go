package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"hotel_train_antonyGG/types"
	"log"
	"net/http"
)

func main() {

	recv, err := NewDataReceiver()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/ws", recv.wsHandle)
	http.ListenAndServe(":30000", nil)
}

type DataReceiver struct {
	msg  chan types.OBUData
	conn *websocket.Conn
	prod DataProducer
}

func (dr *DataReceiver) wsHandle(w http.ResponseWriter, r *http.Request) {

	u := websocket.Upgrader{
		ReadBufferSize:  1028,
		WriteBufferSize: 1028,
	}

	conn, err := u.Upgrade(w, r, nil)

	if err != nil {
		log.Fatal(err)
	}

	dr.conn = conn

	go dr.wsLoop()

}

func NewDataReceiver() (*DataReceiver, error) {

	var (
		p   DataProducer
		err error
	)

	p, err = NewKafkaProducer()

	if err != nil {
		return nil, err
	}

	p = NewLogMiddleware(p)

	return &DataReceiver{
		msg:  make(chan types.OBUData, 128),
		prod: p,
	}, nil
}

func (dr *DataReceiver) produceData(data types.OBUData) error {
	return dr.prod.ProduceData(data)
}

func (dr *DataReceiver) wsLoop() {

	fmt.Println("client connected")

	for {
		var data types.OBUData
		if err := dr.conn.ReadJSON(&data); err != nil {
			log.Println(err)
			continue
		}

		if err := dr.produceData(data); err != nil {
			fmt.Println("kafka producer error", err.Error())
		}

	}

}
