package emoji

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
)

type EmojiList map[string]string

func update() {
	const url = `https://api.github.com/emojis`
	c := http.Client{}
	r, _ := http.NewRequest("GET", url, nil)
	r.Header.Add("Accept", "application/vnd.github.v3+json")
	response, _ := c.Do(r)
	content, _ := io.ReadAll(response.Body)

	e := EmojiList{}
	json.Unmarshal(content, &e)

	const findEmojiCodeRegex = `(unicode\/)(.*?)((\.png)|(-))`
	rgx, err := regexp.Compile(findEmojiCodeRegex)

	if err != nil {
		panic("regex fail " + err.Error())
	}

	l := emojiIndexed{}
	for name, url := range e {
		strNum := rgx.Find([]byte(url))
		strNum = regexp.MustCompile(`(unicode\/)|(\.png)|(-)`).ReplaceAll(strNum, []byte{})

		o := Emoji{IconURL: url}
		fmt.Sscanf(string(strNum), "%x", &o.Code)
		l[name] = o
	}

	fmt.Printf("%#v", l)
}
