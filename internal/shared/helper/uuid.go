package helper

import (
	"github.com/google/uuid"
)

type UuidGenerator interface {
	GenerateUUID() string
}

type uuidGenerator struct{}

func NewUuidGenerator() UuidGenerator {
	return &uuidGenerator{}
}

func (u *uuidGenerator) GenerateUUID() string {
	id := uuid.New()
	return id.String()
}
