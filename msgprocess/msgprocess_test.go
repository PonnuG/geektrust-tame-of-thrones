package msgprocess

import (
	"geektrust/msgprocess/cipher"
	"reflect"
	"testing"
)

type TestComparatorItem struct {
	inputs [2]map[int]int
	result bool
}

func TestReadInput(t *testing.T) {
	var expected []string
	expected = []string{"Hello, this is a test input", "This is scond line in the file with line-break"}
	inputFilename := "testInput1.txt"
	result := ReadFileInput(inputFilename)
	if !reflect.DeepEqual(expected, result) { // Reflect only used for testing purpose as bad performance
		t.Fail()
		t.Errorf("TestReadInput Arguments %s : FAILED, expected %v but got %v", inputFilename, expected, result)
	} else {
		t.Logf("TestReadInput Arguments %s : PASSED, expected %v but got %v", inputFilename, expected, result)
	}
}

func TestComparator(t *testing.T) {
	m1 := map[int]int{1: 1, 3: 1}
	m2 := map[int]int{1: 2, 3: 1} //  True Case:Both maps are same in keys && value of m1 <= val of m2
	m3 := map[int]int{4: 1, 65: 1}
	m4 := map[int]int{5: 1, 65: 1} // False Case: Different Keys
	m5 := map[int]int{87: 3, 66: 4}
	m6 := map[int]int{66: 4, 87: 2} // False Case: same keys with val m1 > val m2
	m7 := map[int]int{1: 1, 2: 1}
	m8 := map[int]int{2: 1, 1: 1} // True Case: same map i.e. nothing is different

	dataItems := []TestComparatorItem{
		{[2]map[int]int{m1, m2}, true},
		{[2]map[int]int{m3, m4}, false},
		{[2]map[int]int{m5, m6}, false},
		{[2]map[int]int{m7, m8}, true},
	}

	for _, val := range dataItems {
		actualResult := comparator(val.inputs[0], val.inputs[1])
		if actualResult != val.result {
			t.Fail()
			t.Errorf("TestComparator Arguments %v : FAILED, expected %v but got %v", val.inputs, val.result, actualResult)
		} else {
			t.Logf("TestComparator Arguments %v : PASSED, expected %v and got %v", val.inputs, val.result, actualResult)
		}
	}
}

func TestDecryptMessage(t *testing.T) {
	var testCases = []struct {
		Inp     string
		Key     int
		Decoded string
	}{
		{"ROZO", 3, "OLWL"},
	}

	for _, val := range testCases {
		res := cipher.DecryptMessage(val.Inp, val.Key)
		if res != val.Decoded {
			t.Fail()
			t.Errorf("TestDecryptMessage Arguments %v and key %v : FAILED, expected %v but got %v", val.Inp, val.Key, val.Decoded, res)
		} else {
			t.Logf("TestDecryptMessage Arguments %v and key %v : PASSED, expected %v and got %v", val.Inp, val.Key, val.Decoded, res)
		}
	}
}