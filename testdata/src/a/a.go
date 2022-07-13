package a // want package:"ctxCheck"

import "context"

type MyString string

type TestInterface interface {
	Test() string
}

type xx struct{}

func newXX() TestInterface {
	return &xx{}
}

func (x xx) Test() string {
	return ""
}

type MyInt int

func (x MyInt) F() int {
	return int(x)
}

func f1(ctx context.Context) {
	defer f2(ctx)
	go f2(ctx)
	f2(ctx)

	ctx = context.WithValue(ctx, MyString("aaa"), "aaaaaa")
	f2(ctx)

	newXX().Test()

	f3() // want "Function `f3` should pass the context parameter"
	f6() // want "Function `f6->f3` should pass the context parameter"

	defer func() {
		f2(ctx)
	}()

	func(ctx context.Context) {
		f2(ctx)
	}(ctx)

	f2(context.Background()) // want "Non-inherited new context, use function like `context.WithXXX` instead"

	thunk := MyInt.F
	thunk(0)

	bound := MyInt(0).F
	bound()
}

func f2(ctx context.Context) {}

func f3() {
	f2(context.TODO())
}

func f4(ctx context.Context) {
	f2(ctx)
	ctx = context.Background()
	f2(ctx) // want "Non-inherited new context, use function like `context.WithXXX` instead"
}

func f5(ctx context.Context) {
	func() {
		f2(ctx)
	}()

	ctx = context.Background() // want "Non-inherited new context, use function like `context.WithXXX` instead"
	f2(ctx)
}

func f6() {
	f3()
}

func f7(ctx context.Context) {
	ctx, cancel := getNewCtx(ctx)
	defer cancel()

	f2(ctx) // OK
}

func getNewCtx(ctx context.Context) (newCtx context.Context, cancel context.CancelFunc) {
	return context.WithCancel(ctx)
}
