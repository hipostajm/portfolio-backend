package main

import (
	"log"
	"net/http"
	"os"
	"portfolio-send-message/handler"
	"portfolio-send-message/service"
	"portfolio-send-message/storage"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)


var(
	BOT_TOKEN string	
	CHANNEL_ID string
)

func loadEnvs() error{
	err := godotenv.Load()
	if err != nil{
		return err
	}
	
	BOT_TOKEN = os.Getenv("BOT_TOKEN")
	CHANNEL_ID = os.Getenv("CHANNEL_ID")
	
	return nil
}

func ServeMessage(){
	err := loadEnvs()
	if err != nil{
		log.Println(err)
		return
	}
	storage, err := storage.NewDiscordStorage(BOT_TOKEN, CHANNEL_ID);
	
	if err != nil {
		log.Println(err)
		return
	}

	service := service.NewService(storage)
	handler := handler.NewHanlder(service)

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.HandleAddMessage)
	corsHandler := cors.Default().Handler(mux)
	http.ListenAndServe(":8080", corsHandler)
}
