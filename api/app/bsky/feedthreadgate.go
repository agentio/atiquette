// Code generated by cmd/lexgenlite (see Makefile's lexgen); DO NOT EDIT.

package bsky

// schema: app.bsky.feed.threadgate

import (
	"encoding/json"
	"fmt"

	"github.com/agentio/atiquette/pkg/lexutil"
)

func init() {
	lexutil.RegisterType("app.bsky.feed.threadgate", &FeedThreadgate{})
}

// RECORDTYPE: FeedThreadgate
type FeedThreadgate struct {
	LexiconTypeID string `json:"$type,const=app.bsky.feed.threadgate" cborgen:"$type,const=app.bsky.feed.threadgate"`
	// allow: List of rules defining who can reply to this post. If value is an empty array, no one can reply. If value is undefined, anyone can reply.
	Allow     []*FeedThreadgate_Allow_Elem `json:"allow,omitempty" cborgen:"allow,omitempty"`
	CreatedAt string                       `json:"createdAt" cborgen:"createdAt"`
	// hiddenReplies: List of hidden reply URIs.
	HiddenReplies []string `json:"hiddenReplies,omitempty" cborgen:"hiddenReplies,omitempty"`
	// post: Reference (AT-URI) to the post record.
	Post string `json:"post" cborgen:"post"`
}

type FeedThreadgate_Allow_Elem struct {
	FeedThreadgate_MentionRule   *FeedThreadgate_MentionRule
	FeedThreadgate_FollowerRule  *FeedThreadgate_FollowerRule
	FeedThreadgate_FollowingRule *FeedThreadgate_FollowingRule
	FeedThreadgate_ListRule      *FeedThreadgate_ListRule
}

func (t *FeedThreadgate_Allow_Elem) MarshalJSON() ([]byte, error) {
	if t.FeedThreadgate_MentionRule != nil {
		t.FeedThreadgate_MentionRule.LexiconTypeID = "app.bsky.feed.threadgate#mentionRule"
		return json.Marshal(t.FeedThreadgate_MentionRule)
	}
	if t.FeedThreadgate_FollowerRule != nil {
		t.FeedThreadgate_FollowerRule.LexiconTypeID = "app.bsky.feed.threadgate#followerRule"
		return json.Marshal(t.FeedThreadgate_FollowerRule)
	}
	if t.FeedThreadgate_FollowingRule != nil {
		t.FeedThreadgate_FollowingRule.LexiconTypeID = "app.bsky.feed.threadgate#followingRule"
		return json.Marshal(t.FeedThreadgate_FollowingRule)
	}
	if t.FeedThreadgate_ListRule != nil {
		t.FeedThreadgate_ListRule.LexiconTypeID = "app.bsky.feed.threadgate#listRule"
		return json.Marshal(t.FeedThreadgate_ListRule)
	}
	return nil, fmt.Errorf("cannot marshal empty enum")
}
func (t *FeedThreadgate_Allow_Elem) UnmarshalJSON(b []byte) error {
	typ, err := lexutil.TypeExtract(b)
	if err != nil {
		return err
	}

	switch typ {
	case "app.bsky.feed.threadgate#mentionRule":
		t.FeedThreadgate_MentionRule = new(FeedThreadgate_MentionRule)
		return json.Unmarshal(b, t.FeedThreadgate_MentionRule)
	case "app.bsky.feed.threadgate#followerRule":
		t.FeedThreadgate_FollowerRule = new(FeedThreadgate_FollowerRule)
		return json.Unmarshal(b, t.FeedThreadgate_FollowerRule)
	case "app.bsky.feed.threadgate#followingRule":
		t.FeedThreadgate_FollowingRule = new(FeedThreadgate_FollowingRule)
		return json.Unmarshal(b, t.FeedThreadgate_FollowingRule)
	case "app.bsky.feed.threadgate#listRule":
		t.FeedThreadgate_ListRule = new(FeedThreadgate_ListRule)
		return json.Unmarshal(b, t.FeedThreadgate_ListRule)

	default:
		return nil
	}
}

// FeedThreadgate_FollowerRule is a "followerRule" in the app.bsky.feed.threadgate schema.
//
// Allow replies from actors who follow you.
//
// RECORDTYPE: FeedThreadgate_FollowerRule
type FeedThreadgate_FollowerRule struct {
	LexiconTypeID string `json:"$type,const=app.bsky.feed.threadgate#followerRule" cborgen:"$type,const=app.bsky.feed.threadgate#followerRule"`
}

// FeedThreadgate_FollowingRule is a "followingRule" in the app.bsky.feed.threadgate schema.
//
// Allow replies from actors you follow.
//
// RECORDTYPE: FeedThreadgate_FollowingRule
type FeedThreadgate_FollowingRule struct {
	LexiconTypeID string `json:"$type,const=app.bsky.feed.threadgate#followingRule" cborgen:"$type,const=app.bsky.feed.threadgate#followingRule"`
}

// FeedThreadgate_ListRule is a "listRule" in the app.bsky.feed.threadgate schema.
//
// Allow replies from actors on a list.
//
// RECORDTYPE: FeedThreadgate_ListRule
type FeedThreadgate_ListRule struct {
	LexiconTypeID string `json:"$type,const=app.bsky.feed.threadgate#listRule" cborgen:"$type,const=app.bsky.feed.threadgate#listRule"`
	List          string `json:"list" cborgen:"list"`
}

// FeedThreadgate_MentionRule is a "mentionRule" in the app.bsky.feed.threadgate schema.
//
// Allow replies from actors mentioned in your post.
//
// RECORDTYPE: FeedThreadgate_MentionRule
type FeedThreadgate_MentionRule struct {
	LexiconTypeID string `json:"$type,const=app.bsky.feed.threadgate#mentionRule" cborgen:"$type,const=app.bsky.feed.threadgate#mentionRule"`
}
