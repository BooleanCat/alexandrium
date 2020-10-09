package memory

import (
	"github.com/BooleanCat/alexandrium/store"
	"github.com/BooleanCat/alexandrium/types"
)

type AuthorStore struct{}

func (_ *AuthorStore) ByID(id string) (types.Author, error) {
	for _, author := range authorsData {
		if author.ID == id {
			return author, nil
		}
	}

	return types.Author{}, store.NotFoundError{}
}

var _ store.Authors = new(AuthorStore)

var authorsData = []types.Author{
	{
		ID:   "ea1ff7d7-67cd-477c-8cb7-8756619e275d",
		Name: "Adrian Tchaikovsky",
	},
}
