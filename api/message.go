package handler

import (
    "net/http"
    "os"
    "portfolio-send-message/handler"
    "portfolio-send-message/service"
    "portfolio-send-message/storage"
    "github.com/rs/cors"
)

func Handler(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    botToken := os.Getenv("BOT_TOKEN")
    channelID := os.Getenv("CHANNEL_ID")

    storage, err := storage.NewDiscordStorage(botToken, channelID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    svc := service.NewService(storage)
    h := handler.NewHanlder(svc)

    mux := http.NewServeMux()
    mux.HandleFunc("/", h.HandleAddMessage)

    corsHandler := cors.Default().Handler(mux)
    corsHandler.ServeHTTP(w, r)
}
