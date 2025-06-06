// Code generated by cmd/lexgenlite (see Makefile's lexgen); DO NOT EDIT.

package atproto

// schema: com.atproto.server.deactivateAccount

import (
	"context"

	"github.com/agentio/atiquette/pkg/xrpc"
)

// ServerDeactivateAccount_Input is the input argument to a com.atproto.server.deactivateAccount call.
type ServerDeactivateAccount_Input struct {
	// deleteAfter: A recommendation to server as to how long they should hold onto the deactivated account before deleting.
	DeleteAfter *string `json:"deleteAfter,omitempty" cborgen:"deleteAfter,omitempty"`
}

// ServerDeactivateAccount calls the XRPC method "com.atproto.server.deactivateAccount".
func ServerDeactivateAccount(ctx context.Context, c xrpc.Client, input *ServerDeactivateAccount_Input) error {
	if err := c.Do(ctx, xrpc.Procedure, "application/json", "com.atproto.server.deactivateAccount", nil, input, nil); err != nil {
		return err
	}

	return nil
}
