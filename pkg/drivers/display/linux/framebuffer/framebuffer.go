//go:build !linux

package framebuffer

import (
	"fmt"
	"image"
)

type FrameBuffer struct{}

func (d *FrameBuffer) Draw(im *image.RGBA) error {
	return nil
}

func (d *FrameBuffer) WithRotation(rotation int) *FrameBuffer {
	return d
}

func (d *FrameBuffer) Size() image.Point {
	return image.Point{}
}

func (d *FrameBuffer) Close() error {
	return nil
}

func Open(device string, width, height int) (*FrameBuffer, error) {

	return nil, fmt.Errorf("not supported on this platform")
}
