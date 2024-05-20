package gofile_test

import (
	"testing"

	"github.com/dvwzj/gofile"
)

func TestNewClient(t *testing.T) {
	client, err := gofile.NewClient()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if client == nil {
		t.Fatalf("unexpected nil client")
	}
}

func TestNewClientWithToken(t *testing.T) {
	client, err := gofile.NewClient(gofile.WithToken("7JyBbtDTF7yakfcfmTWYlXNTL4j5r9HV"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if client == nil {
		t.Fatalf("unexpected nil client")
	}
}

func TestNewClientWithNewGuestAccount(t *testing.T) {
	client, err := gofile.NewClient(gofile.WithNewGuestAccount)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if client == nil {
		t.Fatalf("unexpected nil client")
	}
}
