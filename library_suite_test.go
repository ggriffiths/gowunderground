package main

import (
	"testing"

	"github.com/ggriffiths/gowunderground/fakes"

	. "github.build.ge.com/predix-data-services/gateway/Godeps/_workspace/src/github.com/onsi/ginkgo"
	. "github.build.ge.com/predix-data-services/gateway/Godeps/_workspace/src/github.com/onsi/gomega"
)

var (
	report fakes.Report
)

func TestLibrary(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Library Suite")
}

var _ = BeforeSuite(func() {

})
