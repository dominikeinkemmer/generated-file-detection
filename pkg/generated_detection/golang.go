package generated_detection

import "regexp"

// isGoVendor checks if the blob is part of the Go vendor/ tree,
// not meant for humans in pull requests.
func (d detector) isGoVendor() bool {
	goVendorRe := regexp.MustCompile("vendor/((?!-)[-0-9A-Za-z]+(?<!-)\\.)+(com|edu|gov|in|me|net|org|fm|io)")

	return goVendorRe.MatchString(d.FileName.Base)
}

// isGoLock checks if the blob us a generated Go dep or glide lock file?
func (d detector) isGoLock() bool {
	goLockRe := regexp.MustCompile("(Gopkg|glide)\\.lock")

	return goLockRe.MatchString(d.FileName.Base)
}

// isGoDeps checks if the blob is part of Godeps/,
// which are not meant for humans in pull requests.
func (d detector) isGoDeps() bool {
	goDepsRe := regexp.MustCompile("Godeps/")

	return goDepsRe.MatchString(d.FileName.Base)
}

// isGoGenerated tries to detect auto generated files for golang
//
// We follow the following proposal by the golang community for this approach:
// https://github.com/golang/go/issues/13560#issuecomment-288457920
func (d detector) isGoGenerated() bool {
	// If any line contains a match for this string, it is a generated file
	re := regexp.MustCompile("^// Code generated .* DO NOT EDIT\\.$")

	for _, line := range d.Lines.All() {
		if re.MatchString(line) {
			return true
		}
	}

	return false
}
