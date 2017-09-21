package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/signavio/pimba/pkg/api"
	"github.com/signavio/pimba/pkg/push"
)

const (
	TestStorage   = "/tmp/pimba-storage-testing"
	TestSecret    = "123testing"
	TestPort      = 9398
	TestServerURL = "localhost:9398"
	TestBucket    = "footest"
	TestFilesPath = "/tmp/pimba-test-files"
)

func init() {
	go api.Serve(TestPort, TestStorage, TestSecret)
}

func TestPushURL(t *testing.T) {
	if err := createTestStorage(); err != nil {
		clean()
		t.Errorf("Error: %v", err)
	}
	if err := generateTestFiles(); err != nil {
		clean()
		t.Errorf("Error: %v", err)
	}

	want := "localhost:9398/footest"
	got, err := push.PushFiles(TestFilesPath, TestServerURL, TestBucket, "")
	if err != nil {
		clean()
		t.Errorf("Error: %v", err)
	}

	if got[0] != want {
		clean()
		t.Errorf("Expected '%v' got '%v'", want, got[0])
	}

	clean()
}

func TestPushToken(t *testing.T) {
	if err := createTestStorage(); err != nil {
		clean()
		t.Errorf("Error: %v", err)
	}
	if err := generateTestFiles(); err != nil {
		clean()
		t.Errorf("Error: %v", err)
	}

	resp, err := push.PushFiles(TestFilesPath, TestServerURL, TestBucket, "")
	if err != nil {
		clean()
		t.Errorf("Error: %v", err)
	}

	want := resp[1]
	got, err := push.PushFiles(TestFilesPath, TestServerURL, TestBucket, want)
	if err != nil {
		clean()
		t.Errorf("Error: %v", err)
	}

	if got[1] != want {
		clean()
		t.Errorf("Expected '%v' got '%v'", want, got[1])
	}

	clean()
}

func TestPushAccess(t *testing.T) {
	if err := createTestStorage(); err != nil {
		clean()
		t.Errorf("Error: %v", err)
	}
	if err := generateTestFiles(); err != nil {
		clean()
		t.Errorf("Error: %v", err)
	}

	_, err := push.PushFiles(TestFilesPath, TestServerURL, TestBucket, "")
	if err != nil {
		clean()
		t.Errorf("Error: %v", err)
	}

	fileURL := "http://localhost:9398/footest/foo-1"
	resp, err := http.Get(fileURL)
	if err != nil {
		clean()
		t.Errorf("Error: %v", err)
	}
	defer resp.Body.Close()

	want := "FOOO"
	got, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		clean()
		t.Errorf("Error: %v", err)
	}
	if string(got) != want {
		clean()
		t.Errorf("Expected '%v' got '%v'", want, got)
	}

	clean()
}

func createTestStorage() error {
	return os.MkdirAll(TestStorage, 0755)
}

func generateTestFiles() error {
	os.MkdirAll(TestFilesPath, 0755)

	for n := 1; n <= 5; n++ {
		f, err := os.Create(fmt.Sprintf("%v/foo-%v", TestFilesPath, n))
		if err != nil {
			return err
		}
		defer f.Close()
		f.WriteString("FOOO")
		f.Sync()
	}

	return nil
}

func clean() {
	os.RemoveAll(TestStorage)
	os.RemoveAll(TestFilesPath)
}
