package main

import (
	"encoding/json"
	"fmt"
	"log"
	"mini-k8s/pkg/resources"
	"net/http"
	"os"
)

const apiServerAddr = "http://localhost:8080"

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: mini-kubectl get <pods|services>")
		return
	}

	command := os.Args[1]
	resourceType := os.Args[2]

	if command == "get" {
		switch resourceType {
		case "pods":
			listPods()
		case "services":
			listServices()
		default:
			fmt.Printf("Error: resource type %q not supported\n", resourceType)
		}
		// 🎯 아까 코드에 있던 listPods() 중복 호출 줄을 삭제했습니다.
	} else {
		fmt.Println("Unknown command")
	}
}

// 🎯 인자 이름을 resources에서 resList로 바꿔서 패키지 이름과 충돌을 피했습니다!
func printResources(resList []resources.Resource) {
	fmt.Println("NAME\t\tKIND\t\tDETAIL")

	for _, r := range resList {
		switch v := r.(type) {
		case resources.Pod:
			fmt.Printf("%s\t\t%s\t\t-\n", v.GetName(), v.GetKind())
		case resources.Service:
			// 🎯 이제 resources.Service 패키지 타입을 정상적으로 인식합니다.
			fmt.Printf("%s\t\t%s\t\tPort: %d\n", v.GetName(), v.GetKind(), v.Port)
		default:
			fmt.Printf("%s\t\t%s\t\tUnknown\n", r.GetName(), r.GetKind())
		}
	}
}

func listPods() {
	resp, err := http.Get(apiServerAddr + "/api/v1/pods")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer resp.Body.Close()

	var pods []resources.Pod
	json.NewDecoder(resp.Body).Decode(&pods)

	// 🎯 []Pod를 []Resource로 변환 (Go의 필수 과정)
	resList := make([]resources.Resource, len(pods))
	for i, p := range pods {
		resList[i] = p
	}
	printResources(resList)
}

func listServices() {
	resp, err := http.Get(apiServerAddr + "/api/v1/services")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer resp.Body.Close()

	var svcs []resources.Service
	json.NewDecoder(resp.Body).Decode(&svcs)

	// 🎯 []Service를 []Resource로 변환
	resList := make([]resources.Resource, len(svcs))
	for i, s := range svcs {
		resList[i] = s
	}
	printResources(resList)
}
