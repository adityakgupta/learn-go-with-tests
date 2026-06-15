package maps

type Dictionary map[string]string

const (
	ErrNotfound = DictionaryErr("not found")
	ErrExist    = DictionaryErr("exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(s string) (string, error) {
	res, ok := d[s]
	if !ok {
		return res, ErrNotfound
	}

	return res, nil
}

func (d Dictionary) Add(key, val string) (res error) {
	_, err := d.Search(key)

	switch err {
	case ErrNotfound:
		d[key] = val
	case nil:
		res = ErrExist
	default:
		res = err
	}

	return
}

func (d Dictionary) Update(key, val string) (res error) {
	_, res = d.Search(key)

	if res == nil {
		d[key] = val
	}

	return
}

func (d Dictionary) Delete(key string) (res error) {
	_, res = d.Search(key)

	if res == nil {
		delete(d, key)
	}

	return
}
