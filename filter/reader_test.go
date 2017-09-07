package filter

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	doc := []byte("<xml>\n\t<doc>can\tdog\tfrog\n</doc>\n\t</xml>")

	filter := func(r rune) rune {
		switch r {
		case '\t', '\n':
			return -1
		default:
			return r
		}
	}

	unfiltered, _ := ioutil.ReadAll(bytes.NewReader(doc))
	assert.Equal(t, string(doc), string(unfiltered))

	filtered, _ := ioutil.ReadAll(NewReader(bytes.NewReader(doc), filter))
	assert.Equal(t, "<xml><doc>candogfrog</doc></xml>", string(filtered))
}
