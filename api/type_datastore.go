package api

type Datastore struct {
	Type    string   `json:"type,omitempty"`
	Servers []string `json:"servers,omitempty" yaml:"servers,omitempty"`
}
