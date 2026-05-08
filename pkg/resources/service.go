package resources

type Service struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
	Port int    `json:"port"`
}

func (s Service) GetName() string { return s.Name }
func (s Service) GetKind() string { return s.Kind }
