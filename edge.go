package burrowdb

type Edge struct {
	Name string
	To   Node
	From Node
}

// Not planning on adding properities
// not sure what to do here since we can use edge to find to and from
// edgename:to:from ?, don't need to store?
// edge:to -> []from?
