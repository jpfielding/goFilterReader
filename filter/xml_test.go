package filter

import (
	"bytes"
	"encoding/xml"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBadXml(t *testing.T) {
	doc := []byte("<?xml version=\"1.0\"?>\x0b<doc/>")

	unfiltered := xml.NewDecoder(bytes.NewReader(doc))
	var err error

	for err == nil {
		_, err = unfiltered.Token()
	}
	assert.NotNil(t, err)
}

func TestGoodXml(t *testing.T) {
	doc := []byte("<?xml version=\"1.0\"?>\x0b<doc/>")

	filtered := xml.NewDecoder(NewReader(bytes.NewReader(doc), XML10Filter(DropChar)))
	var err error

	for err == nil {
		_, err = filtered.Token()
	}
	switch err {
	case io.EOF:
	default:
		assert.Nil(t, err)

	}
}
