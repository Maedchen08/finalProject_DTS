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
	return [...]string{"menunggu konfirmasi anda", "agent dalam perjalanan", "dibatalkan", "selesai"}[st-1]
}
