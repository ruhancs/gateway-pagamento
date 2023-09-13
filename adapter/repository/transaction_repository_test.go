//gerar o mock do repository: mockgen -destination=domain/repository/mock/mock.go -source=domain/repository/repository.go
//gerar o mock do repository: mockgen -destination=adapter/broker/mock/mock.go -source=adapter/broker/interface.go

package repository

import (
	"os"
	"testing"

	"github.com/ruhancs/gateway-pagamento/adapter/repository/fixture"
	"github.com/ruhancs/gateway-pagamento/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestTransactionRepositoryInsert(t *testing.T) {
	//buscar arquivos das migracoes
	migrationsDir := os.DirFS("fixture/sql")
	db := fixture.Up(migrationsDir)//subir o banco de dados
	defer fixture.Down(db, migrationsDir)

	repository := NewTransactionRepository(db)
	err := repository.Insert("1", "1", 15.8, entity.APPROVED, "")
	assert.Nil(t, err)
}