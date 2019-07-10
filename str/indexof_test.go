package str

import (
	"testing"
	"strings"
)

var t *testing.T

func TestIndexOf(tt *testing.T) {
	t = tt
	p0 := "1234567890"
	p1 := "XYZ"

	string0 := "" + p0 + p1
	ei0 := len(p0) + len(p1) - len(p1)

	string1 := string0 + p0 + p1
	ei1 := len(string0) + len(p0) + len(p1) - len(p1)

	substr := p1
	index := check(string1, substr, ei0, 0)

	check(string1, substr, ei0, index)
	check(string1, substr, ei1, index+1)

}

func TestInsert(tt *testing.T) {
	str := `	dev, art := blockchain.CalcBlockTaxSubsidy(cache, height, voters,
	s.server.chainParams)
	pos := blockchain.CalcStakeVoteSubsidy(cache, height,
		s.server.chainParams) * int64(voters)
	pow := blockchain.CalcBlockWorkSubsidy(cache, height, voters,
		s.server.chainParams)
	total := dev + pos + pow

	rep := pfcjson.GetBlockSubsidyResult{
		Developer: dev,
		PoS:       pos,
		PoW:       pow,
		Total:     total,
	}

	return rep, nil"
`
	expected := `	dev, art := blockchain.CalcBlockTaxSubsidy(cache, height, voters,
	s.server.chainParams)
	pos := blockchain.CalcStakeVoteSubsidy(cache, height,
		s.server.chainParams) * int64(voters)
	pow := blockchain.CalcBlockWorkSubsidy(cache, height, voters,
		s.server.chainParams)
	total := dev + pos + pow

	rep := pfcjson.GetBlockSubsidyResult{
		Developer: dev,
	XYZ
		PoS:       pos,
		PoW:       pow,
		Total:     total,
	}

	return rep, nil"
`
	lines := strings.Split(str, "\n")
	//i := IndexOf(str, "		Dev: dev,", 0)
	i := IndexOfLine(lines, "		Developer: dev,")
	lines = InsertLineAt(i, lines, "	XYZ")
	result := strings.Join(lines, "\n")

	if result != expected {
		t.Logf("  result <%s>", result)
		t.Logf("expected <%s>", expected)
		t.FailNow()
	}

}
func check(string string, substr string, expected int, offset int) int {
	index := IndexOf(string, substr, offset)
	if index != expected {
		t.Logf("string <%s>", string)
		t.Logf("substr <%s>", substr)
		t.Logf("expected <%v>", expected)
		t.Logf("   index <%v>", index)
		t.Logf("bad index")
		t.FailNow()
	}
	return index
}
