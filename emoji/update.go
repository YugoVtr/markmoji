package emoji

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

type List map[string]string

func Update(out io.Writer) {
	const url = `https://api.github.com/emojis`
	c := http.Client{}
	r, _ := http.NewRequestWithContext(context.Background(), "GET", url, http.NoBody)
	r.Header.Add("Accept", "application/vnd.github.v3+json")
	response, err := c.Do(r)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)

	e := List{}
	_ = json.Unmarshal(content, &e)

	const findEmojiCodeRegex = `(unicode/)(.*?)((\.png))`
	rgx := regexp.MustCompile(findEmojiCodeRegex)

	l := emojiIndexed{}
	for name, url := range e {
		strNum := rgx.Find([]byte(url))
		strNum = regexp.MustCompile(`(unicode/)|(\.png)`).ReplaceAll(strNum, []byte{})

		o := Emoji{IconURL: url}
		for _, unicode := range strings.Split(string(strNum), "-") {
			var r rune

			fmt.Sscanf(unicode, "%x", &r)
			o.Unicodes = append(o.Unicodes, r)
		}

		l[name] = o
	}

	fmt.Fprintf(out, "%s", l)
}
