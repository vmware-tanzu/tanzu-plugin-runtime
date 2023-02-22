package compatibility_tests_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCompatibilityTests(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CompatibilityTests Suite")
}
