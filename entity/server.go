package entity

type Server struct {
	Name string `json:"name"`
	Zone string `json:"zone"`
}

type Servers struct {
	Servers []Server `json:"servers"`
}
