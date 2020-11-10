package internal

// NodeType this type is going to be used like an enum in Java or rust (thats the hope anyway)
type NodeType int

const (
	//Array The Node is an Array type
	Array NodeType = iota
	//BooleanValue The Node is a BooleanValue type
	BooleanValue
	NullValue
	NumberValue
	Object
	ObjectBooleanValue
	ObjectKey
	ObjectNullValue
	ObjectStringValue
	StringValue
)
