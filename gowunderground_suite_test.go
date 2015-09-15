package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGowunderground(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gowunderground Suite")
}
