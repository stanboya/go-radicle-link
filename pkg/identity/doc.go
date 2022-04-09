// This file is part of go-radicle-link, distributed under the GPLv3
// For full terms see the included LICENSE file.
package identity

type Doc struct {
	replaces    Revision
	payload     DocPayload
	delegations []Delegation
}
