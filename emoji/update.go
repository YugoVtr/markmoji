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

func Update(out io.Writer) Indexed {
	const url = `https://api.github.com/emojis`
	r, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, url, http.NoBody)
	r.Header.Add("Accept", "application/vnd.github.v3+json")
	response, err := http.DefaultClient.Do(r)
	if err != nil {
		panic(fmt.Errorf("failed to fetch emoji list: %w", err))
	}
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)

	e := List{}
	if err = json.Unmarshal(content, &e); err != nil {
		panic(fmt.Errorf("failed to unmarshal emoji list: %w", err))
	}

	const findEmojiCodeRegex = `(unicode/)(.*?)((\.png))`
	rgx := regexp.MustCompile(findEmojiCodeRegex)

	l := Indexed{}
	for name, url := range e {
		strNum := rgx.Find([]byte(url))
		strNum = regexp.MustCompile(`(unicode/)|(\.png)`).ReplaceAll(strNum, []byte{})
		o := Emoji{IconURL: url}
		for _, unicode := range strings.Split(string(strNum), "-") {
			var r rune
			fmt.Sscanf(unicode, "%x", &r)
			o.Unicodes = append(o.Unicodes, r)
		}
		if a := strings.Split(o.IconURL, "-"); len(a) > 1 {
			continue
		}
		o.Code = int(o.Unicodes[0])
		l[name] = o
	}
	fmt.Fprintf(out, "%s", l)
	return l
}
