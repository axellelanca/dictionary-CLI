package dictionary

import (
	"bytes"
	"encoding/gob"
	"github.com/dgraph-io/badger/v3"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"sort"
	"time"
)

func (d *Dictionary) Add(word string, definition string) error {
	caser := cases.Title(language.French)
	entry := Entry{
		Word:       caser.String(word),
		Definition: definition,
		CreatedAt:  time.Now(),
	}

	// On créé un buffer
	var buffer bytes.Buffer
	// On créé un new Encoder et on lui dit ou stocker l'encodage
	enc := gob.NewEncoder(&buffer)
	// On lui passe entry, qui va être stockée dans le buffer
	err := enc.Encode(entry)
	if err != nil {
		return err
	}

	// On peut faire un return direct car la fonction Update et Set peuvent retourner des erreurs
	return d.db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(word), buffer.Bytes())
	})
}

func (d *Dictionary) Get(word string) (Entry, error) {
	var entry Entry
	err := d.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(word))
		if err != nil {
			return err
		}
		entry, err = getEntry(item)
		return err
	})
	return entry, err
}

// List retrieves all the dictionnary content.
// []string is an alphabetically sorted array with the words
// [string]Entry is a map of the words and their definition
func (d *Dictionary) List() ([]string, map[string]Entry, error) {
	entries := make(map[string]Entry)
	err := d.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		// Va prendre en avance des groupe de 10 par 10, dépend de la taille
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		defer it.Close()
		// On le met à 0; tant qu'il est valide; on continue
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			entry, err := getEntry(item)
			if err != nil {
				return err
			}
			entries[entry.Word] = entry
		}
		return nil
	})
	return sortedKeys(entries), entries, err
}

func (d *Dictionary) Remove(word string) error {
	return d.db.Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte(word))
	})
}

func sortedKeys(entries map[string]Entry) []string {
	keys := make([]string, len(entries))
	for key := range entries {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func getEntry(item *badger.Item) (Entry, error) {
	var entry Entry
	// Déclaration buffer bytes qui contiendra ce qui a été lu par l'item
	var buffer bytes.Buffer
	// On récupe la valeur
	err := item.Value(func(val []byte) error {
		_, err := buffer.Write(val)
		return err
	})
	// Décodage du buffer, qui prend l'adresse du buffer
	dec := gob.NewDecoder(&buffer)
	err = dec.Decode(&entry)
	return entry, err
}
