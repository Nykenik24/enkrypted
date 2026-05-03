package id

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"regexp"

	"time"

	"gorm.io/gorm"
)

type IDModel struct {
	ID        string `gorm:"primaryKey;size:32"` // if id_len changes, update size parameter
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (i *IDModel) BeforeCreate(tx *gorm.DB) (err error) {
	if i.ID == "" {
		i.ID = string(generate())
	} else {
		if !validate(i.ID) {
			return fmt.Errorf("invalid ID")
		}
	}

	return nil
}

type ID struct {
	IDModel
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const id_len = 32

var idRegex = regexp.MustCompile(fmt.Sprintf(`[a-zA-Z0-9]{%d}`, id_len))

func validate(id string) bool {
	return len(id) == id_len && idRegex.Match([]byte(id))
}

func newID(contents string) *ID {
	return &ID{IDModel: IDModel{ID: contents}}
}

func fromBytes(str []byte) *ID {
	id := make([]byte, id_len)
	copy(id, str)
	return newID(string(id))
}

func FromString(str string) *ID {
	return fromBytes([]byte(str))
}

func generate() []byte {
	result := make([]byte, id_len)
	max := big.NewInt(int64(len(charset)))

	for i := range result {
		num, _ := rand.Int(rand.Reader, max)
		result[i] = charset[num.Int64()]
	}

	return result
}

func RandomID() *ID {
	result := generate()
	i := fromBytes(result)
	return i
}

func (i *ID) Bytes() []byte {
	return []byte(i.ID)
}

func (i *ID) String() string {
	if i == nil {
		return ""
	}
	return i.ID
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

	i.ID = s
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
