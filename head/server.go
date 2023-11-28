package head

import (
	"google.golang.org/grpc"
	"goremote/proto"
)

type serverAPI struct {
	proto.UnimplementedRemoteControlServer
}

func Register(gRPC *grpc.Server) {
	proto.RegisterRemoteControlServer(gRPC, &serverAPI{})
}

func (s *serverAPI) GetScreenshot(req *proto.ScreenshotRequest, rc proto.RemoteControl_GetScreenshotServer) error {
	panic("implement me")
}
