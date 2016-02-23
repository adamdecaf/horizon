package storage

import (
	"testing"
	"github.com/adamdecaf/horizon/utils"
)

func TestHostnameReadWrite(t *testing.T) {
	value := utils.RandString(20)
	empty, err := SearchHostnameByValue(value)

	if err == nil {
		t.Fatalf("no error when we expected one -- row should not exist")
	}

	nothing := Hostname{}
	if empty != nothing {
		t.Fatal("found hostname some how when we didn't expect to find any")
	}

	id := utils.RandString(20)
	hostname := Hostname{id, "host"}

	if written := WriteHostname(hostname); written != nil {
		t.Fatalf("error when writing hostname id=%s, value=%s, err=%s", id, value, *written)
	}

	found, err := SearchHostnameByValue("host")

	if err != nil {
		t.Fatalf("error finding hostname that should exist id=%s, err=%s\n", id, err)
	}

	if found.Id != hostname.Id || found.Value != hostname.Value {
		t.Fatalf("hostname don't match (written=%s) (found=%s)", hostname, found)
	}
}
