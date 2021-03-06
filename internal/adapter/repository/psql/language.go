package psql

import (
	"context"
	"fmt"

	"github.com/answersuck/vault/internal/domain/language"

	"github.com/answersuck/vault/pkg/logging"
	"github.com/answersuck/vault/pkg/postgres"
)

const languageTable = "language"

type languageRepo struct {
	l logging.Logger
	c *postgres.Client
}

func NewLanguageRepo(l logging.Logger, c *postgres.Client) *languageRepo {
	return &languageRepo{
		l: l,
		c: c,
	}
}

func (r *languageRepo) FindAll(ctx context.Context) ([]*language.Language, error) {
	sql := fmt.Sprintf(`
		SELECT id, name 
		FROM %s
	`, languageTable)

	r.l.Info("psql - language - FindAll: %s", sql)

	rows, err := r.c.Pool.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("psql - language - FindAll - r.c.Pool.Query: %w", err)
	}

	defer rows.Close()

	var langs []*language.Language

	for rows.Next() {
		var l language.Language

		if err = rows.Scan(&l.Id, &l.Name); err != nil {
			return nil, fmt.Errorf("psql - language - FindAll - rows.Scan: %w", err)
		}

		langs = append(langs, &l)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("psql - language - FindAll - rows.Err: %w", err)
	}

	return langs, nil
}
