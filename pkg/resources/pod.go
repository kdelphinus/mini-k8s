package resources

type Pod struct {
	Name string `json:"name"`
	Kind string `json:"kind"`
}

func (p Pod) GetName() string { return p.Name }
func (p Pod) GetKind() string { return "Pod" }
