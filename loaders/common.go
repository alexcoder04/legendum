package loaders

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"

	"github.com/alexcoder04/legendum/common"
)

func Hash(p common.Post) string {
	hasher := sha1.New()
	hasher.Write([]byte(fmt.Sprintf("%s:%s:%v", p.Title, p.Channel.Url, p.TimeCreated)))
	return base64.StdEncoding.EncodeToString(hasher.Sum(nil))
}
