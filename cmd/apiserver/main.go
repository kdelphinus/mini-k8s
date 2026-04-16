package main

import (
	"fmt"
	"mini-k8s/pkg/resources"
)

func main() {
	p := resources.Pod{Name: "my-pod"}

	fmt.Println(p.GetName())
}
