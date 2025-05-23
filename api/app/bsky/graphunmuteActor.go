// Code generated by cmd/lexgenlite (see Makefile's lexgen); DO NOT EDIT.

package bsky

// schema: app.bsky.graph.unmuteActor

import (
	"context"

	"github.com/agentio/atiquette/pkg/xrpc"
)

// GraphUnmuteActor_Input is the input argument to a app.bsky.graph.unmuteActor call.
type GraphUnmuteActor_Input struct {
	Actor string `json:"actor" cborgen:"actor"`
}

// GraphUnmuteActor calls the XRPC method "app.bsky.graph.unmuteActor".
func GraphUnmuteActor(ctx context.Context, c xrpc.Client, input *GraphUnmuteActor_Input) error {
	if err := c.Do(ctx, xrpc.Procedure, "application/json", "app.bsky.graph.unmuteActor", nil, input, nil); err != nil {
		return err
	}

	return nil
}
