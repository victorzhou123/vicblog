package entity

import "io"

const pictureSizeLimit = 1 << 22

type Picture struct {
	Name PictureName
	Data io.Reader
	Size int64
}

func (p *Picture) OverSizeLimited() bool {
	return p.Size > pictureSizeLimit
}
