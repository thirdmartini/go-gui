package framebuffer

import (
	"fmt"
	"image"

	_ "github.com/nfnt/resize"
	"golang.org/x/sys/unix"

	"github.com/thirdmartini/gogui/pkg/drivers/display"
)

type FrameBuffer struct {
	Width    int
	Height   int
	buffer   []byte
	fd       int
	rotation int
}

func (f *FrameBuffer) WithRotation(rotation int) *FrameBuffer {
	f.rotation = rotation
	return f
}

func (f *FrameBuffer) Draw(im *image.RGBA) error {
	rgb := im.Pix
	switch f.rotation {
	case display.RotationNone, display.Rotation180:
		for pos := 0; pos < len(f.buffer); pos += 4 {
			f.buffer[pos] = rgb[pos+2]
			f.buffer[pos+1] = rgb[pos+1]
			f.buffer[pos+2] = rgb[pos]
			f.buffer[pos+3] = rgb[pos+3]
		}
	case display.Rotation90, display.Rotation270:
		// FIXME only do 90
		b := im.Bounds()
		w := b.Dx()
		h := b.Dy()

		stride := f.Width * 4 // we assume 32bit color, might want to muck with that
		for y := 0; y < h; y++ {
			for x := 0; x < w; x++ {
				srcIdx := (y*im.Stride + x*4)

				// Destination: (y, w-1-x)
				dstX := y
				dstY := w - 1 - x
				dstIdx := (dstY*stride + dstX*4)

				f.buffer[dstIdx] = rgb[srcIdx+2]
				f.buffer[dstIdx+1] = rgb[srcIdx+1]
				f.buffer[dstIdx+2] = rgb[srcIdx]
				f.buffer[dstIdx+3] = rgb[srcIdx+3]
			}
		}
	}
	return nil
}

func (f *FrameBuffer) Size() image.Point {
	switch f.rotation {
	case display.Rotation90, display.Rotation270:
		return image.Point{
			X: f.Height,
			Y: f.Width,
		}
	}

	return image.Point{
		X: f.Width,
		Y: f.Height,
	}
}

func (f *FrameBuffer) Close() error {
	if len(f.buffer) != 0 {
		unix.Munmap(f.buffer)
		f.buffer = nil
		return unix.Close(f.fd)
	}
	return nil
}

func ioctl(fd, op, arg uintptr) error {
	_, _, ep := unix.Syscall(unix.SYS_IOCTL, fd, op, arg)
	if ep != 0 {
		return ep
	}
	return nil
}

func Open(device string, width, height int) (*FrameBuffer, error) {
	fd, err := unix.Open(device, unix.O_RDWR|unix.O_NONBLOCK, 0666)
	if err != nil {
		panic(err)
		return nil, err
	}
	fmt.Printf("Frame buffer opened\n")
	length := width * height * 4

	buffer, err := unix.Mmap(fd, int64(0), length, unix.PROT_READ|unix.PROT_WRITE, unix.MAP_SHARED)
	if err != nil {
		unix.Close(fd)
		panic(err)
		return nil, err
	}

	fmt.Printf("Frame buffer mmap\n")
	ttyfd, err := unix.Open("/dev/tty0", unix.O_RDWR, 0666)
	if err != nil {
		unix.Munmap(buffer)
		unix.Close(fd)
		panic(err)
		return nil, err
	}
	defer unix.Close(ttyfd)

	fmt.Printf("Frame setfd\n")
	err = ioctl(uintptr(ttyfd), 0x4B3A, uintptr(0x1))
	if err != nil {
		unix.Munmap(buffer)
		unix.Close(fd)
		panic(err)
		return nil, err
	}

	r := &FrameBuffer{
		Width:  width,
		Height: height,
		fd:     fd,
		buffer: buffer,
	}

	return r, nil
}
