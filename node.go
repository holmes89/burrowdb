package burrowdb

type propertyKey string

type Node struct {
	ID         string
	Name       string
	Labels     []Label
	Properties map[string]interface{}
}

// node:id as key
// store properties as byte array
// need to wrap in transaction with label
