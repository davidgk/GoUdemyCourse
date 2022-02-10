package evenAndOdd

import (
	"fmt"
	"testing"
)

func TestOddAndEven(t *testing.T) {
	fmt.Println("Starting Test of OddAndEven")
	msg := testInitializeArray(t)
	msgs := []string{"when want to initialize" + msg}

	msg = testGetOdds(t)
	msgs = append(msgs, "when we want to print the odd ones"+msg)

	msg = testGetEven(t)
	msgs = append(msgs, "when we want to print the even ones"+msg)

	printTestsMessages(msgs)
}

func testGetEven(t *testing.T) string {
	sut := startOddAndEven()
	msg := " should obtain an slice with 1,3,5,7,9"
	res := sut.getEven()
	expected := []int{1, 3, 5, 7, 9}
	msg = evaluateTest(t, res, expected, msg)
	return msg
}

func evaluateTest(t *testing.T, res oddAndEven, expected []int, msg string) string {
	if len(res) != 5 {
		t.Errorf("error: expected to print %v and got %v", expected, res)
		msg = " and FAIL"
	}
	for i, val := range res {
		if val != expected[i] {
			t.Errorf("error: on index %v for val %v and obtained %v", i, val, expected[i])
			msg = " and FAIL"
		}
	}
	return msg
}

func testGetOdds(t *testing.T) string {
	sut := startOddAndEven()
	msg := " should obtain an slice with 0,2,4,6,8"
	res := sut.getOdds()
	expected := []int{0, 2, 4, 6, 8}
	msg = evaluateTest(t, res, expected, msg)
	return msg
}

func testInitializeArray(t *testing.T) string {
	msg := " should be initialized with 10 elements"
	res := startOddAndEven()
	if len(res) != 10 {
		t.Errorf("error: expected 10 elements and got %v", len(res))
		msg = " and FAIL"
	}
	return msg
}

func printTestsMessages(msgs []string) {
	for _, m := range msgs {
		fmt.Println(m)
	}
}
