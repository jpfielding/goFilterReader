package filter

import (
	"bytes"
	"encoding/xml"
	"io"
	"testing"

	testutils "github.com/jpfielding/goTest/testutils"
)

func TestBadXml(t *testing.T) {
	doc := []byte("<?xml version=\"1.0\"?>\x0b<doc/>")

	unfiltered := xml.NewDecoder(bytes.NewReader(doc))
	var err error

	for err == nil {
		_, err = unfiltered.Token()
	}
	testutils.NotOk(t, err)
}

func TestGoodXml(t *testing.T) {
	doc := []byte("<?xml version=\"1.0\"?>\x0b<doc/>")

	filtered := xml.NewDecoder(NewReader(bytes.NewReader(doc), XML10Filter))
	var err error

	for err == nil {
		_, err = filtered.Token()
	}
	switch err {
	case io.EOF:
	default:
		testutils.Ok(t, err)

	}
}
