package enum

type StatusTransaction int

const (
	InConfirm StatusTransaction = iota  //enum indec 1
	Confirm
	Canceled
	Done
)

func (st StatusTransaction) EnumIndex() int {
	return int(st)
}

func (st StatusTransaction) String() string {
	return [...]string{"menunggu konfirmasi agen", "agent dalam perjalanan", "dibatalkan", "selesai"}[st]
}
