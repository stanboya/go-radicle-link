// This file is part of go-radicle-link, distributed under the GPLv3
// For full terms see the included LICENSE file.
package identity

type DocPayload interface{}

type ProjectDocPayload struct {
	// A short name
	name string
	// A slightly longer description (should fit in a headline)
	description string
	// The default branch, "master" is assumed for git repositories
	// if unspecified
	default_branch string
}

type UserDocPayload struct {
	// short name (nickname, handle), without any prefix such as the `@`
	// character
	name string
}
