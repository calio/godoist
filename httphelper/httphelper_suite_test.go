package httphelper_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestHttphelper(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Httphelper Suite")
}
