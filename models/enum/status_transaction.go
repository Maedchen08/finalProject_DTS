package enum

type StatusTransaction int

const (
	InConfirm StatusTransaction = iota + 1 //enum indec 1
	Otw
	Cancel
	Done
)

func (st StatusTransaction) EnumIndex() int {
	return int(st)
}

func (st StatusTransaction) String() string {
	return [...]string{"konfirmasi", "inprogress", "canceled", "done"}[st-1]
}
