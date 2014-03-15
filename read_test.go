package dpkgdb

import (
	"os/exec"
	"strings"
	"testing"
)

func TestArchitectures(t *testing.T) {
	db, err := Read()
	if err != nil {
		t.Fatal(err)
	}

	out, err := exec.Command("dpkg", "--print-architecture").Output()
	if err != nil {
		t.Fatal(err)
	}

	archs := strings.Split(string(out), "\n")
	for _, dirtyArch := range archs {
		arch := strings.TrimSpace(dirtyArch)
		if !db.HasArchitecture(arch) {
			t.Errorf("Expect DB has architecture %v", arch)
		}
	}

	out, err = exec.Command("dpkg", "--print-foreign-architectures").Output()
	if err != nil {
		t.Fatal(err)
	}

	archs = strings.Split(string(out), "\n")
	for _, dirtyArch := range archs {
		arch := strings.TrimSpace(dirtyArch)
		if !db.HasArchitecture(arch) {
			t.Errorf("Expect DB has architecture %v", arch)
		}
	}
}

func TestPackages(t *testing.T) {
	db, err := Read()
	if err != nil {
		t.Fatal(err)
	}

	pkg, ok := db.Package("dpkg")
	if !ok {
		t.Fatal("Can't found package 'dpkg'")
	}
	if pkg.Name() != "dpkg" {
		t.Errorf("Expect package name %v but got %v", "dpkg", pkg.Name())
	}
	if pkg.Want() != INSTALL {
		t.Errorf("Expect package desired action '%s' but got '%s'", "Install", pkg.Want())
	}
	if pkg.ErrorFlag() != OK {
		t.Errorf("Expect package error flag '%s' but got '%s'", "OK", pkg.ErrorFlag())
	}
	if pkg.Status() != INSTALLED {
		t.Errorf("Expect package status '%s' but got '%s'", "Installed", pkg.Status())
	}
}
