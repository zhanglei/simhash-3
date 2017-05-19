package simhash

import (
	"log"
	"strconv"
	"testing"
)

func init() {
	if err := LoadDictionary("jieba.dict.txt", "idf.txt", "stop_words.txt"); err != nil {
		log.Fatal("Failed to load dictionary:", err)
	}
}

func TestSimhash(t *testing.T) {
	s := `江南皮革厂倒闭了`
	Simhash(&s, -1)
}

func BenchmarkSimhash(b *testing.B) {
	s := `江南皮革厂倒闭了`
	for i := 0; i < b.N; i++ {
		Simhash(&s, -1)
	}
}

func TestCalWeights(t *testing.T) {
	hashes := []hashWeigth{}
	weights := calWeights(hashes)
	AssertIntEqual(t, len(weights), 64)
	if len(weights) != 64 {
		t.Error("weights length should be 64")
	}
	for _, weight := range weights {
		AssertFloat64Equal(t, 0.0, weight)
	}

	hashes = []hashWeigth{{1, 11.11}, {1 << 1, 0.5}, {1 << 63, 30.3}}
	weights = calWeights(hashes)
	AssertFloat64Equal(t, weights[0], 11.11-0.5-30.3)
	AssertFloat64Equal(t, weights[1], -11.11+0.5-30.3)
	for i := 2; i < 63; i++ {
		AssertFloat64Equal(t, weights[i], -11.11-0.5-30.3)
	}
	AssertFloat64Equal(t, weights[63], -11.11-0.5+30.3)
}

func TestFingerprint(t *testing.T) {
	weights := [64]float64{
		-1.0, -2.0, -3.0, -4.0, -5.0, -6.0, -7.0, -8.0,
		0.0, 1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0,
		-1.0, -2.0, -3.0, -4.0, -5.0, -6.0, -7.0, -8.0,
		0.0, 1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0,
		-1.0, -2.0, -3.0, -4.0, -5.0, -6.0, -7.0, -8.0,
		0.0, 1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0,
		-1.0, -2.0, -3.0, -4.0, -5.0, -6.0, -7.0, -8.0,
		0.0, 1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0,
	}
	fp := fingerprint(weights)
	AssertStringEqual(t, strconv.FormatUint(fp, 2), "1111111100000000111111110000000011111111000000001111111100000000")

	weights = [64]float64{
		-1.0, -2.0, -3.0, -4.0, -5.0, -6.0, -7.0, -8.0,
		0.0, 1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0,
		-1.0, -2.0, -3.0, -4.0, -5.0, -6.0, -7.0, -8.0,
		0.0, 1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0,
		-1.0, -2.0, -3.0, -4.0, -5.0, -6.0, -7.0, -8.0,
		0.0, 1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0,
		-1.0, -2.0, -3.0, -4.0, -5.0, -6.0, -7.0, -8.0,
		-1.0, -2.0, -3.0, -4.0, -5.0, -6.0, -7.0, -8.0,
	}
	fp = fingerprint(weights)
	AssertStringEqual(t, strconv.FormatUint(fp, 2), "111111110000000011111111000000001111111100000000")
	AssertUint64Equal(t, fp, uint64(280379743338240))
}

func AssertIntEqual(t *testing.T, a, b int) {
	if a != b {
		t.Errorf("%d should equal %d", a, b)
	}
}

func AssertUint64Equal(t *testing.T, a, b uint64) {
	if a != b {
		t.Errorf("%d should equal %d", a, b)
	}
}

func AssertFloat64Equal(t *testing.T, a, b float64) {
	if a != b {
		t.Errorf("%f should equal %f", a, b)
	}
}

func AssertStringEqual(t *testing.T, a, b string) {
	if a != b {
		t.Errorf("%f should equal %f", a, b)
	}
}
