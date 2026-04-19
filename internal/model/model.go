package model

type Model interface {
	Marshal() ([]byte, error)
}
