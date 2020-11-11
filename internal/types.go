package internal

//Commit A struct representing a single commit
type Commit struct {
	revisionTimestamp string
	revision          uint
	author            string
	commitMessage     string
}

//DBInfo A struct for database info
type DBInfo struct {
	resources []string
}

//DeleteDiff A diff from a delete operation
type DeleteDiff struct {
	nodeKey uint
	deweyID string
	depth   uint
}

//InfoResult A result from the global info request
type InfoResult struct {
	name     string
	infoType string
	//*Note: in the rust code "resources" is wrapped in an "option" construct so there will likely be error handling involved
	resources []string
}

//InfoResults A complete list of the info within the info request
type InfoResults struct {
	results []InfoResult
}

//InsertDiff a diff from an insert operation
type InsertDiff struct {
	nodeKey          uint
	insertPosNodeKey uint
	insertPosition   string
	deweyID          string
	depth            uint
	insertType       string
	data             string
}

//TransactionMetadata Transaction Metadata
type TransactionMetadata struct {
	nodeKey         uint
	hash            int
	nt              NodeType
	descendantCount uint
	childCount      uint
	//*NOTE: descendantCount and childCount are only provided in the rust code if nodeType is a NodeType::Object or NodeType::Array
}

//QueryResult the result from a query
type QueryResult struct {
	revisionNumber    uint
	revisionTimestamp string
	revision          Revision
}

//Revision A timestamped revision ID
type Revision struct {
	timestamp string
	revision  uint
}

//ReplaceDiff A diff from a replace operation
type ReplaceDiff struct {
	nodeKey     uint
	replaceTyoe string
	data        string
}
