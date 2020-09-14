package router_test

import (
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/BooleanCat/alexandrium/router"
)

var _ = Describe("Router", func() {
	var recorder *httptest.ResponseRecorder

	BeforeEach(func() {
		recorder = httptest.NewRecorder()
	})

	Describe("/ping", func() {
		It("responds with 204 No Content", func() {
			http.HandlerFunc(router.Ping).ServeHTTP(recorder, newRequest(http.MethodGet, "/ping", nil))
			Expect(recorder.Code).To(Equal(http.StatusNoContent))
		})
	})
})
