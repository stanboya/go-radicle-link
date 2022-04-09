// This file is part of go-radicle-link, distributed under the GPLv3
// For full terms see the included LICENSE file.
package identity

type Delegation interface{}

type User struct {
	identity Identity
}

type ProjectDelegation int

const (
	KeyDelegationType ProjectDelegation = iota + 1
	UserDelegationType
)
