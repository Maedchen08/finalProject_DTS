package enum

type Role int

const(

	Agent StatusUser = iota + 1 //enum indec 1
	Customer
)

func( ss Role ) EnumIndex() int {
	return int (ss)
}

func (ss Role ) String() string {
	return [...]string{"Agent", "Customer"}[ss-1]
}