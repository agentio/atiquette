// Code generated by cmd/lexgenlite (see Makefile's lexgen); DO NOT EDIT.

package bsky

// schema: app.bsky.graph.getListMutes

import (
	"context"

	"github.com/agentio/atiquette/pkg/xrpc"
)

// GraphGetListMutes_Output is the output of a app.bsky.graph.getListMutes call.
type GraphGetListMutes_Output struct {
	Cursor *string               `json:"cursor,omitempty" cborgen:"cursor,omitempty"`
	Lists  []*GraphDefs_ListView `json:"lists" cborgen:"lists"`
}

// GraphGetListMutes calls the XRPC method "app.bsky.graph.getListMutes".
func GraphGetListMutes(ctx context.Context, c xrpc.Client, cursor string, limit int64) (*GraphGetListMutes_Output, error) {
	var out GraphGetListMutes_Output

	params := map[string]interface{}{
		"cursor": cursor,
		"limit":  limit,
	}
	if err := c.Do(ctx, xrpc.Query, "", "app.bsky.graph.getListMutes", params, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
