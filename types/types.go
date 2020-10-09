package types

type Book struct {
	ID        string   `json:"id"`
	ISBN      string   `json:"isbn"`
	Name      string   `json:"name"`
	Publisher string   `json:"publisher"`
	Authors   []string `json:"author"`
}

type Author struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
