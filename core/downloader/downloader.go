package downloader

import (
	"spider/utils/context"
	"spider/utils/page"
)

// The Downloader interface.
// You can implement the interface by implement function Download.
// Function Download need to return Page instance pointer that has request result downloaded from Request.
type Downloader interface {
	Download(req *context.Request) *page.Page
}

