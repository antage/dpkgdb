package dpkgdb

import (
	"errors"
	"fmt"
	"strings"
)

type Want int

const (
	UNKNOWN Want = iota + 1
	INSTALL
	HOLD
	REMOVE
	PURGE
)

func (v Want) String() string {
	switch v {
	case UNKNOWN:
		return "Unknown"
	case INSTALL:
		return "Install"
	case HOLD:
		return "Hold"
	case REMOVE:
		return "Remove"
	case PURGE:
		return "Purge"
	default:
		panic("Unknown action")
	}
}

type ErrorFlag int

const (
	OK ErrorFlag = iota + 1
	REINSTALL_REQUIRED
)

func (v ErrorFlag) String() string {
	switch v {
	case OK:
		return "OK"
	case REINSTALL_REQUIRED:
		return "Reinst-required"
	default:
		panic("Unknown error flag")
	}
}

type Status int

const (
	NOT_INSTALLED Status = iota + 1
	CONFIG_FILES
	HALF_INSTALLED
	UNPACKED
	HALF_CONFIGURED
	TRIGGERS_AWAITING
	TRIGGERS_PENDING
	INSTALLED
)

func (v Status) String() string {
	switch v {
	case NOT_INSTALLED:
		return "Not-installed"
	case CONFIG_FILES:
		return "Config-files"
	case HALF_INSTALLED:
		return "Half-installed"
	case UNPACKED:
		return "Unpacked"
	case HALF_CONFIGURED:
		return "Half-configured"
	case TRIGGERS_AWAITING:
		return "Triggers-awaiting"
	case TRIGGERS_PENDING:
		return "Triggers-pending"
	case INSTALLED:
		return "Installed"
	default:
		panic("Unknown status")
	}
}

type Package interface {
	Name() string

	Want() Want
	ErrorFlag() ErrorFlag
	Status() Status
}

type packageT struct {
	name   string
	want   Want
	eflag  ErrorFlag
	status Status
}

// It returns package name.
func (pkg packageT) Name() string {
	return pkg.name
}

// It returns package desired action.
func (pkg packageT) Want() Want {
	return pkg.want
}

// It returns package error flag.
func (pkg packageT) ErrorFlag() ErrorFlag {
	return pkg.eflag
}

// It returns package status.
func (pkg packageT) Status() Status {
	return pkg.status
}

func (pkg *packageT) parseStatus(st string) error {
	parseErr := errors.New(fmt.Sprintf("Can't parse status field for package '%s'", pkg.name))

	parts := strings.Split(st, " ")
	if len(parts) != 3 {
		return parseErr
	}

	switch parts[0] {
	case "unknown":
		pkg.want = UNKNOWN
	case "install":
		pkg.want = INSTALL
	case "hold":
		pkg.want = HOLD
	case "deinstall":
		pkg.want = REMOVE
	case "purge":
		pkg.want = PURGE
	default:
		return parseErr
	}

	switch parts[1] {
	case "ok":
		pkg.eflag = OK
	case "reinstreq":
		pkg.eflag = REINSTALL_REQUIRED
	default:
		return parseErr
	}

	switch parts[2] {
	case "not-installed":
		pkg.status = NOT_INSTALLED
	case "config-files":
		pkg.status = CONFIG_FILES
	case "half-installed":
		pkg.status = HALF_INSTALLED
	case "unpacked":
		pkg.status = UNPACKED
	case "half-configured":
		pkg.status = HALF_CONFIGURED
	case "triggers-awaited":
		pkg.status = TRIGGERS_AWAITING
	case "triggers-pending":
		pkg.status = TRIGGERS_PENDING
	case "installed":
		pkg.status = INSTALLED
	default:
		return parseErr
	}

	return nil
}
