package errs

// 框架层接口调用的返回码定义
type TrpcRetCode int32

const (
	// 调用成功
	TrpcRetCode_TRPC_INVOKE_SUCCESS TrpcRetCode = 0
	// 协议错误码
	// 服务端解码错误
	TrpcRetCode_TRPC_SERVER_DECODE_ERR TrpcRetCode = 1
	// 服务端编码错误
	TrpcRetCode_TRPC_SERVER_ENCODE_ERR TrpcRetCode = 2
	// service或者func路由错误码
	// 服务端没有调用相应的service实现
	TrpcRetCode_TRPC_SERVER_NOSERVICE_ERR TrpcRetCode = 11
	// 服务端没有调用相应的接口实现
	TrpcRetCode_TRPC_SERVER_NOFUNC_ERR TrpcRetCode = 12
	// 超时/过载/限流错误码
	// 请求在服务端超时
	TrpcRetCode_TRPC_SERVER_TIMEOUT_ERR TrpcRetCode = 21
	// 请求在服务端被过载保护而丢弃请求
	// 主要用在框架内部实现的过载保护插件上
	TrpcRetCode_TRPC_SERVER_OVERLOAD_ERR TrpcRetCode = 22
	// 请求在服务端被限流
	// 主要用在外部服务治理系统的插件或者业务自定义的限流插件上，比如: 北极星限流
	TrpcRetCode_TRPC_SERVER_LIMITED_ERR TrpcRetCode = 23
	// 请求在服务端因全链路超时时间而超时
	TrpcRetCode_TRPC_SERVER_FULL_LINK_TIMEOUT_ERR TrpcRetCode = 24
	// 服务端系统错误
	TrpcRetCode_TRPC_SERVER_SYSTEM_ERR TrpcRetCode = 31
	// 服务端鉴权失败错误
	TrpcRetCode_TRPC_SERVER_AUTH_ERR TrpcRetCode = 41
	// 服务端请求参数自动校验失败错误
	TrpcRetCode_TRPC_SERVER_VALIDATE_ERR TrpcRetCode = 51
	// 超时错误码
	// 请求在客户端调用超时
	TrpcRetCode_TRPC_CLIENT_INVOKE_TIMEOUT_ERR TrpcRetCode = 101
	// 请求在客户端因全链路超时时间而超时
	TrpcRetCode_TRPC_CLIENT_FULL_LINK_TIMEOUT_ERR TrpcRetCode = 102
	// 网络相关错误码
	// 客户端连接错误
	TrpcRetCode_TRPC_CLIENT_CONNECT_ERR TrpcRetCode = 111
	// 协议相关错误码
	// 客户端编码错误
	TrpcRetCode_TRPC_CLIENT_ENCODE_ERR TrpcRetCode = 121
	// 客户端解码错误
	TrpcRetCode_TRPC_CLIENT_DECODE_ERR TrpcRetCode = 122
	// 过载保护/限流相关错误码
	// 请求在客户端被限流
	// 主要用在外部服务治理系统的插件或者业务自定义的限流插件上，比如: 北极星限流
	TrpcRetCode_TRPC_CLIENT_LIMITED_ERR TrpcRetCode = 123
	// 请求在客户端被过载保护而丢弃请求
	// 主要用在框架内部实现的过载保护插件上
	TrpcRetCode_TRPC_CLIENT_OVERLOAD_ERR TrpcRetCode = 124
	// 路由相关错误码
	// 客户端选ip路由错误
	TrpcRetCode_TRPC_CLIENT_ROUTER_ERR TrpcRetCode = 131
	// 客户端网络错误
	TrpcRetCode_TRPC_CLIENT_NETWORK_ERR TrpcRetCode = 141
	// 客户端响应参数自动校验失败错误
	TrpcRetCode_TRPC_CLIENT_VALIDATE_ERR TrpcRetCode = 151
	// 上游主动断开连接，提前取消请求错误
	TrpcRetCode_TRPC_CLIENT_CANCELED_ERR TrpcRetCode = 161
	// 客户端读取 Frame 错误
	TrpcRetCode_TRPC_CLIENT_READ_FRAME_ERR TrpcRetCode = 171
	// 服务端流式网络错误, 详细错误码需要在实现过程中再梳理
	TrpcRetCode_TRPC_STREAM_SERVER_NETWORK_ERR TrpcRetCode = 201
	// 服务端流式传输错误, 详细错误码需要在实现过程中再梳理
	// 比如：流消息过大等
	TrpcRetCode_TRPC_STREAM_SERVER_MSG_EXCEED_LIMIT_ERR TrpcRetCode = 211
	// 服务端流式编码错误
	TrpcRetCode_TRPC_STREAM_SERVER_ENCODE_ERR TrpcRetCode = 221
	// 客户端流式编解码错误
	TrpcRetCode_TRPC_STREAM_SERVER_DECODE_ERR TrpcRetCode = 222
	// 服务端流式写错误, 详细错误码需要在实现过程中再梳理
	TrpcRetCode_TRPC_STREAM_SERVER_WRITE_END          TrpcRetCode = 231
	TrpcRetCode_TRPC_STREAM_SERVER_WRITE_OVERFLOW_ERR TrpcRetCode = 232
	TrpcRetCode_TRPC_STREAM_SERVER_WRITE_CLOSE_ERR    TrpcRetCode = 233
	TrpcRetCode_TRPC_STREAM_SERVER_WRITE_TIMEOUT_ERR  TrpcRetCode = 234
	// 服务端流式读错误, 详细错误码需要在实现过程中再梳理
	TrpcRetCode_TRPC_STREAM_SERVER_READ_END         TrpcRetCode = 251
	TrpcRetCode_TRPC_STREAM_SERVER_READ_CLOSE_ERR   TrpcRetCode = 252
	TrpcRetCode_TRPC_STREAM_SERVER_READ_EMPTY_ERR   TrpcRetCode = 253
	TrpcRetCode_TRPC_STREAM_SERVER_READ_TIMEOUT_ERR TrpcRetCode = 254
	// 客户端流式网络错误, 详细错误码需要在实现过程中再梳理
	TrpcRetCode_TRPC_STREAM_CLIENT_NETWORK_ERR TrpcRetCode = 301
	// 客户端流式传输错误, 详细错误码需要在实现过程中再梳理
	// 比如：流消息过大等
	TrpcRetCode_TRPC_STREAM_CLIENT_MSG_EXCEED_LIMIT_ERR TrpcRetCode = 311
	// 客户端流式编码错误
	TrpcRetCode_TRPC_STREAM_CLIENT_ENCODE_ERR TrpcRetCode = 321
	// 客户端流式编解码错误
	TrpcRetCode_TRPC_STREAM_CLIENT_DECODE_ERR TrpcRetCode = 322
	// 客户端流式写错误, 详细错误码需要在实现过程中再梳理
	TrpcRetCode_TRPC_STREAM_CLIENT_WRITE_END          TrpcRetCode = 331
	TrpcRetCode_TRPC_STREAM_CLIENT_WRITE_OVERFLOW_ERR TrpcRetCode = 332
	TrpcRetCode_TRPC_STREAM_CLIENT_WRITE_CLOSE_ERR    TrpcRetCode = 333
	TrpcRetCode_TRPC_STREAM_CLIENT_WRITE_TIMEOUT_ERR  TrpcRetCode = 334
	// 客户端流式读错误, 详细错误码需要在实现过程中再梳理
	TrpcRetCode_TRPC_STREAM_CLIENT_READ_END         TrpcRetCode = 351
	TrpcRetCode_TRPC_STREAM_CLIENT_READ_CLOSE_ERR   TrpcRetCode = 352
	TrpcRetCode_TRPC_STREAM_CLIENT_READ_EMPTY_ERR   TrpcRetCode = 353
	TrpcRetCode_TRPC_STREAM_CLIENT_READ_TIMEOUT_ERR TrpcRetCode = 354
	// 未明确的错误
	TrpcRetCode_TRPC_INVOKE_UNKNOWN_ERR TrpcRetCode = 999
	// 未明确的错误
	TrpcRetCode_TRPC_STREAM_UNKNOWN_ERR TrpcRetCode = 1000
)
