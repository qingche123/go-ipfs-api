package shell

//CopyState represent a node add object state
type CopyState uint

//CopyState types
const (
	NoUsed      CopyState = iota //the node is not used for adding
	CopyFailed                   //the node is used for adding, but add failed
	CopySuccess                  //the node is used for adding, and add success
)

//AddedResult, file hash and nodes's copy state
type AddedResult struct {
	Hash string `json:",omitempty"`
	Copy map[string]CopyState
}
