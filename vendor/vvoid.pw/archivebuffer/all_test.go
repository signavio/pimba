package archivebuffer

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"vvoid.pw/randomizer"
)

func TestArchive(t *testing.T) {
	testfilePath, err := createTestFile()
	if err != nil {
		t.Errorf("Some errors occurred when creating the test file. %v", err)
	}
	tarBuf, err := NewTarballBuffer(testfilePath)
	if err != nil {
		t.Errorf("Some errors occurred when archiving. %v", err)
	}
	unarchivePath := "/tmp/unarchiving-test"
	os.Mkdir(unarchivePath, 0777)
	err = UntarToFile(tarBuf, unarchivePath)
	if err != nil {
		t.Errorf("Some errors occurred when unarchiving. %v", err)
	}
	want := "hello, foobar!\n"
	sl := strings.Split(testfilePath, "/")
	got, err := ioutil.ReadFile(unarchivePath + "/" + sl[len(sl)-1])
	if err != nil {
		t.Errorf("Some errors occurred when reading the file. %v", err)
	}
	if string(got) != want {
		t.Errorf("Expected '%v' got '%v'", want, got)
	}
	os.Remove(testfilePath)
	os.RemoveAll(unarchivePath)
}

func TestCompress(t *testing.T) {
	testfilePath, err := createTestFile()
	if err != nil {
		t.Errorf("Some errors occurred when creating the test file. %v", err)
	}
	tarBuf, err := NewTarballBuffer(testfilePath)
	if err != nil {
		t.Errorf("Some errors occurred when archiving. %v", err)
	}
	gzipBuf, err := NewGzipBuffer(tarBuf)
	if err != nil {
		t.Errorf("Some errors occurred when compressing. %v", err)
	}
	ungzipBuf, err := UngzipToBuffer(gzipBuf)
	if err != nil {
		t.Errorf("Some errors occurred when uncompressing. %v", err)
	}
	unarchivePath := "/tmp/unarchiving-test"
	os.Mkdir(unarchivePath, 0777)
	err = UntarToFile(ungzipBuf, unarchivePath)
	if err != nil {
		t.Errorf("Some errors occurred when unarchiving. %v", err)
	}
	want := "hello, foobar!\n"
	sl := strings.Split(testfilePath, "/")
	got, err := ioutil.ReadFile(unarchivePath + "/" + sl[len(sl)-1])
	if err != nil {
		t.Errorf("Some errors occurred when reading the file. %v", err)
	}
	if string(got) != want {
		t.Errorf("Expected '%v' got '%v'", want, got)
	}
	os.Remove(testfilePath)
	os.RemoveAll(unarchivePath)
}

func createTestFile() (string, error) {
	path := fmt.Sprintf("/tmp/foobar-%v", randomizer.GenerateRandomString(8))
	d := []byte("hello, foobar!\n")
	err := ioutil.WriteFile(path, d, 0644)
	if err != nil {
		return "", err
	}
	return path, nil
}
