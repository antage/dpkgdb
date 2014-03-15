package dpkgdb

// Architectures returns all supported architectures.
func (db *DB) Architectures() []string {
	db.RLock()
	defer db.RUnlock()

	archs := make([]string, len(db.archs))
	copy(archs, db.archs)
	return archs
}

// HasArchitecture checks that arch architectures is supported.
// It returns true if arch is supported architecture.
func (db *DB) HasArchitecture(arch string) bool {
	db.RLock()
	defer db.RUnlock()

	_, found := db.archsIndex[arch]
	return found
}

// Package finds package by name.
// It returns package information and 'found' flag.
func (db *DB) Package(name string) (pkg Package, found bool) {
	db.RLock()
	defer db.RUnlock()

	pkg, found = db.packagesIndexByName[name]
	return
}

// Packages returns all packages.
func (db *DB) Packages() []Package {
	db.RLock()
	defer db.RUnlock()

	pkgs := make([]Package, len(db.packages))
	for i := range db.packages {
		pkgs[i] = db.packages[i]
	}

	return pkgs
}
