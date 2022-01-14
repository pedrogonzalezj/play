package custom

import "testing"

func TestBitcoinIsANumber(t *testing.T) {
	var got interface{} = Bitcoin(10)
	if v, ok := got.(Bitcoin); !ok {
		t.Errorf("got %v want Bitcoint", v)
	}
	got = Bitcoin(20)
	if v, ok := got.(Bitcoin); !ok {
		t.Errorf("got %v want Bitcoint", v)
	}
}
func TestBitcoinStringFormat(t *testing.T) {
	tenBTC := Bitcoin(10)
	got := tenBTC.String()
	want := "10 BTC"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
