// Code generated by cmd/lexgenlite (see Makefile's lexgen); DO NOT EDIT.

package bsky

// schema: app.bsky.feed.getAuthorFeed

import (
	"context"

	"github.com/agentio/atiquette/pkg/xrpc"
)

// FeedGetAuthorFeed_Output is the output of a app.bsky.feed.getAuthorFeed call.
type FeedGetAuthorFeed_Output struct {
	Cursor *string                  `json:"cursor,omitempty" cborgen:"cursor,omitempty"`
	Feed   []*FeedDefs_FeedViewPost `json:"feed" cborgen:"feed"`
}

// FeedGetAuthorFeed calls the XRPC method "app.bsky.feed.getAuthorFeed".
//
// filter: Combinations of post/repost types to include in response.
func FeedGetAuthorFeed(ctx context.Context, c xrpc.Client, actor string, cursor string, filter string, includePins bool, limit int64) (*FeedGetAuthorFeed_Output, error) {
	var out FeedGetAuthorFeed_Output

	params := map[string]interface{}{
		"actor":       actor,
		"cursor":      cursor,
		"filter":      filter,
		"includePins": includePins,
		"limit":       limit,
	}
	if err := c.Do(ctx, xrpc.Query, "", "app.bsky.feed.getAuthorFeed", params, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
