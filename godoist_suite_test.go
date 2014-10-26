package godoist_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGodoist(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Godoist Suite")
}
