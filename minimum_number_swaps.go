/*
Minimum Number of Swaps to Make the String Balanced
You are given a 0-indexed string s of even length n. The string consists of exactly n / 2 opening brackets '[' and n / 2 closing brackets ']'.

A string is called balanced if and only if:

- It is the empty string, or
- It can be written as AB, where both A and B are balanced strings, or
- It can be written as [C], where C is a balanced string.
You may swap the brackets at any two indices any number of times.

Return the minimum number of swaps to make s balanced.

Example 1:
  Input: s = "][]["
  Output: 1
  Explanation: You can make the string balanced by swapping index 0 with index 3.
  The resulting string is "[[]]".

Example 2:
  Input: s = "]]][[["
  Output: 2
  Explanation: You can do the following to make the string balanced:
    - Swap index 0 with index 4. s = "[]][][".
    - Swap index 1 with index 5. s = "[[][]]".
  The resulting string is "[[][]]".

Example 3:
  Input: s = "[]"
  Output: 0
  Explanation: The string is already balanced.

Observations:
  - Equal count of opening & closing brackets
  - We don´t care about valid part of the input. If it´s ok, we just ignore it and we only care about incorrect ones
*/

type TestCase struct {
  input string
  expectedValue int
}

func main() {
  inputCases := map[string]TestCase{
    {"case1": TestCase{input:"][][", expectedValue:1}},
    {"case2": TestCase{input:"]]][[[", expectedValue:2}},
    {"case3": TestCase{input:"[]", expectedValue:0}},
  }

  desiredCase := inputCases["case1"]
  result := minSwaps(desiredCase.input

  fmt.Println("Running input: %v - result: %v - expected: %v", desiredCase.input, result, desiredCase.expectedValue)
}

func minSwaps(s string) int {
	if len(s) == 0 {
		return 0
	}

	var inbalance int
	for i := 0; i < len(s); i++ {
		if isOpeningBracket(string(s[i])) {
            inbalance++
        } else if isClosingBracket(string(s[i])) && inbalance > 0 {
            inbalance--
        } else if isClosingBracket(string(s[i])) {
            inbalance++
        }
	}

	fmt.Printf("total inbalance: %v", inbalance)

    // we divide in two because are always pairs
	return inbalance/2

}

func isOpeningBracket(value string) bool {
	return value == "["
}

func isClosingBracket(value string) bool {
	return value == "]"
}
