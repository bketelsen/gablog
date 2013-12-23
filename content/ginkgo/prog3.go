package user_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"user"
)

var _ = Describe("User", func() {
	var u *user.User

	BeforeEach(func() {
		var err error
		u, err = user.New()
		Expect(err).NotTo(HaveOccurred())
	})

	Describe("Full Name", func() {
		Context("With a first and last name", func() {
			It("should concatenate the names with a ' '", func() {
				u.FirstName = "Peyton"
				u.LastName = "Manning"
				Expect(u.FullName()).To(Equal("Peyton Manning"))
			})
		})

		Context("With only a first name", func() {
			It("should return the first name", func() {
				u.FirstName = "Peyton"
				Expect(u.FullName()).To(Equal("Peyton"))
			})
		})

		Context("With only a last name", func() {
			It("should return the last name", func() {
				u.LastName = "Manning"
				Expect(u.FullName()).To(Equal("Manning"))
			})
		})

		Context("When first and last name are missing", func() {
			It("should return the empty string", func() {
				Expect(u.FullName()).To(BeEmpty())
			})
		})
	})
})
