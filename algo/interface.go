package algo

// Element is any type that supports comparing
type Element interface {
	LessThan(e Element) bool
}
