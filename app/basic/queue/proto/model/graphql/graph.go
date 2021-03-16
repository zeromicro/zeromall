package graphql

// graphql query: 用于接收 graphql 查询结果
type Query struct {
	Markets []struct {
		ID          string
		Description string
	}
}
