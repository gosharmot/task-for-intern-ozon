package store

import (
	"task-for-intern/internal/shortener"
)

type URLRepository struct {
	store *Store
}

func (u *URLRepository) GetShortUrl(longUrl string) (string, error) {
	var shortUrl string
	if err := u.store.db.QueryRow(
		"SELECT short_url FROM urls WHERE long_url = $1",
		longUrl,
	).Scan(
		&shortUrl,
	); err != nil {
		if err.Error() == "sql: no rows in result set" {
			var id int

			if err := u.store.db.QueryRow(
				"insert into urls (long_url) values ($1) RETURNING id",
				longUrl,
			).Scan(
				&id,
			); err != nil {
				return "", err
			}

			shortUrl = shortener.GetShortUrl(longUrl, 9, id)

			if err := u.store.db.QueryRow(
				"update urls SET short_url = $1 where id = $2;",
				shortUrl,
				id,
			).Err(); err != nil {
				return "", err
			}

			return shortUrl, nil
		}
		return "", err
	}

	return shortUrl, nil
}

func (u *URLRepository) GetLongUrl(shortUrl string) (string, error) {
	var longUrl string
	if err := u.store.db.QueryRow(
		"SELECT long_url FROM urls WHERE short_url = $1",
		shortUrl,
	).Scan(
		&longUrl,
	); err != nil {
		return "", err
	}

	return longUrl, nil
}
