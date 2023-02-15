package util

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"encoding/json"
	"io/ioutil"
	"log"
)

func SerializeToFile[T any](fn string, data T) error {
	out, err := json.Marshal(data)
	log.Printf("TRACE: DATA = %s", string(out))
	if err != nil {
		return err
	}
	return ioutil.WriteFile(fn, out, 0644)
}

func UnserializeFromFile[T any](fn string) (T, error) {
	var out T
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		return out, err
	}
	err = json.Unmarshal(data, &out)
	return out, nil
}

func ToGOB64[T any](m T) (string, error) {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	err := e.Encode(m)
	if err != nil {
		log.Printf(`failed gob Encode`, err)
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b.Bytes()), nil
}

func FromGOB64[T any](str string) (T, error) {
	m := new(T)
	by, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		log.Printf("ERR: failed base64 Decode: %s", err.Error())
		return *m, err
	}
	b := bytes.Buffer{}
	b.Write(by)
	d := gob.NewDecoder(&b)
	err = d.Decode(&m)
	if err != nil {
		log.Printf("ERR: failed gob Decode: %s", err.Error())
		return *m, err
	}
	return *m, nil
}
