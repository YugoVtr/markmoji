package emoji_test

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yugovtr/markmoji/emoji"
)

func TestUpdate(t *testing.T) {
	buffer := bytes.Buffer{}
	indexed := emoji.Update(&buffer)
	require.NotEmpty(t, buffer.String())
	file, err := os.Create("emojis.json")
	require.NoError(t, err)
	t.Cleanup(func() { file.Close() })
	err = json.NewEncoder(file).Encode(indexed)
	require.NoError(t, err)
}
