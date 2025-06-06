// Code generated by cmd/lexgenlite (see Makefile's lexgen); DO NOT EDIT.

package bsky

// schema: app.bsky.feed.getActorFeeds

import (
	"context"

	"github.com/agentio/atiquette/pkg/xrpc"
)

// FeedGetActorFeeds_Output is the output of a app.bsky.feed.getActorFeeds call.
type FeedGetActorFeeds_Output struct {
	Cursor *string                   `json:"cursor,omitempty" cborgen:"cursor,omitempty"`
	Feeds  []*FeedDefs_GeneratorView `json:"feeds" cborgen:"feeds"`
}

// FeedGetActorFeeds calls the XRPC method "app.bsky.feed.getActorFeeds".
func FeedGetActorFeeds(ctx context.Context, c xrpc.Client, actor string, cursor string, limit int64) (*FeedGetActorFeeds_Output, error) {
	var out FeedGetActorFeeds_Output

	params := map[string]interface{}{
		"actor":  actor,
		"cursor": cursor,
		"limit":  limit,
	}
	if err := c.Do(ctx, xrpc.Query, "", "app.bsky.feed.getActorFeeds", params, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
