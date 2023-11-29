package head

import (
	"fmt"
	"github.com/kbinani/screenshot"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"goremote/proto"
	"image/png"
	"io"
	"os"
)

type serverAPI struct {
	proto.UnimplementedRemoteControlServer
}

func Register(gRPC *grpc.Server) {
	proto.RegisterRemoteControlServer(gRPC, &serverAPI{})
}

func (s *serverAPI) GetScreenshot(req *proto.ScreenshotRequest, stream proto.RemoteControl_GetScreenshotServer) error {
	buffer := make([]byte, 64*1024)

	scr := TakeScreenshot()

	for {
		n, err := scr.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		resp := &proto.ScreenshotResponse{
			Chunk: buffer[:n],
			Error: "",
		}
		if err := stream.Send(resp); err != nil {
			return status.Error(codes.Internal, err.Error())
		}
	}
	return nil
}

func TakeScreenshot() *os.File {
	n := screenshot.NumActiveDisplays()

	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)

		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			panic(err)
		}
		fileName := fmt.Sprintf("%d_%dx%d.png", i, bounds.Dx(), bounds.Dy())
		scr, _ := os.Create(fileName)
		defer scr.Close()
		err = png.Encode(scr, img)
		if err != nil {
			return nil
		}
		return scr

	}
	return nil
}
