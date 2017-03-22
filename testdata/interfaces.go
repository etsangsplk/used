package pkg

type I interface {
	fn1()
}

type t struct{}

func (t) fn1() {}
func (t) fn2() {} // MATCH /fn2 is used 0 times/

func init() {
	var _ I
	var _ t
}
