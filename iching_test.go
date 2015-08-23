package iching

import (
	"math/rand"
	"testing"
	"time"
)

func TestRoundtrip(t *testing.T) {
	rand.Seed(time.Now().Unix())
	for i := 0; i < 1000; i++ {
		n := rand.Uint32()
		ich := Itoiching(uint64(n))
		m, _ := Ichingtoi(ich)
		t.Logf("0%o → %s\n", n, ich)
		if n != uint32(m) {
			t.Fatalf("0%o → %s → 0%o\n", n, ich, m)
		}
	}
}

func TestPadding(t *testing.T) {
	for i := 0; i < 1000; i++ {
		ich := Itoiching(uint64(i))
		ichp := QianPad(ich, 3)
		t.Logf("0%o → %s → %s\n", i, ich, ichp)
		if len([]rune(ichp)) != 3 {
			t.Fatalf("0%o → %s → %s\n", i, ich, ichp)
		}
	}

	ichp := QianPad(Itoiching(07777), 2)
	if len([]rune(ichp)) != 2 {
		t.Fatalf("bad padding %s\n", ichp)
	}

	ichp = QianPad(Itoiching(077777), 2)
	if len([]rune(ichp)) != 3 {
		t.Fatalf("bad padding %s\n", ichp)
	}

}
