package main

import (
	"encoding/json"
	"fmt"
	"io"
	"mini-k8s/pkg/registry"
	"mini-k8s/pkg/resources"
	"net/http"
)

func main() {
	// 저장소 생성
	reg := registry.NewRegistry()

	http.HandleFunc("/api/v1/pods", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			// 저장소에서 전체 목록 조회
			list := reg.List()

			// 목록을 JSON으로 변환
			data, err := json.Marshal(list)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// 응답 헤더를 설정하고 데이터 전송
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(data)

		case http.MethodPost:
			body, _ := io.ReadAll(r.Body)
			defer r.Body.Close() // 함수가 끝날 때 자동으로 실행

			var p resources.Pod
			// if문 안에서만 err 변수 유효
			if err := json.Unmarshal(body, &p); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			reg.Register(p)

			w.WriteHeader(http.StatusCreated)
			fmt.Fprintf(w, "Pod %s registered successfully!", p.GetName())
		}
	})

	fmt.Println("서버가 8080 포트에서 시작됩니다.")
	http.ListenAndServe(":8080", nil)
}
