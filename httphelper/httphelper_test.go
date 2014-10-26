package httphelper_test

import (
	"bytes"
	"io"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/robdimsdale/godoist/httphelper"
)

type nopCloser struct {
	io.Reader
}

func (nopCloser) Close() error { return nil }

var _ = Describe("Httphelper", func() {
	var a httphelper.HTTPHelper
	var resp *http.Response

	BeforeEach(func() {
		a = &httphelper.ActualHTTPHelper{}
		resp = &http.Response{}
	})

	Describe("Converting response body to bytes", func() {
		Context("When the response is nil", func() {
			BeforeEach(func() {
				resp = nil
			})
			It("returns an error", func() {
				_, err := a.ResponseBodyAsBytes(nil)
				Expect(err).To(HaveOccurred())
			})
		})

		Context("When the response body is nil", func() {
			It("returns an error", func() {
				_, err := a.ResponseBodyAsBytes(resp)
				Expect(err).To(HaveOccurred())
			})
		})

		Context("When reading the response body returns successfully", func() {
			BeforeEach(func() {
				resp.Body = nopCloser{bytes.NewBufferString("test")}
			})

			It("returns without error", func() {
				result, err := a.ResponseBodyAsBytes(resp)
				Expect(result).To(Equal([]byte("test")))
				Expect(err).ToNot(HaveOccurred())
			})
		})

	})

})
