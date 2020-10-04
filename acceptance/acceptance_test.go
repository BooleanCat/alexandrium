package acceptance_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/BooleanCat/alexandrium/books"
)

var _ = Describe("Acceptance", func() {
	var cmd *exec.Cmd

	BeforeEach(func() {
		cmd = exec.Command(bin_path)
		cmd.Stdout = GinkgoWriter
		cmd.Stderr = GinkgoWriter
		Expect(cmd.Start()).To(Succeed())
		Eventually(ping("http://127.0.0.1:3000")).Should(Succeed())
	})

	AfterEach(func() {
		Expect(cmd.Process.Signal(os.Interrupt)).To(Succeed())
		Expect(cmd.Wait()).To(Succeed())
	})

	It("does nothing", func() {
		Expect(true).To(BeTrue())
	})

	Describe("GET /books/9781788547383", func() {
		It("responds with the correct book data", func() {
			response, err := http.Get("http://127.0.0.1:3000/books/9781788547383")
			Expect(err).NotTo(HaveOccurred())
			defer closeIgnoreError(response.Body)

			Expect(response.StatusCode).To(Equal(http.StatusOK))

			var book books.Book
			Expect(json.NewDecoder(response.Body).Decode(&book)).To(Succeed())
			Expect(book.ID).To(Equal("76341e07-911c-44fd-aafa-13b43daf3494"))
		})
	})
})

func ping(url string) func() error {
	return func() error {
		response, err := http.Get(url + "/ping")
		if err != nil {
			return err
		}

		if response.StatusCode == http.StatusNoContent {
			return nil
		}

		return fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}
}

func closeIgnoreError(c io.Closer) {
	_ = c.Close()
}
