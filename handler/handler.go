package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"portfolio-send-message/model"
	"portfolio-send-message/service"
)

type Handler struct{
	service service.Service	
}

func NewHanlder(service service.Service) Handler{
	return Handler{service: service}
}

func writeResponse(w http.ResponseWriter, sucess bool,message string, status int){
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(model.NewOuput(sucess, message))
}

func (h *Handler)HandleAddMessage(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "application/json")
	switch r.Method{
	case http.MethodPost:
		h.handleAddMessagePost(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	} 
}

func (h *Handler)handleAddMessagePost(w http.ResponseWriter, r *http.Request){
	var messageInput model.MessageInput;

	err := json.NewDecoder(r.Body).Decode(&messageInput)
	if err != nil{
		log.Println(err)
		writeResponse(w, false,"Decode error. U sure u passed everything right?", http.StatusBadRequest)
		return
	}

	err = h.service.AddMessage(messageInput)
	if err != nil{
		log.Println(err)
		writeResponse(w, false,err.Error(), http.StatusInternalServerError)
		return
	}
	
	writeResponse(w, true, "", http.StatusCreated)
}
