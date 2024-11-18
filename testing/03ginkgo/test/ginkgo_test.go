package newginkgo_test

// run "ginkgo generate set"

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestADTSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ADT Suite")
}
