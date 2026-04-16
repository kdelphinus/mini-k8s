package registry

import (
	"fmt"
	"mini-k8s/pkg/resources"
)

type Registry struct {
	// Resources 인터페이스를 지킨 모든 것을 담음
	resources map[string]resources.Resource
}

// Registry의 생성자
func NewRegistry() *Registry {
	return &Registry{
		resources: make(map[string]resources.Resource),
	}
}

// 등록(Create/Update)
func (r *Registry) Register(res resources.Resource) {
	r.resources[res.GetName()] = res
}

// 단일 조회(Read)
func (r *Registry) Get(name string) (resources.Resource, error) {
	res, ok := r.resources[name] // 'comma ok' 패턴
	if !ok {
		return nil, fmt.Errorf("resource %s not found", name)
	}
	return res, nil
}

// 전체 목록 조회(List)
func (r *Registry) List() []resources.Resource {
	list := make([]resources.Resource, 0, len(r.resources))
	for _, v := range r.resources {
		list = append(list, v)
	}
	return list
}

// 삭제(Delete)
func (r *Registry) Delete(name string) {
	delete(r.resources, name)
}
