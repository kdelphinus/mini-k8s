package registry

import (
	"encoding/json"
	"fmt"
	"mini-k8s/pkg/resources"
	"os"
	"sync"
)

type Registry struct {
	mu        sync.RWMutex
	resources map[string]resources.Resource
	filePath  string // 저장할 파일 경로
}

// NewRegistry Registry의 생성자
func NewRegistry(path string) *Registry {
	reg := &Registry{
		resources: make(map[string]resources.Resource),
		filePath:  path,
	}
	reg.load() // 기존 데이터 불러오기
	return reg
}

// save 파일에 저장하는 내부 메서드
func (r *Registry) save() error {
	data, err := json.MarshalIndent(r.resources, "", "	")
	if err != nil {
		return err
	}
	return os.WriteFile(r.filePath, data, 0644)
}

// load 파일에서 불어오는 내부 메서드
func (r *Registry) load() {
	data, err := os.ReadFile(r.filePath)
	if err != nil {
		return // 파일 없으면 빈 상태로 시작
	}

	var tempMap map[string]json.RawMessage
	if err := json.Unmarshal(data, &tempMap); err != nil {
		return
	}

	for name, raw := range tempMap {
		var typeChecker struct {
			Kind string `json:"kind"`
		}
		json.Unmarshal(raw, &typeChecker)

		switch typeChecker.Kind {
		case "Pod":
			var p resources.Pod
			json.Unmarshal(raw, &p)
			r.resources[name] = p
		case "Service":
			var s resources.Service
			json.Unmarshal(raw, &s)
			r.resources[name] = s
		}
	}
}

// Register 등록(Create/Update)
func (r *Registry) Register(res resources.Resource) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.resources[res.GetName()]; exists {
		return fmt.Errorf("resource %s already exists", res.GetName())
	}

	r.resources[res.GetName()] = res
	return r.save() // 등록할 때마다 파일에 저장
}

// Get 단일 조회(Read)
func (r *Registry) Get(name string) (resources.Resource, error) {
	res, ok := r.resources[name] // 'comma ok' 패턴
	if !ok {
		return nil, fmt.Errorf("resource %s not found", name)
	}
	return res, nil
}

// List 전체 목록 조회(List)
func (r *Registry) List() []resources.Resource {
	list := make([]resources.Resource, 0, len(r.resources))
	for _, v := range r.resources {
		list = append(list, v)
	}
	return list
}

// Delete 삭제(Delete)
func (r *Registry) Delete(name string) {
	delete(r.resources, name)
}
