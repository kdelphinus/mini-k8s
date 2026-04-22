package main

import (
	"log"
	"mini-k8s/pkg/api"
	"mini-k8s/pkg/registry"
	"net/http"
)

func main() {
	reg := registry.NewRegistry("./resources.json")
	server := api.NewServer(reg)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/v1/pods", server.ListPods)
	mux.HandleFunc("POST /api/v1/pods", server.CreatePod)
	mux.HandleFunc("GET /api/v1/pods/{name}", server.GetPod)

	log.Println("🚀 mini-k8s API Server starting on :8080...")
	// 에러 처리를 위해 log.Fatal로 감싸줍니다.
	log.Fatal(http.ListenAndServe(":8080", mux))
}
