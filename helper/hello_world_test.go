package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Tono")
	if result != "Hello Tono" {
		t.FailNow()
	}
	fmt.Println("TestHelloWorld Done")
}

func TestHelloWorldJean(t *testing.T) {
	result := HelloWorld("Jean")
	if result != "Hello Jean" {
		t.Fail()
	}
	fmt.Println("TestHelloWorldJean Done")
}

func TestHelloWorldError(t *testing.T) {
	result := HelloWorld("Jean")
	if result != "Hello Jean" {
		t.Error("Should be Hello Jean")
	}
	fmt.Println("Executed if fail")
}

func TestHelloWorldJeanFatal(t *testing.T) {
	result := HelloWorld("Jean")
	if result != "Hello Jean2" {
		t.Fatal("Should be Hello Jean")
	}
	fmt.Println("Not executed if fail")
}

func TestHelloWorldAssertion(t *testing.T) {
	result := HelloWorld("Jean")
	assert.Equal(t, "Hello Jean", result, "Should be Hello Jean")
	fmt.Println("Executed")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Jean")
	require.Equal(t, "Hello Jean", result, "Should be Hello Jean")
	fmt.Println("Not Executed")
}
