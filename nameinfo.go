package nameinfo

var generators = []Generator{}

type Generator interface {
	Name() string
	Generate(firstname, lastname string) string
}
