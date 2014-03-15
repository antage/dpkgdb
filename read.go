package dpkgdb

import (
	"bufio"
	"io"
	"os"
	"path"
	"strings"
)

const (
	archFilename   = "arch"
	statusFilename = "status"
	defaultDir     = "/var/lib/dpkg"

	hdrPackage    = "Package:"
	hdrPackageLen = len(hdrPackage)
	hdrStatus     = "Status:"
	hdrStatusLen  = len(hdrStatus)
)

// Read reads dpkg database from default path (/var/lib/dpkg).
func Read() (*DB, error) {
	return ReadFromDir(defaultDir)
}

// ReadFromDir reads dpkg database from given path.
func ReadFromDir(dir string) (*DB, error) {
	db := newDB()

	err := db.readArchitecturesFromDir(dir)
	if err != nil {
		return nil, err
	}

	err = db.readStatusFromDir(dir)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (db *DB) addArchitecture(arch string) {
	db.Lock()
	defer db.Unlock()

	db.archs = append(db.archs, arch)
	db.archsIndex[arch] = struct{}{}
}

func (db *DB) readArchitecturesFromDir(dir string) error {
	archFile, err := os.Open(path.Join(dir, archFilename))
	if err != nil {
		return err
	}
	defer archFile.Close()

	bufArchFile := bufio.NewReader(archFile)
	err = nil
	for err == nil {
		var line string
		line, err = bufArchFile.ReadString('\n')

		db.addArchitecture(strings.TrimSpace(line))
	}
	if err != io.EOF {
		return err
	}

	return nil
}

func (db *DB) addPackage(pkg *packageT) {
	db.Lock()
	defer db.Unlock()

	db.packages = append(db.packages, pkg)
	db.packagesIndexByName[pkg.name] = pkg
}

func (db *DB) readStatusFromDir(dir string) error {
	statusFile, err := os.Open(path.Join(dir, statusFilename))
	if err != nil {
		return err
	}
	defer statusFile.Close()

	bufStatusFile := bufio.NewReader(statusFile)

	var pkg *packageT

	err = nil
	for err == nil {
		var line string

		line, err = bufStatusFile.ReadString('\n')
		tline := strings.TrimSpace(line)

		switch {
		case strings.HasPrefix(tline, hdrPackage):
			if pkg != nil {
				db.addPackage(pkg)
			}

			pkg = new(packageT)

			name := strings.TrimSpace(tline[hdrPackageLen:])
			pkg.name = name
		case pkg != nil && strings.HasPrefix(tline, hdrStatus):
			status_value := strings.TrimSpace(tline[hdrStatusLen:])
			pkg.parseStatus(status_value)
		default:
			// ignore this line
		}
	}
	if err != io.EOF {
		return err
	}
	if pkg != nil {
		db.addPackage(pkg)
	}

	return nil
}
