package repositoryProducts

import (
	"context"
	"log"

	uuid "github.com/satori/go.uuid"
)

func (r *ProductsRepository) CheckIfAllItemsExist(ctx context.Context, uuids []uuid.UUID) error {
	log.Printf("productsRepository.CheckIfAllItemsExist")

	return nil
}
