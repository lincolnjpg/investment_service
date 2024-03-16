package repositories_test

import (
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/lincolnjpg/investment_service/internal/adapters/repositories"
	"github.com/lincolnjpg/investment_service/internal/infra"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var db *pgx.Conn
var repo repositories.UserRepository
var err error

func TestRepositories(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Repositories Suite")
}

var _ = BeforeSuite(func() {
	db, err = infra.NewPostgres(infra.DBConnParams{
		Host:     "localhost",
		UserName: "postgres",
		Password: "example",
		Database: "postgres",
		Port:     5433,
	})

	Expect(err).To(Succeed())

	repo = repositories.NewUserRepository(db)
})
