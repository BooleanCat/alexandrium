package router_test

import (
	"io"
	"net/http"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRouter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Router Suite")
}

func newRequest(method, url string, body io.Reader) *http.Request {
	request, err := http.NewRequest(method, url, body)
	ExpectWithOffset(1, err).NotTo(HaveOccurred())
	return request
}
