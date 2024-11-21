package pkg

import "encoding/json"

func JsonMarshal(i interface{}) (string, error) {
	v, err := json.Marshal(i)
	if err != nil {
		return "", err
	}
	return string(v), nil
}

func JsonUnMarshal(value string, i interface{}) error {
	data := []byte(value)
	err := json.Unmarshal(data, i)
	if err != nil {
		return err
	}
	return nil
}
