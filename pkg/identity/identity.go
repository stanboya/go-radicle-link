// This file is part of go-radicle-link, distributed under the GPLv3
// For full terms see the included LICENSE file.
package identity

import "crypto"

type Signature = crypto.Hash

type ContentId = crypto.Hash

type Identity struct {
	id         ContentId
	root       Revision
	revision   Revision
	doc        Doc
	signatures map[crypto.PublicKey]Signature
}
