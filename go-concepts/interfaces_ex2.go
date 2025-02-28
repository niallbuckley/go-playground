package main

import (
    "net"
    "fmt"
)

type Status interface {
    String() string
    Error() error
}

type CoreCtxInterface interface {
    Status() []Status
    CLIPrint()
}

type ArchStatusInterface interface {
	Error()					error
	String()				string
}

type CoreCtx struct {
    ips     []net.IP

    status  []ArchStatusInterface
}

func (ctx *CoreCtx) CLIPrint() {
    fmt.Println("Called CoreCtx CLIPrint")
}

type playLiveVideoCtx struct {
    CoreCtx

    primaryArchiver     string

    stats   []string
}

func (ctx *playLiveVideoCtx) Status() []Status {
    fmt.Println("Called playLiveVideoCtx Status")
    return nil
}

type playLivePreviewCtx struct {
    CoreCtx
}

func (ctx *playLivePreviewCtx) Status() []Status {
    fmt.Println("Called playLivePreviewCtx Status")
    return nil
}

func (ctx *playLivePreviewCtx) CLIPrint() {
    fmt.Println("Called playLivePreviewCtx CLIPrint")
}

func MonitorStatusBase(ctx CoreCtxInterface) {
    ctx.CLIPrint()
    ctx.Status()
}

func main() {
    vidctx := playLiveVideoCtx{}
    prevctx := playLivePreviewCtx{}

    fmt.Println("Video Ctx")
    MonitorStatusBase(&vidctx)

    fmt.Println("Prev Ctx")
    MonitorStatusBase(&prevctx)
}

