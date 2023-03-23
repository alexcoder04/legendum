package processor

import "github.com/alexcoder04/legendum/common"

var ContentBuffer = make([]common.Post, 0)

func Load() []common.Post {
	contentChunk := ContentBuffer[:5]
	ContentBuffer = ContentBuffer[5:]
	return contentChunk
}
