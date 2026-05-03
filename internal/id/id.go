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

func fromBytes(str []byte) *ID {
	id := make(ID, id_len)
	copy(id, str)
	return &id
}

func FromString(str string) *ID {
	return fromBytes([]byte(str))
}

func RandomID() *ID {
	result := make([]byte, id_len)
	max := big.NewInt(int64(len(charset)))

	for i := range result {
		num, _ := rand.Int(rand.Reader, max)
		result[i] = charset[num.Int64()]
	}

	i := fromBytes(result)
	return i
}

func (i *ID) Bytes() []byte {
	return ([]byte)(*i)
}

func (i *ID) String() string {
	if i == nil {
		return ""
	}
	return string(*i)
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

func (i *ID) Short() []byte {
	return fmt.Appendf(i.Bytes()[:6], "...")
}

func (a *ID) Compare(b *ID) bool {
	return a.String() == b.String()
}

func (a *ID) CompareString(b string) bool {
	return a.String() == b
}
