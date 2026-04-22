package api

import (
	"encoding/json"
	"fmt"
	"io"
	"mini-k8s/pkg/registry"
	"mini-k8s/pkg/resources"
	"net/http"
)

type Server struct {
	reg *registry.Registry
}

// NewServer APIServer 생성자
func NewServer(reg *registry.Registry) *Server {
	return &Server{reg: reg}
}

// Error 응답
func sendJSONError(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"message": message})
}

// ListPods [GET] 목록 조회
func (s *Server) ListPods(w http.ResponseWriter, r *http.Request) {
	list := s.reg.List()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}

// GetPod [GET] 개별 조회
func (s *Server) GetPod(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	pod, err := s.reg.Get(name)
	if err != nil {
		sendJSONError(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pod)
}

// CreatePod [POST] 생성
func (s *Server) CreatePod(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		sendJSONError(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var p resources.Pod
	if err := json.Unmarshal(body, &p); err != nil {
		sendJSONError(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// TODO 임시, 저장 전 Kind를 명시적으로 저장
	p.Kind = "Pod"

	if err := s.reg.Register(p); err != nil {
		sendJSONError(w, err.Error(), http.StatusConflict)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Pod %s registered successfully!", p.GetName())
}
