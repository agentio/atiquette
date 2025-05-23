// Code generated by cmd/lexgenlite (see Makefile's lexgen); DO NOT EDIT.

package bsky

// schema: app.bsky.graph.unmuteActorList

import (
	"context"

	"github.com/agentio/atiquette/pkg/xrpc"
)

// GraphUnmuteActorList_Input is the input argument to a app.bsky.graph.unmuteActorList call.
type GraphUnmuteActorList_Input struct {
	List string `json:"list" cborgen:"list"`
}

// GraphUnmuteActorList calls the XRPC method "app.bsky.graph.unmuteActorList".
func GraphUnmuteActorList(ctx context.Context, c xrpc.Client, input *GraphUnmuteActorList_Input) error {
	if err := c.Do(ctx, xrpc.Procedure, "application/json", "app.bsky.graph.unmuteActorList", nil, input, nil); err != nil {
		return err
	}

	return nil
}
