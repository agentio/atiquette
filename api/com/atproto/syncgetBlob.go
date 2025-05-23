// Code generated by cmd/lexgenlite (see Makefile's lexgen); DO NOT EDIT.

package atproto

// schema: com.atproto.sync.getBlob

import (
	"bytes"
	"context"

	"github.com/agentio/atiquette/pkg/xrpc"
)

// SyncGetBlob calls the XRPC method "com.atproto.sync.getBlob".
//
// cid: The CID of the blob to fetch
// did: The DID of the account.
func SyncGetBlob(ctx context.Context, c xrpc.Client, cid string, did string) ([]byte, error) {
	buf := new(bytes.Buffer)

	params := map[string]interface{}{
		"cid": cid,
		"did": did,
	}
	if err := c.Do(ctx, xrpc.Query, "", "com.atproto.sync.getBlob", params, nil, buf); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
