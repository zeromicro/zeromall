package inner


/*
	内部 RPC Server: 对内提供 gRPC API

*/
type Server struct {
}

func NewServer() *Server {
	return &Server{}
}
