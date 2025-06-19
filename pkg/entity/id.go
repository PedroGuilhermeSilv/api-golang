package entity

import (
	erros "github.com/PedroGuilhermeSilv/api-golang/pkg/erros"
	"github.com/google/uuid"
)

type ID = uuid.UUID

func NewID() ID {
	return ID(uuid.New())
}

func ParseID(id string) (ID, error) {
	idParsed, err := uuid.Parse(id)

	return erros.HandleError(ID(idParsed), err)
}
