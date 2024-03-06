package emoji

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
)

//go:embed emojis.json
var emojisJSON []byte
var Emojis Indexed

func init() {
	_ = json.NewDecoder(bytes.NewBuffer(emojisJSON)).Decode(&Emojis)
}

type Emoji struct {
	IconURL  string
	Code     int
	Unicodes []rune
}

func (e Emoji) String() string {
	return fmt.Sprintf("%s\t%s", string(e.Unicodes), e.IconURL)
}

const BaseURL = "https://github.githubassets.com/images/icons/emoji/unicode"

type Indexed map[string]Emoji

func (i Indexed) String() (s string) {
	for name, code := range i {
		s = fmt.Sprintf("%s\n%-40s %s", s, name, code)
	}
	return fmt.Sprintln(s)
}
