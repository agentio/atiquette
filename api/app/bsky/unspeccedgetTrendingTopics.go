// Code generated by cmd/lexgenlite (see Makefile's lexgen); DO NOT EDIT.

package bsky

// schema: app.bsky.unspecced.getTrendingTopics

import (
	"context"

	"github.com/agentio/atiquette/pkg/xrpc"
)

// UnspeccedGetTrendingTopics_Output is the output of a app.bsky.unspecced.getTrendingTopics call.
type UnspeccedGetTrendingTopics_Output struct {
	Suggested []*UnspeccedDefs_TrendingTopic `json:"suggested" cborgen:"suggested"`
	Topics    []*UnspeccedDefs_TrendingTopic `json:"topics" cborgen:"topics"`
}

// UnspeccedGetTrendingTopics calls the XRPC method "app.bsky.unspecced.getTrendingTopics".
//
// viewer: DID of the account making the request (not included for public/unauthenticated queries). Used to boost followed accounts in ranking.
func UnspeccedGetTrendingTopics(ctx context.Context, c xrpc.Client, limit int64, viewer string) (*UnspeccedGetTrendingTopics_Output, error) {
	var out UnspeccedGetTrendingTopics_Output

	params := map[string]interface{}{
		"limit":  limit,
		"viewer": viewer,
	}
	if err := c.Do(ctx, xrpc.Query, "", "app.bsky.unspecced.getTrendingTopics", params, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
