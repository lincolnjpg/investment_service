package repositories_test

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/lincolnjpg/investment_service/internal/domain"
)

var _ = Describe("user creation", func() {
	It("creates a single in PostgreSQL", func() {
		var count int

		_, err := repo.Create(
			context.Background(),
			domain.CreateUserInput{
				Name:            "user_1",
				InvestorProfile: "moderate",
			},
		)

		Expect(err).To(Succeed())

		row := db.QueryRow(
			context.Background(),
			`
				SELECT COUNT(*) FROM users;
			`,
		)

		Expect(row.Scan(&count)).To(Succeed())
		Expect(count).To(BeEquivalentTo(1))
	})

	It("creates a user and returns a non empty ID that represents the created user", func() {
		user, err := repo.Create(
			context.Background(),
			domain.CreateUserInput{
				Name:            "user_1",
				InvestorProfile: "moderate",
			},
		)

		Expect(err).To(Succeed())
		Expect(user.ID).NotTo(Equal(""))
	})

	It("creates a user and returns it to the caller", func() {
		var id, name, investor_profile string

		user, err := repo.Create(
			context.Background(),
			domain.CreateUserInput{
				Name:            "user_1",
				InvestorProfile: "moderate",
			},
		)

		Expect(err).To(Succeed())

		row := db.QueryRow(
			context.Background(),
			`
				SELECT * FROM users
				WHERE id = $1;
			`,
			user.ID,
		)

		Expect(row.Scan(&id, &name, &investor_profile)).To(Succeed())
		Expect(user.ID).To(Equal(id))
		Expect(user.Name).To(Equal(name))
		Expect(user.InvestorProfile).To(BeEquivalentTo(investor_profile))
	})
})

var _ = BeforeEach(func() {
	cleanUserTable()
})

var _ = AfterEach(func() {
	cleanUserTable()
})

func cleanUserTable() {
	db.Exec(
		context.Background(),
		`
			DELETE FROM users;
		`,
	)
}