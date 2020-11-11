package internal

// NodeType this type is going to be used like an enum in Java or rust (thats the hope anyway)
type NodeType int

//Insert All possible options for a resource update
type Insert int

//DBType All possible options database (and resource) types
type DBType int

//MetaData The scope of the metadata to return
type MetaData int

const (
	//Array The Node is an Array type
	Array NodeType = iota

	//BooleanValue The Node is a BooleanValue type
	BooleanValue

	//NullValue The Node is a NullValue type
	NullValue

	//NumberValue The Node is a NumberValue type
	NumberValue

	//Object The Node is an Object type
	Object

	//ObjectBooleanValue The Noide is of type ObjectBooleanValue
	ObjectBooleanValue

	//ObjectKey The Node is of type ObjectKey
	ObjectKey

	//ObjectNullValue The Node has a type ObjectNullValue
	ObjectNullValue

	//ObjectStringValue The Node has a type of ObjectStringValue
	ObjectStringValue

	//StringValue The Node is type StringValue
	StringValue
)

const (
	//Child do a child insert
	Child Insert = iota
	//Left Do a left insert
	Left
	//Right Do a right insert
	Right
	//Replace Do a replace
	Replace
)

const (
	//JSON Json Data
	JSON DBType = iota
	//XML XML Data
	XML
)
const (
	//All Returns all Metadata
	All MetaData = iota
	//Key Returns Key Metadata
	Key
	//KeyAndChild Returns KeyAndChild metadata
	KeyAndChild
)

var strInserts = [...]string{
	"Child",
	"Left",
	"Right",
	"Replace",
}

var strNodeType = [...]string{
	"Array",
	"BooleanValue",
	"NullValue",
	"NumberValue",
	"Object",
	"ObjectBooleanValue",
	"ObjectKey",
	"ObjectNullValue",
	"ObjectStringValue",
	"StringValue",
}

var strDBTypes = [...]string{
	"JSON",
	"XML",
}

var strMD = [...]string{
	"ALL",
	"Key",
	"KeyAndChild",
}

//*These functions are probably not necessary...

//GetNodeType Returns the String representation of the "enum" field for NodeType
func (nt NodeType) GetNodeType() string {
	return strNodeType[nt]
}

//GetInsertVal Returns the String representation of the "enum" field for Insert
func (i Insert) GetInsertVal() string {
	return strInserts[i]
}

//GetDBType Fetches the string representation of the "enum" field for DBType
func (db DBType) GetDBType() string {
	return strDBTypes[db]
}

//GetMetadataType Returns the string representation of the "enum" field for the scope of the metadata to return
func (md MetaData) GetMetadataType() string {
	return strMD[md]
}
