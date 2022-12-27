package dictionary

import (
	"fmt"
	"github.com/dgraph-io/badger/v3"
	"time"
)

type Dictionary struct {
	db *badger.DB // Point d'entrée de la bd pour récup un dictionnaire
}

type Entry struct {
	Word       string
	Definition string
	CreatedAt  time.Time
}

// Fonction helper pour l'affichage
func (e Entry) String() string {
	created := e.CreatedAt.Format(time.Stamp)
	return fmt.Sprintf("%-10v\t%-50v%-6v", e.Word, e.Definition, created)
}

func New(dir string) (*Dictionary, error) {
	opts := badger.DefaultOptions(dir)
	opts.Dir = dir
	opts.ValueDir = dir

	db, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}

	dict := &Dictionary{
		db: db,
	}
	return dict, nil
}

func (d *Dictionary) Close() {
	d.db.Close()
}
