package ccontext

import (
	"context"
	"github.com/OpenIMSDK/Open-IM-Server/pkg/common/constant"
	"github.com/OpenIMSDK/Open-IM-Server/pkg/errs"
)

const (
	Platform             = "platform"
	ApiAddr              = "apiAddr"
	WsAddr               = "wsAddr"
	DataDir              = "dataDir"
	LogLevel             = "logLevel"
	ObjectStorage        = "objectStorage"
	EncryptionKey        = "encryptionKey"
	IsCompression        = "isCompression"
	IsExternalExtensions = "isExternalExtensions"
)

var mapper = []string{Platform, ApiAddr, WsAddr, DataDir, LogLevel,
	ObjectStorage, EncryptionKey, IsCompression, IsExternalExtensions}

func WithOpUserIDContext(ctx context.Context, opUserID string) context.Context {
	return context.WithValue(ctx, constant.OpUserID, opUserID)
}
func WithOpUserPlatformContext(ctx context.Context, platform string) context.Context {
	return context.WithValue(ctx, constant.OpUserPlatform, platform)
}
func WithTriggerIDContext(ctx context.Context, triggerID string) context.Context {
	return context.WithValue(ctx, constant.TriggerID, triggerID)
}
func NewCtx(operationID string) context.Context {
	c := context.Background()
	ctx := context.WithValue(c, constant.OperationID, operationID)
	return SetOperationID(ctx, operationID)
}

func SetOperationID(ctx context.Context, operationID string) context.Context {
	return context.WithValue(ctx, constant.OperationID, operationID)
}

func SetOpUserID(ctx context.Context, opUserID string) context.Context {
	return context.WithValue(ctx, constant.OpUserID, opUserID)
}

func SetConnID(ctx context.Context, connID string) context.Context {
	return context.WithValue(ctx, constant.ConnID, connID)
}

func GetOperationID(ctx context.Context) string {
	if ctx.Value(constant.OperationID) != nil {
		s, ok := ctx.Value(constant.OperationID).(string)
		if ok {
			return s
		}
	}
	return ""
}
func GetOpUserID(ctx context.Context) string {
	if ctx.Value(constant.OpUserID) != "" {
		s, ok := ctx.Value(constant.OpUserID).(string)
		if ok {
			return s
		}
	}
	return ""
}
func GetConnID(ctx context.Context) string {
	if ctx.Value(constant.ConnID) != "" {
		s, ok := ctx.Value(constant.ConnID).(string)
		if ok {
			return s
		}
	}
	return ""
}

func GetTriggerID(ctx context.Context) string {
	if ctx.Value(constant.TriggerID) != "" {
		s, ok := ctx.Value(constant.TriggerID).(string)
		if ok {
			return s
		}
	}
	return ""
}
func GetPlatform(ctx context.Context) int32 {
	if ctx.Value(Platform) != "" {
		s, ok := ctx.Value(Platform).(int32)
		if ok {
			return s
		}
	}
	return 0
}
func GetOpUserPlatform(ctx context.Context) string {
	if ctx.Value(constant.OpUserPlatform) != "" {
		s, ok := ctx.Value(constant.OpUserPlatform).(string)
		if ok {
			return s
		}
	}
	return ""
}
func GetDataDir(ctx context.Context) string {
	if ctx.Value(DataDir) != "" {
		s, ok := ctx.Value(DataDir).(string)
		if ok {
			return s
		}
	}
	return ""
}
func GetRemoteAddr(ctx context.Context) string {
	if ctx.Value(constant.RemoteAddr) != "" {
		s, ok := ctx.Value(constant.RemoteAddr).(string)
		if ok {
			return s
		}
	}
	return ""
}

func GetMustCtxInfo(ctx context.Context) (operationID, opUserID, platform, connID string, err error) {
	operationID, ok := ctx.Value(constant.OperationID).(string)
	if !ok {
		err = errs.ErrArgs.Wrap("ctx missing operationID")
		return
	}
	opUserID, ok1 := ctx.Value(constant.OpUserID).(string)
	if !ok1 {
		err = errs.ErrArgs.Wrap("ctx missing opUserID")
		return
	}
	platform, ok2 := ctx.Value(constant.OpUserPlatform).(string)
	if !ok2 {
		err = errs.ErrArgs.Wrap("ctx missing platform")
		return
	}
	connID, _ = ctx.Value(constant.ConnID).(string)
	return

}
func WithMustInfoCtx(values []interface{}) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())
	for i, v := range values {
		ctx = context.WithValue(ctx, mapper[i], v)

	}
	return ctx, cancel

}
