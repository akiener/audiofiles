package database

import (
	"fmt"
	"net/url"
	"os"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jackc/pgx/v4"
	"github.com/jmoiron/sqlx"
)

type Database struct {
	conn *sqlx.DB
}

type Song struct {
	Id          int    `db:"id"`
	Name        string `db:"name"`
	Artist      string `db:"artist_name"`
	Album       string `db:"album_name"`
	Genre       string `db:"genre"`
	Description string `db:"description"`
	DcUser      string `db:"dc_user_name"`
}

func New() (*Database, error) {
	connConfig, err := pgx.ParseConfig("")
	if err != nil {
		return nil, fmt.Errorf("failed to create db connection config: %w", err)
	}
	connConfig.Host = os.Getenv("AUDIOFILES_DB_HOST")
	connConfig.Port = 5432
	connConfig.Database = "audiofiles"
	connConfig.User = "audiofiles"
	connConfig.Password = os.Getenv("AUDIOFILES_DB_PASSWORD")

	println(connConfig.ConnString())
	conn, err := sqlx.Connect("pgx",
		"postgresql://audiofiles:"+
			url.QueryEscape(os.Getenv("AUDIOFILES_DB_PASSWORD"))+
			"@"+os.Getenv("AUDIOFILES_DB_HOST")+
			":5432/audiofiles")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db: %w", err)
	}

	return &Database{conn: conn}, nil
}

func (d *Database) Test() (*[]Song, error) {
	var songs []Song
	err := d.conn.Select(&songs, `
SELECT s.id AS id, s.name AS name, a.name AS artist_name, bum.name AS album_name, genre, description, u.name AS dc_user_name
FROM song s
         JOIN dc_user u ON s.dc_user_id = u.id
         JOIN album bum ON s.album_id = bum.id
         JOIN artist a ON a.id = s.artist_id;`,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to select songs: %w", err)
	}

	return &songs, nil
}

func (d *Database) Close() error {
	return d.conn.Close()
}
