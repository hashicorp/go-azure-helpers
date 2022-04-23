package normalizers

type Normalize interface {
	Normalize[T]() T
}
