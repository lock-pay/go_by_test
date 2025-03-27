package maps

const (
	ErrNotFound = DictionaryErr("can't")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionnary) Search(key string) (string, error) {
	definition, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionnary) Add(key, definition string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = definition
		return nil
	case nil:
		return ErrNotFound
	default:
		return err
	}

}

func (d Dictionnary) Update(key, definition string) error {
	_, err := d.Search(key)

	switch err {
	case nil:
		d[key] = definition
		return nil
	default:
		return err
	}
}

func (d Dictionnary) Delete(key string) error {
	_, err := d.Search(key)

	switch err {
	case nil:
		delete(d, key)
		return nil
	default:
		return err
	}
}

type Dictionnary map[string]string
