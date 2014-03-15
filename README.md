# dpkgdb

dpkg database API for Go.

## Installation

``` sh
go get github.com/antage/dpkgdb
```

## Usage

``` go
import (
    "fmt"
    "github.com/antage/dpkgdb"
    "os"
    "strings"
)

func main() {
    db, err := dpkgdb.Read()
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }

    fmt.Printf("Architectures: %s\n", strings.Join(db.Architectures(), ", "))

    // Is amd64 supported?
    if db.HasArchitecture("amd64") {
        fmt.Println("amd64 is supported")
    }

    // query package information
    pkg, found := db.Package("dpkg")
    if !found {
        fmt.Println("'dpkg' package is not found")
        return
    }

    fmt.Printf("Package name: %s\n", pkg.Name())
    fmt.Printf("Package status: %s\n", pkg.Status())

    // Error flag ('OK' or 'Reinstall required')
    fmt.Printf("Package error flag: %s\n", pkg.ErrorFlag())

    // Desired action ('Install', 'Remove', 'Purge', etc)
    fmt.Printf("Package desired action: %s\n", pkg.Want())
}
```
