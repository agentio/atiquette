// Code generated by cmd/lexgenlite (see Makefile's lexgen); DO NOT EDIT.

package atproto

// schema: com.atproto.identity.getRecommendedDidCredentials

import (
	"context"

	"github.com/agentio/atiquette/pkg/xrpc"
)

// IdentityGetRecommendedDidCredentials_Output is the output of a com.atproto.identity.getRecommendedDidCredentials call.
type IdentityGetRecommendedDidCredentials_Output struct {
	AlsoKnownAs []string `json:"alsoKnownAs,omitempty" cborgen:"alsoKnownAs,omitempty"`
	// rotationKeys: Recommended rotation keys for PLC dids. Should be undefined (or ignored) for did:webs.
	RotationKeys        []string     `json:"rotationKeys,omitempty" cborgen:"rotationKeys,omitempty"`
	Services            *interface{} `json:"services,omitempty" cborgen:"services,omitempty"`
	VerificationMethods *interface{} `json:"verificationMethods,omitempty" cborgen:"verificationMethods,omitempty"`
}

// IdentityGetRecommendedDidCredentials calls the XRPC method "com.atproto.identity.getRecommendedDidCredentials".
func IdentityGetRecommendedDidCredentials(ctx context.Context, c xrpc.Client) (*IdentityGetRecommendedDidCredentials_Output, error) {
	var out IdentityGetRecommendedDidCredentials_Output
	if err := c.Do(ctx, xrpc.Query, "", "com.atproto.identity.getRecommendedDidCredentials", nil, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
