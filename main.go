package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ScriptBlock struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	Type    string `json:"type"`
}

func getScriptBlock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	scriptBlocks := []ScriptBlock{
		{
			ID:      "script1",
			Content: "console.log('Dynamic script loaded!'); alert('Hello from server-loaded script!');",
			Type:    "javascript",
		},
		{
			ID:      "script2",
			Content: "document.body.style.backgroundColor = document.body.style.backgroundColor === 'lightblue' ? '' : 'lightblue'; console.log('Background color toggled!');",
			Type:    "javascript",
		},
	}

	vars := mux.Vars(r)
	scriptID := vars["id"]

	if scriptID == "" {
		json.NewEncoder(w).Encode(scriptBlocks)
		return
	}

	for _, script := range scriptBlocks {
		if script.ID == scriptID {
			json.NewEncoder(w).Encode(script)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Script not found"})
}

func main() {
	r := mux.NewRouter()

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")
	})

	r.HandleFunc("/api/scripts", getScriptBlock).Methods("GET")
	r.HandleFunc("/api/scripts/{id}", getScriptBlock).Methods("GET")

	fmt.Println("Server starting on :9090")
	log.Fatal(http.ListenAndServe(":9090", r))
}
