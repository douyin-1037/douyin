package main

import (
	"context"
	messageproto "douyin/code_gen/kitex_gen/messageproto"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct{}

// CreateMessage implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) CreateMessage(ctx context.Context, req *messageproto.CreateMessageReq) (resp *messageproto.CreateMessageResp, err error) {
	// TODO: Your code here...
	return
}

// GetMessageList implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) GetMessageList(ctx context.Context, req *messageproto.GetMessageListReq) (resp *messageproto.GetMessageListResp, err error) {
	// TODO: Your code here...
	return
}
