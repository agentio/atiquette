// Code generated by cmd/lexgenlite (see Makefile's lexgen); DO NOT EDIT.

package chat

// schema: chat.bsky.actor.exportAccountData

import (
	"bytes"
	"context"

	"github.com/agentio/atiquette/pkg/xrpc"
)

// ActorExportAccountData calls the XRPC method "chat.bsky.actor.exportAccountData".
func ActorExportAccountData(ctx context.Context, c xrpc.Client) ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := c.Do(ctx, xrpc.Query, "", "chat.bsky.actor.exportAccountData", nil, nil, buf); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
