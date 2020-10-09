package types

type Book struct {
	ID        string `json:"id"`
	ISBN      string `json:"isbn"`
	Name      string `json:"name"`
	Publisher string `json:"publisher"`
	Author    string `json:"author"`
}
