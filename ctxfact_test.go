package contextcheck

import (
	"bytes"
	"encoding/gob"
	"reflect"
	"testing"
)

var testCtxFactEntries = []ctxFactEntry{
	{Key: "entry:a", Res: resInfo{Valid: true, Funcs: []string{"a"}, ReqCtx: true, EntryType: EntryWithCtx}},
	{Key: "entry:b", Res: resInfo{Valid: false, Funcs: []string{"b", "a"}, Skip: true, EntryType: EntryNormal}},
	{Key: "entry:c", Res: resInfo{Valid: true, Funcs: []string{"c"}, EntryType: EntryWithHttpHandler}},
	{Key: "entry:d", Res: resInfo{Valid: false, Funcs: []string{"d", "c"}, ReqCtx: true}},
	{Key: "entry:e", Res: resInfo{Valid: true, Funcs: []string{"e"}}},
	{Key: "entry:f", Res: resInfo{Valid: false, Funcs: []string{"f", "e"}}},
	{Key: "entry:g", Res: resInfo{Valid: true, Funcs: []string{"g"}}},
	{Key: "entry:h", Res: resInfo{Valid: false, Funcs: []string{"h", "g"}}},
	{Key: "entry:i", Res: resInfo{Valid: true, Funcs: []string{"i"}}},
	{Key: "entry:j", Res: resInfo{Valid: false, Funcs: []string{"j", "i"}}},
	{Key: "entry:k", Res: resInfo{Valid: true, Funcs: []string{"k"}}},
	{Key: "entry:l", Res: resInfo{Valid: false, Funcs: []string{"l", "k"}}},
}

func TestCtxFactGobEncodingIsDeterministic(t *testing.T) {
	want := encodeCtxFact(t, newTestCtxFact(false))

	for i := 0; i < 200; i++ {
		if got := encodeCtxFact(t, newTestCtxFact(false)); !bytes.Equal(got, want) {
			t.Fatalf("gob encoding changed on iteration %d", i)
		}
		if got := encodeCtxFact(t, newTestCtxFact(true)); !bytes.Equal(got, want) {
			t.Fatalf("gob encoding changed with reversed insertion order on iteration %d", i)
		}
	}
}

func TestCtxFactGobRoundTrip(t *testing.T) {
	fact := newTestCtxFact(false)
	encoded := encodeCtxFact(t, fact)

	var decoded ctxFact
	if err := gob.NewDecoder(bytes.NewReader(encoded)).Decode(&decoded); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(decoded, fact) {
		t.Fatalf("decoded fact mismatch:\n got: %#v\nwant: %#v", decoded, fact)
	}
}

func newTestCtxFact(reverse bool) ctxFact {
	fact := make(ctxFact, len(testCtxFactEntries))
	if reverse {
		for i := len(testCtxFactEntries) - 1; i >= 0; i-- {
			entry := testCtxFactEntries[i]
			fact[entry.Key] = entry.Res
		}
		return fact
	}
	for _, entry := range testCtxFactEntries {
		fact[entry.Key] = entry.Res
	}
	return fact
}

func encodeCtxFact(t *testing.T, fact ctxFact) []byte {
	t.Helper()

	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(&fact); err != nil {
		t.Fatal(err)
	}
	return buf.Bytes()
}
