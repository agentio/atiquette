// Code generated by cmd/lexgenlite (see Makefile's lexgen); DO NOT EDIT.

package atproto

// schema: com.atproto.identity.resolveDid

import (
	"context"

	"github.com/agentio/atiquette/pkg/xrpc"
)

// IdentityResolveDid_Output is the output of a com.atproto.identity.resolveDid call.
type IdentityResolveDid_Output struct {
	// didDoc: The complete DID document for the identity.
	DidDoc interface{} `json:"didDoc" cborgen:"didDoc"`
}

// IdentityResolveDid calls the XRPC method "com.atproto.identity.resolveDid".
//
// did: DID to resolve.
func IdentityResolveDid(ctx context.Context, c xrpc.Client, did string) (*IdentityResolveDid_Output, error) {
	var out IdentityResolveDid_Output

	params := map[string]interface{}{
		"did": did,
	}
	if err := c.Do(ctx, xrpc.Query, "", "com.atproto.identity.resolveDid", params, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
