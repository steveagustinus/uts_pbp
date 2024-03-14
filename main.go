package main

import (
	"fmt"
	"log"
	"net/http"
	"uts/controllers"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	os.setEnv("token", "GLwIjWz4uGH8Bnc4YpDNl8JUKSna9hQpbmm38mDpY9O")
	router := mux.NewRouter()
	router.HandleFunc("/room/get", controllers.GetAllRooms).Methods("GET")
	router.HandleFunc("/room/getdetail", controllers.GetRoomDetail).Methods("GET")
	router.HandleFunc("/room/enter", controllers.EnterRoom).Methods("POST")
	router.HandleFunc("/room/leave", controllers.LeaveRoom).Methods("DELETE")

	http.Handle("/", router)
	fmt.Println("Connected to port 22345")
	log.Println("Connected to port 22345")
	log.Fatal(http.ListenAndServe(":22345", router))
}
