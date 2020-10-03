package router_test

import (
	"net/http"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRouter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Router Suite")
}

func httpGet(url string) *http.Response {
	response, err := http.Get(url)
	ExpectWithOffset(1, err).NotTo(HaveOccurred())
	return response
}
