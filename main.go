package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Request struct {
	Websites []string `json:"websites"`
}

type Result struct {
	Website string        `json:"website"`
	Status  string        `json:"status"`
	Latency time.Duration `json:"latency"`
	Error   string        `json:"error,omitempty"`
}

func CheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	resultsChan := make(chan Result, len(req.Websites))

	for _, site := range req.Websites {
		go checkWebsite(site, resultsChan)
	}

	var results []Result
	for range req.Websites {
		res := <-resultsChan
		results = append(results, res)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func checkWebsite(url string, ch chan<- Result) {
	start := time.Now()

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(url)
	latency := time.Since(start)

	if err != nil {
		ch <- Result{Website: url, Status: "DOWN", Latency: latency, Error: err.Error()}
		return
	}
	defer resp.Body.Close()

	status := "DOWN"
	if resp.StatusCode == 200 {
		status = "UP"
	}

	ch <- Result{Website: url, Status: status, Latency: latency}
}

func main() {
	http.HandleFunc("/check", CheckHandler)
	
	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server failed:", err)
	}
}