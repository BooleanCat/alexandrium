package acceptance_test

import (
	"net/http"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	binPath       = "../targets/alexandrium"
	serverAddress = "http://127.0.0.1:3000"
)

func TestAcceptance(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Acceptance Suite")
}

func httpGet(url string) *http.Response {
	response, err := http.Get(url)
	ExpectWithOffset(1, err).NotTo(HaveOccurred())
	return response
}
