// Code generated by cmd/lexgenlite (see Makefile's lexgen); DO NOT EDIT.

package bsky

// schema: app.bsky.graph.getList

import (
	"context"

	"github.com/agentio/atiquette/pkg/xrpc"
)

// GraphGetList_Output is the output of a app.bsky.graph.getList call.
type GraphGetList_Output struct {
	Cursor *string                   `json:"cursor,omitempty" cborgen:"cursor,omitempty"`
	Items  []*GraphDefs_ListItemView `json:"items" cborgen:"items"`
	List   *GraphDefs_ListView       `json:"list" cborgen:"list"`
}

// GraphGetList calls the XRPC method "app.bsky.graph.getList".
//
// list: Reference (AT-URI) of the list record to hydrate.
func GraphGetList(ctx context.Context, c xrpc.Client, cursor string, limit int64, list string) (*GraphGetList_Output, error) {
	var out GraphGetList_Output

	params := map[string]interface{}{
		"cursor": cursor,
		"limit":  limit,
		"list":   list,
	}
	if err := c.Do(ctx, xrpc.Query, "", "app.bsky.graph.getList", params, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
