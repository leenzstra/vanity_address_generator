package gentype

type GenInterface interface {
	Generate() ([]byte, []byte, error)
	Check([]byte) bool
	Encode([]byte, []byte) (string, string)
}