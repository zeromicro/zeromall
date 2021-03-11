package graphql

// graphql query:
type Query struct {
	Markets []struct {
		ID          string
		Description string
	}
}
