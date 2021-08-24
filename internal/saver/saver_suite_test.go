package saver_test

import (
	"testing"

	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

func TestSaver(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Saver Suite")
}

// Declarations for Ginkgo DSL
var GinkgoT = ginkgo.GinkgoT
var RunSpecs = ginkgo.RunSpecs
var Fail = ginkgo.Fail
var Describe = ginkgo.Describe
var Context = ginkgo.Context
var It = ginkgo.It
var BeforeEach = ginkgo.BeforeEach

// Declarations for Gomega DSL
var RegisterFailHandler = gomega.RegisterFailHandler
var Expect = gomega.Expect

// Declarations for Gomega Matchers
var Equal = gomega.Equal
var BeNil = gomega.BeNil
