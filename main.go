package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

// LumpSumRequest represents the input data structure
type LumpSumRequest struct {
    Principal float64 `json:"principal"`
    Rate      float64 `json:"rate"`
    Time      float64 `json:"time"`
}

// LumpSumResponse represents the output data structure
type LumpSumResponse struct {
    Amount float64 `json:"amount"`
}

// calculateLumpSum returns the maturity amount for a lump sum investment
func calculateLumpSum(principal, rate, time float64) float64 {
    return principal * (1 + (rate / 100) * time)
}

func lumpSumHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        var req LumpSumRequest
        if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
            http.Error(w, "Invalid request payload", http.StatusBadRequest)
            return
        }
        amount := calculateLumpSum(req.Principal, req.Rate, req.Time)
        res := LumpSumResponse{Amount: amount}
        json.NewEncoder(w).Encode(res)
    } else {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
    }
}

func main() {
    fs := http.FileServer(http.Dir("."))
    http.Handle("/", fs)
    http.HandleFunc("/calculate", lumpSumHandler)
    fmt.Println("Server starting on port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
