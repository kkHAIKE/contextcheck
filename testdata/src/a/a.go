package a // want package:"ctxCheck"

import (
	"context"
	"net/http"
)

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

/* ----------------- http ----------------- */

func f8(ctx context.Context, w http.ResponseWriter, r *http.Request) {
}

func f9(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	f8(ctx, w, r)
	f8(context.Background(), w, r) // want "Non-inherited new context, use function like `context.WithXXX` or `r.Context` instead"
}

func f10(in bool, w http.ResponseWriter, r *http.Request) {
	f8(r.Context(), w, r)
	f8(context.Background(), w, r) // want "Non-inherited new context, use function like `context.WithXXX` or `r.Context` instead"
}

// nolint: contextcheck
func f14(w http.ResponseWriter, r *http.Request, err error) {
	f8(context.Background(), w, r)
}

// @contextcheck(req_has_ctx)
func f15(w http.ResponseWriter, r *http.Request, err error) {
	f8(r.Context(), w, r)
}

func f11() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		f8(r.Context(), w, r)
		f8(context.Background(), w, r) // want "Non-inherited new context, use function like `context.WithXXX` or `r.Context` instead"

		f9(w, r)

		f10(true, w, r)
		f14(w, r, nil)
		f15(w, r, nil)
	})
}

/* ----------------- generics ----------------- */

type MySlice[T int | float32] []T

func (s MySlice[T]) f12(ctx context.Context) T {
	f3() // want "Function `f3` should pass the context parameter"

	var sum T
	for _, value := range s {
		sum += value
	}
	return sum
}

func f13[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64](ctx context.Context, a, b T) T {
	f3() // want "Function `f3` should pass the context parameter"

	if a > b {
		return a
	}

	return b
}

/* ----------------- issue 21 ----------------- */

func f16(ctx context.Context, k string) func() {
	return func() { // want "Function `f16\\$1` should pass the context parameter"
		f16(context.Background(), k)
	}
}

func f17(ctx context.Context, k string) func() func() {
	return func() func() { // want "Function `f17\\$1->f17\\$1\\$1` should pass the context parameter"
		return func() {
			f16(context.Background(), k)
		}
	}
}
