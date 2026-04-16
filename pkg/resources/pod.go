package resources

type Pod struct {
	Name string
}

func (p Pod) GetName() string { return p.Name }
func (p Pod) GetKind() string { return "Pod" }
