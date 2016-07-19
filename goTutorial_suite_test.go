package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGoTutorial(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoTutorial Suite")

}
var _ = Describe("Title", func(){
		var t *template.Template


}

It("Has a title", func(done Done){
	go func(){
		defer GinkgoRecover()
		Ω(doSomething()).Should(BeTrue())

		close(done)
}()
})
