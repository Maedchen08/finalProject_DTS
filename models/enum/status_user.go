package enum

type StatusUser int

const(

	Active StatusUser = iota + 1 //enum indec 1
	NonActive
)

func( ss StatusUser ) EnumIndex() int {
	return int (ss)
}

func (ss StatusUser ) String() string {
	return [...]string{"active", "nonactive"}[ss-1]
}