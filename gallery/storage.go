package gallery

import (
	"encoding/json"
	"errors"
	"io/fs"
	"log"
	"os"
)

func OpenOrCreate(fn string) (*os.File, error) {
	f, err := os.Open(fn)
	if errors.Is(err, fs.ErrNotExist) {
		f, err = os.Create(fn)
		if err != nil {
			return nil, err
		}
		log.Println("Created file", fn)
		_, err = f.Write([]byte("{}\n"))
		if err != nil {
			return nil, err
		}
		_, err = f.Seek(0, 0)
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}
	return f, nil
}

func DecodeJSON(fn string, target interface{}) error {
	f, err := OpenOrCreate(fn)
	if err != nil {
		return err
	}
	defer f.Close()
	d := json.NewDecoder(f)
	return d.Decode(target)
}

func EncodeJSON(fn string, data interface{}) error {
	f, err := os.Create(fn)
	if err != nil {
		return err
	}
	defer f.Close()
	e := json.NewEncoder(f)
	return e.Encode(data)
}

var Users = map[string]*User{}
var Images = map[uint32]*Image{}

func LoadUsers() error {
	return DecodeJSON("data/users.json", &Users)
}

func LoadImages() error {
	return DecodeJSON("data/images.json", &Images)
}

func SaveUsers() error {
	return EncodeJSON("data/users.json", Users)
}

func SaveImages() error {
	return EncodeJSON("data/images.json", Images)
}
