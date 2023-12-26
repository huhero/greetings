package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Hugo"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Hugo")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Hugo") = %q, %v, quiere coinciencia para %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v,  quiere "", error`, msg, err)
	}
}
