// Code generated by cmd/lexgenlite (see Makefile's lexgen); DO NOT EDIT.

package bsky

// schema: app.bsky.feed.getPosts

import (
	"context"

	"github.com/agentio/atiquette/pkg/xrpc"
)

// FeedGetPosts_Output is the output of a app.bsky.feed.getPosts call.
type FeedGetPosts_Output struct {
	Posts []*FeedDefs_PostView `json:"posts" cborgen:"posts"`
}

// FeedGetPosts calls the XRPC method "app.bsky.feed.getPosts".
//
// uris: List of post AT-URIs to return hydrated views for.
func FeedGetPosts(ctx context.Context, c xrpc.Client, uris []string) (*FeedGetPosts_Output, error) {
	var out FeedGetPosts_Output

	params := map[string]interface{}{
		"uris": uris,
	}
	if err := c.Do(ctx, xrpc.Query, "", "app.bsky.feed.getPosts", params, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
