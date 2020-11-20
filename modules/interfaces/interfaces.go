package interfaces

// consider the functionality(actions) that is common between datatypes
// interface{} is not any type but a type of interface{}
// interfaces are two words wide (type, value)
// returning an interface{} is generally not practical
// accepting an interface{}, like marshalJSON(interface{}), is more practical
//

// Processor
type Processor interface {
	Process(int)
}
