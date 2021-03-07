package maps

import "errors"

type Dictionary map[string]string

var (
	ErrCantAdd = errors.New("could not add word because word already exist")
	ErrNotFound = errors.New("could not find the word you were looking for")
	ErrNotExist = errors.New("update error")
	ErrDelete = errors.New("delete error")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition,nil
}

func (d Dictionary) add(word, definition string) error  {
	_, e := d.Search(word)

	switch e {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrCantAdd
	default:
		return e
	}
	return nil
}

func (d Dictionary) update(word, definition string) error  {
	_, e := d.Search(word)
	switch e {
	case ErrNotFound:
		return ErrNotExist
	case nil:
		d[word] = definition
	}
	return nil
}

func (d Dictionary) delete(word string) error {
	delete(d, word)
	return nil
}

