// Code generated by cmd/lexgenlite (see Makefile's lexgen); DO NOT EDIT.

package atproto

// schema: com.atproto.temp.fetchLabels

import (
	"context"

	"github.com/agentio/atiquette/pkg/xrpc"
)

// TempFetchLabels_Output is the output of a com.atproto.temp.fetchLabels call.
type TempFetchLabels_Output struct {
	Labels []*LabelDefs_Label `json:"labels" cborgen:"labels"`
}

// TempFetchLabels calls the XRPC method "com.atproto.temp.fetchLabels".
func TempFetchLabels(ctx context.Context, c xrpc.Client, limit int64, since int64) (*TempFetchLabels_Output, error) {
	var out TempFetchLabels_Output

	params := map[string]interface{}{
		"limit": limit,
		"since": since,
	}
	if err := c.Do(ctx, xrpc.Query, "", "com.atproto.temp.fetchLabels", params, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
