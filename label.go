package burrowdb

type Label string

//storage should be
// label:node:id
// the values should be key/value of properties for querying
// prefix query -> label
// compare on key value passed in
// Maybe should be a query bucket?
// label can be nil:node:id
