package godoist_test

import (
	"errors"
	"net/http"

	"github.com/robdimsdale/godoist"
	fakeHTTP "github.com/robdimsdale/godoist/httphelper/fakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	c              godoist.Client
	fakeHTTPHelper *fakeHTTP.FakeHTTPHelper
	response       *http.Response
)

var _ = Describe("Godoist", func() {

	BeforeEach(func() {
		response = &http.Response{}
		fakeHTTPHelper = &fakeHTTP.FakeHTTPHelper{}
		godoist.HTTPHelper = fakeHTTPHelper
		fakeHTTPHelper.PostFormReturns(response, nil)

		client, err := godoist.NewClient("email", "password")
		Expect(err).NotTo(HaveOccurred())
		c = client
	})

	Context("With empty string as email", func() {
		It("fails to create client", func() {
			_, err := godoist.NewClient("", "password")
			Expect(err).To(HaveOccurred())
		})
	})

	Context("With empty string as password", func() {
		It("fails to create client", func() {
			_, err := godoist.NewClient("email", "")
			Expect(err).To(HaveOccurred())
		})
	})

	Context("With a valid email and password", func() {
		Context("When performing login", func() {

			Context("When posting form returns error", func() {
				BeforeEach(func() {
					fakeHTTPHelper.PostFormReturns(nil, errors.New("Error during login"))
				})
				It("fowards the error", func() {
					err := c.Login()
					Expect(err).To(HaveOccurred())
				})
			})

			Context("When getting response body as bytes returns an error", func() {
				BeforeEach(func() {
					fakeHTTPHelper.ResponseBodyAsBytesReturns(nil, errors.New("Error converting response body to bytes"))
				})
				It("forwards the error", func() {
					err := c.Login()
					Expect(err).To(HaveOccurred())
				})
			})

			Context("When response body is nil", func() {
				BeforeEach(func() {
					fakeHTTPHelper.ResponseBodyAsBytesReturns(nil, nil)
				})
				It("returns an error", func() {
					err := c.Login()
					Expect(err).To(HaveOccurred())
				})
			})

			Context("When response body is empty", func() {
				BeforeEach(func() {
					fakeHTTPHelper.ResponseBodyAsBytesReturns([]byte{}, nil)
				})
				It("returns an error", func() {
					err := c.Login()
					Expect(err).To(HaveOccurred())
				})
			})

			Context("When response body is \"LOGIN_ERROR\"", func() {
				BeforeEach(func() {
					fakeHTTPHelper.ResponseBodyAsBytesReturns([]byte("\"LOGIN_ERROR\""), nil)
				})
				It("returns an error", func() {
					err := c.Login()
					Expect(err).To(HaveOccurred())
				})
			})

			Context("When unmarshalling json returns error", func() {
				BeforeEach(func() {
					fakeHTTPHelper.ResponseBodyAsBytesReturns([]byte("Invalid Json"), nil)
				})

				It("forwards the error", func() {
					err := c.Login()
					Expect(err).To(HaveOccurred())
				})
			})

			Context("When login succeeds succesfully", func() {
				BeforeEach(func() {
					fakeBody := []byte("{\"api_token\": \"some-api-token\"}")
					fakeHTTPHelper.ResponseBodyAsBytesReturns(fakeBody, nil)
				})

				It("sets the token on the client", func() {
					c.Login()
					Expect(c.ApiToken()).To(Equal("some-api-token"))
				})

				It("does not return an error", func() {
					err := c.Login()
					Expect(err).ToNot(HaveOccurred())
				})
			})
		})
	})

})
