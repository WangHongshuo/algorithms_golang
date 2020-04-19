// Package sortalg : sort slgorithms
package sortalg

// Interface : common sort interface
type Interface interface {
	Len() int
	Swap(i, j int)
	Less(i, j int) bool
}
