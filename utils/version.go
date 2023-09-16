package utils

import (
	"github.com/blang/semver/v4"
)

// VersionGTE returns true if the a is greater than or equal to the e.
func VersionGTE(a, e string) bool {
	if e == "" {
		return true
	}

	if a == "" {
		return false
	}

	ev, err := semver.ParseTolerant(e)
	if err != nil {
		return false
	}

	av, err := semver.ParseTolerant(a)
	if err != nil {
		return false
	}

	return av.GTE(ev)
}

// VersionGT returns true if the a is greater than the e.
func VersionGT(a, e string) bool {
	if e == "" {
		return true
	}

	if a == "" {
		return false
	}

	ev, err := semver.ParseTolerant(e)
	if err != nil {
		return false
	}

	av, err := semver.ParseTolerant(a)
	if err != nil {
		return false
	}

	return av.GT(ev)
}

// VersionLTE returns true if the a is less than or equal to the e.
func VersionLTE(a, e string) bool {
	if a == "" {
		return true
	}

	if e == "" {
		return false
	}

	ev, err := semver.ParseTolerant(e)
	if err != nil {
		return false
	}

	av, err := semver.ParseTolerant(a)
	if err != nil {
		return false
	}

	return av.LTE(ev)
}

// VersionLT returns true if the a is less than the e.
func VersionLT(a, e string) bool {
	if a == "" {
		return true
	}

	if e == "" {
		return false
	}

	ev, err := semver.ParseTolerant(e)
	if err != nil {
		return false
	}

	av, err := semver.ParseTolerant(a)
	if err != nil {
		return false
	}

	return av.LT(ev)
}
