package main

import (
	"fmt"
	"strconv"
)

type CharType string

const (
	CharTypeCap        CharType = "cap"
	CharTypeSmall      CharType = "small"
	CharTypeNum        CharType = "number"
	CharTypeOpenGroup  CharType = "open_group"
	CharTypeCloseGroup CharType = "close_group"
)

func (ct CharType) IsLastCharNum() bool {
	return ct == CharTypeNum
}

type res map[string]int

func (r res) add(elm, num string) {
	var n int = 1
	var err error

	if num != "" {
		n, err = strconv.Atoi(num)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	pn, ok := r[elm]
	if !ok {
		r[elm] = n
		return
	}
	r[elm] = pn + n
}

func (r res) update(elm, num string) {
	pn, ok := r[elm]
	if !ok {
		fmt.Errorf("invalid element to update")
		return
	}
	num = fmt.Sprintf("%d%s", pn, num)
	n, err := strconv.Atoi(num)
	if err != nil {
		fmt.Println(err)
		return
	}

	r[elm] = n
}

func main() {
	// H2(SO)4
	// H13O11
	// H2SO4
	// H2O
	// HCl

	input := "H2(SO)4"

	// form the element untill we reach either 1 number or cap letter
	response := make(res)

	// main the state of what pre element and preChar Type
	element := ""
	lastElem := ""
	var lastCharNum CharType

	isGroup := false
	groupElements := make([]string, 0)

	for _, r := range input {
		// fmt.Println(r, string(r))
		// group start
		if r == 40 {
			isGroup = true
			continue
		}
		if r == 41 && element != "" {
			groupElements = append(groupElements, element)
			element = ""
			continue
		}

		if isCap(r) {
			lastCharNum = CharTypeCap
			if element == "" {
				// append
				element = string(r)
				continue
			}
			if isGroup {
				groupElements = append(groupElements, element)
			} else {
				// add the pre element to result map
				response.add(element, "")
				lastElem = element
			}

			// now track this new element
			element = string(r)
			continue
		}

		if isSmall(r) {
			lastCharNum = CharTypeSmall
			// append
			element += string(r)
			continue
		}

		if isNumber(r) {
			// add this to result map
			// fmt.Println("calling to add ", element, string(r))
			if isGroup {
				for _, elm := range groupElements {
					response.add(elm, string(r))
				}
				isGroup = false
				lastCharNum = CharTypeNum
				continue
			}

			if lastCharNum.IsLastCharNum() {
				// update the elem freq
				response.update(lastElem, string(r))
			} else {
				response.add(element, string(r))
				lastElem = element
				element = ""
			}

			lastCharNum = CharTypeNum
			continue
		}
	}

	if element != "" {
		response.add(element, "")
		element = ""
	}

	fmt.Println("Response: ", response)
}

// CAP letters 65 - 90
// Small case 97 - 122
// numbers 48 - 57
func isCap(r rune) bool {
	if r >= 65 && r <= 90 {
		return true
	}
	return false
}

func isSmall(r rune) bool {
	if r >= 97 && r <= 122 {
		return true
	}
	return false
}

func isNumber(r rune) bool {
	if r >= 48 && r <= 57 {
		return true
	}
	return false
}
