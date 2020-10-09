package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"

	"github.com/BooleanCat/alexandrium/store"
	"github.com/BooleanCat/alexandrium/types"
)

type AuthorStore struct {
	Connection *pgx.Conn
}

func (authors AuthorStore) ByID(id string) (types.Author, error) {
	var author types.Author

	if err := authors.Connection.QueryRow(context.Background(), `SELECT uuid, name FROM author WHERE uuid = $1`, id).Scan(&author.ID, &author.Name); err != nil {
		if err == pgx.ErrNoRows {
			return types.Author{}, store.NotFoundError{}
		}

		return types.Author{}, fmt.Errorf(`find author "%s": %w`, id, err)
	}

	return author, nil
}

var _ store.Authors = new(AuthorStore)
