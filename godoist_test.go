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
	fakeHTTPHelper *fakeHTTP.FakeHTTPHelper
	response       *http.Response
)

var _ = Describe("Godoist", func() {
	BeforeEach(func() {
		response = &http.Response{}
		fakeHTTPHelper = &fakeHTTP.FakeHTTPHelper{}
		godoist.HTTPHelper = fakeHTTPHelper
		fakeHTTPHelper.PostFormReturns(response, nil)
	})

	Context("When creating client", func() {
		Context("With empty string as email", func() {
			It("returns an error", func() {
				_, err := godoist.NewClient("", "password")
				Expect(err).To(HaveOccurred())
			})
		})

		Context("With empty string as password", func() {
			It("returns an error", func() {
				_, err := godoist.NewClient("email", "")
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Context("With a successfully created client", func() {
		var c godoist.Client
		BeforeEach(func() {
			client, err := godoist.NewClient("email", "password")
			Expect(err).NotTo(HaveOccurred())
			c = client
		})

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

				It("returns without error", func() {
					err := c.Login()
					Expect(err).ToNot(HaveOccurred())
				})
			})
		})

		Context("When performing ping", func() {
			Context("When posting form returns error", func() {
				BeforeEach(func() {
					fakeHTTPHelper.PostFormReturns(nil, errors.New("Error during ping"))
				})
				It("fowards the error", func() {
					_, err := c.Ping()
					Expect(err).To(HaveOccurred())
				})
			})

			Context("When response status code is http.StatusUnauthorized", func() {
				BeforeEach(func() {
					response.StatusCode = http.StatusUnauthorized
				})
				It("returns an error", func() {
					_, err := c.Ping()
					Expect(err).To(HaveOccurred())
				})
			})

			Context("When response status code is http.StatusForbidden", func() {
				BeforeEach(func() {
					response.StatusCode = http.StatusForbidden
				})
				It("returns an error", func() {
					_, err := c.Ping()
					Expect(err).To(HaveOccurred())
				})
			})

			Context("When getting response body as bytes returns an error", func() {
				BeforeEach(func() {
					fakeHTTPHelper.ResponseBodyAsBytesReturns(nil, errors.New("Error converting response body to bytes"))
				})
				It("forwards the error", func() {
					_, err := c.Ping()
					Expect(err).To(HaveOccurred())
				})
			})

			Context("When ping succeeds succesfully", func() {
				BeforeEach(func() {
					fakeHTTPHelper.ResponseBodyAsBytesReturns([]byte("ping-result"), nil)
				})
				It("returns without error", func() {
					result, err := c.Ping()
					Expect(err).ToNot(HaveOccurred())
					Expect(result).To(Equal("ping-result"))
				})
			})
		})
	})

})
