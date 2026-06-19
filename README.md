# go-gui

Toolkit for making raw UI Stuff in golang for use with RPIs and touch displays.

# Building

You can build the examples by running:

```aiignore
roland@rpi go-gui % make
go build -o build/demo github.com/thirdmartini/gogui/cmd/demo
```

# Running

If you copy the demo binary someplace else make sure you copy the assets folder as well.

## VNC Mode

The demo comes with very basic VNC server, but its functional enough to run the demo and test the UI without running 
on an RPI.

![VNC Mode](assets/vnc.png)

```aiignore
roland@rpi go-gui % ./build/demo -vnc

2026/06/18 22:12:11 Listening on: vnc://localhost:9000
...
```
I've tested this with TigerVNC on macOS.  It does NOT work with the builtin VNC viewer on the mac as the server does 
not support authentication required by the MacOS Builting VNC viewer.  Use TigerVNC from brew (or add the missing vnc features.)

## Real Mode

Note that the demo app is configured for the [WaveShare 11.9" touch display](https://www.amazon.com/dp/B092LSDMP8) 
which has a resolution of 320x1480.  The demo is running it as a 1480x320 by setting rotation to 90degrees in the 
demo code. 

```aiignore
roland@rpi go-gui % sudo ./build/demo
 
Frame buffer opened
Frame buffer mmap
Frame setfd
FB Acquired
...
```

Short Video:

[![Watch the video](assets/youtube.png)](https://www.youtube.com/watch?v=aj_DWQwIO-I)


# Annoying Bits

## Waveshare 11.9" Touch Display 320x1480

To get this to work on an RPI4:  ( Note: RPI5 is different )
```/boot/firmware/config.txt

# DISABLE DT Overlay
# dtoverlay=vc4-kms-v3d

# Add the following lines to the end of /boot/firmware/config.txt
[all]
max_framebuffer_height=1480
hdmi_group=2
hdmi_mode=87
hdmi_timings=320 0 80 16 32 1480 0 16 4 12 0 0 0 60 0 42000000 3
framebuffer_depth=32
```

```/boot/firmware/cmdline.txt
console=serial0,115200 console=tty1 root=PARTUUID=87e0f569-02 rootfstype=ext4 fsck.repair=yes rootwait video=HDMI-A-1:320x1480-32@60
```