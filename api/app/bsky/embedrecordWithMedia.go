// Code generated by cmd/lexgenlite (see Makefile's lexgen); DO NOT EDIT.

package bsky

// schema: app.bsky.embed.recordWithMedia

import (
	"encoding/json"
	"fmt"

	"github.com/agentio/atiquette/pkg/lexutil"
)

func init() {
	lexutil.RegisterType("app.bsky.embed.recordWithMedia#main", &EmbedRecordWithMedia{})
}

// EmbedRecordWithMedia is a "main" in the app.bsky.embed.recordWithMedia schema.
//
// RECORDTYPE: EmbedRecordWithMedia
type EmbedRecordWithMedia struct {
	LexiconTypeID string                      `json:"$type,const=app.bsky.embed.recordWithMedia" cborgen:"$type,const=app.bsky.embed.recordWithMedia"`
	Media         *EmbedRecordWithMedia_Media `json:"media" cborgen:"media"`
	Record        *EmbedRecord                `json:"record" cborgen:"record"`
}

type EmbedRecordWithMedia_Media struct {
	EmbedImages   *EmbedImages
	EmbedVideo    *EmbedVideo
	EmbedExternal *EmbedExternal
}

func (t *EmbedRecordWithMedia_Media) MarshalJSON() ([]byte, error) {
	if t.EmbedImages != nil {
		t.EmbedImages.LexiconTypeID = "app.bsky.embed.images"
		return json.Marshal(t.EmbedImages)
	}
	if t.EmbedVideo != nil {
		t.EmbedVideo.LexiconTypeID = "app.bsky.embed.video"
		return json.Marshal(t.EmbedVideo)
	}
	if t.EmbedExternal != nil {
		t.EmbedExternal.LexiconTypeID = "app.bsky.embed.external"
		return json.Marshal(t.EmbedExternal)
	}
	return nil, fmt.Errorf("cannot marshal empty enum")
}
func (t *EmbedRecordWithMedia_Media) UnmarshalJSON(b []byte) error {
	typ, err := lexutil.TypeExtract(b)
	if err != nil {
		return err
	}

	switch typ {
	case "app.bsky.embed.images":
		t.EmbedImages = new(EmbedImages)
		return json.Unmarshal(b, t.EmbedImages)
	case "app.bsky.embed.video":
		t.EmbedVideo = new(EmbedVideo)
		return json.Unmarshal(b, t.EmbedVideo)
	case "app.bsky.embed.external":
		t.EmbedExternal = new(EmbedExternal)
		return json.Unmarshal(b, t.EmbedExternal)

	default:
		return nil
	}
}

// EmbedRecordWithMedia_View is a "view" in the app.bsky.embed.recordWithMedia schema.
//
// RECORDTYPE: EmbedRecordWithMedia_View
type EmbedRecordWithMedia_View struct {
	LexiconTypeID string                           `json:"$type,const=app.bsky.embed.recordWithMedia#view" cborgen:"$type,const=app.bsky.embed.recordWithMedia#view"`
	Media         *EmbedRecordWithMedia_View_Media `json:"media" cborgen:"media"`
	Record        *EmbedRecord_View                `json:"record" cborgen:"record"`
}

type EmbedRecordWithMedia_View_Media struct {
	EmbedImages_View   *EmbedImages_View
	EmbedVideo_View    *EmbedVideo_View
	EmbedExternal_View *EmbedExternal_View
}

func (t *EmbedRecordWithMedia_View_Media) MarshalJSON() ([]byte, error) {
	if t.EmbedImages_View != nil {
		t.EmbedImages_View.LexiconTypeID = "app.bsky.embed.images#view"
		return json.Marshal(t.EmbedImages_View)
	}
	if t.EmbedVideo_View != nil {
		t.EmbedVideo_View.LexiconTypeID = "app.bsky.embed.video#view"
		return json.Marshal(t.EmbedVideo_View)
	}
	if t.EmbedExternal_View != nil {
		t.EmbedExternal_View.LexiconTypeID = "app.bsky.embed.external#view"
		return json.Marshal(t.EmbedExternal_View)
	}
	return nil, fmt.Errorf("cannot marshal empty enum")
}
func (t *EmbedRecordWithMedia_View_Media) UnmarshalJSON(b []byte) error {
	typ, err := lexutil.TypeExtract(b)
	if err != nil {
		return err
	}

	switch typ {
	case "app.bsky.embed.images#view":
		t.EmbedImages_View = new(EmbedImages_View)
		return json.Unmarshal(b, t.EmbedImages_View)
	case "app.bsky.embed.video#view":
		t.EmbedVideo_View = new(EmbedVideo_View)
		return json.Unmarshal(b, t.EmbedVideo_View)
	case "app.bsky.embed.external#view":
		t.EmbedExternal_View = new(EmbedExternal_View)
		return json.Unmarshal(b, t.EmbedExternal_View)

	default:
		return nil
	}
}
