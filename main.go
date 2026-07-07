package main

import (
 "fmt"
 "log"
 "net/http"
 "os"
 "time"
)

func main() {
 // Get pod name from environment
 hostname, _ := os.Hostname()
 http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello from pod: %s\n", hostname)
  fmt.Fprintf(w, "Path: %s\n", r.URL.Path)
  fmt.Fprintf(w, "Time: %s\n", time.Now().Format("2006-01-02 15:04:05"))
 })
 http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusOK)
  fmt.Fprintf(w, "Healthy")
 })
 http.HandleFunc("/ready", func(w http.ResponseWriter, r *http.Request) {
  // Simulate readiness check (in real app, check dependencies)
  w.WriteHeader(http.StatusOK)
  fmt.Fprintf(w, "Ready")
 })
 log.Println("Server starting on :8080")
 log.Fatal(http.ListenAndServe(":8080", nil))
}
