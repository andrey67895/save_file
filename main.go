package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Данные будут храниться в этой переменной
var data []string

func saveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Метод %s не разрешен\n", r.Method)
		return
	}

	// Чтение тела запроса
	buf := make([]byte, 1024)
	n, err := r.Body.Read(buf)
	if err != nil {
		log.Printf("Ошибка чтения: %v", err)
		http.Error(w, "Ошибка чтения данных", http.StatusInternalServerError)
		return
	}

	// Сохраняем полученные данные
	data = append(data, string(buf[:n]))
	fmt.Fprintln(w, "Данные сохранены")
}

func getDataHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Метод %s не разрешен\n", r.Method)
		return
	}

	marshal, err := json.Marshal(data)
	if err != nil {
		return
	}
	w.Write(marshal)
}

func getPing(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Метод %s не разрешен\n", r.Method)
		return
	}
	fmt.Println("1")

	w.Write([]byte("OK"))
}

func main() {
	http.HandleFunc("/save", saveHandler)
	http.HandleFunc("/get", getDataHandler)
	http.HandleFunc("/ping", getPing)

	log.Println("Запуск сервера...")
	err := http.ListenAndServe(":8383", nil)
	if err != nil {
		log.Fatal(err)
	}
}
