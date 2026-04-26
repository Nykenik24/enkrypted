package id

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
)

type ID []byte

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const id_len = 32

func FromBytes(str []byte) (*ID, error) {
	if len(str) != id_len {
		return nil, fmt.Errorf("ID string is not %d characters long", id_len)
	}

	id := make(ID, id_len)
	copy(id, str)
	return &id, nil
}

func FromString(str string) (*ID, error) {
	return FromBytes([]byte(str))
}

func RandomID() *ID {
	result := make([]byte, id_len)
	max := big.NewInt(int64(len(charset)))

	for i := range result {
		num, _ := rand.Int(rand.Reader, max)
		result[i] = charset[num.Int64()]
	}

	i, _ := FromBytes(result)
	return i
}

func (i *ID) String() string {
	if i == nil {
		return ""
	}
	return (string)(*i)
}

func (i ID) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

func (i *ID) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	if i == nil {
		return errors.New("id: cannot unmarshal into nil pointer")
	}

	*i = ID(s)
	return nil
}
