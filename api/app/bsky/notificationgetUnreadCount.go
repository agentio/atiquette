// Code generated by cmd/lexgenlite (see Makefile's lexgen); DO NOT EDIT.

package bsky

// schema: app.bsky.notification.getUnreadCount

import (
	"context"

	"github.com/agentio/atiquette/pkg/xrpc"
)

// NotificationGetUnreadCount_Output is the output of a app.bsky.notification.getUnreadCount call.
type NotificationGetUnreadCount_Output struct {
	Count int64 `json:"count" cborgen:"count"`
}

// NotificationGetUnreadCount calls the XRPC method "app.bsky.notification.getUnreadCount".
func NotificationGetUnreadCount(ctx context.Context, c xrpc.Client, priority bool, seenAt string) (*NotificationGetUnreadCount_Output, error) {
	var out NotificationGetUnreadCount_Output

	params := map[string]interface{}{
		"priority": priority,
		"seenAt":   seenAt,
	}
	if err := c.Do(ctx, xrpc.Query, "", "app.bsky.notification.getUnreadCount", params, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
