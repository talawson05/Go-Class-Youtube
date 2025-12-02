package main

import (
	"fmt"
	"regexp"
	"strings"
)



func main() {	

	te := "aba abba abbba"
	re := regexp.MustCompile("b+")
	mm := re.FindAllString(te, -1)
	id := re.FindAllStringIndex(te, -1)

	fmt.Println(mm) // [b bb bbb]
	fmt.Println(id) // [[1 2] [5 7] [10 13]]

	for _, d := range id {
		fmt.Println(te[d[0]:d[1]]) // b bb bbb on new lines
	}

	up := re.ReplaceAllStringFunc(te, strings.ToUpper)
	fmt.Println(up) // aBa aBBa aBBBa

	// incomplete example
	ustr := "072665ee-a034-4cc3-a2e8-9f1822c4ebbb"
	uure := `[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}`
	ufmt := regexp.MustCompile(uure)
	if !ufmt.MatchString(ustr) {
		fmt.Printf("%s is not a UUID\n", ustr)
	}
	// could be uppercase
	// Format is really xxxxxxxx-xxxx-Vxxx-Wxxx-xxxxxxxxxxxx
	// where V is a version 1-5
	// Where W is format marker/type: 8,9,a,b

	uu := regexp.MustCompile(`^[[:xdigit:]]{8}-[[:xdigit:]]{4}-[1-5][[:xdigit:]]{3}-[89abAB][[:xdigit:]]{3}-[[:xdigit:]]{12}$`)
	var test = []string {
		"072665ee-a034-4cc3-a2e8-9f1822c4ebbb",
		"072665ee-a034-6cc3-a2e8-9f1822c4ebbb", // incorrect version
		"072665ee-a034-4cc3-72e8-9f1822c4ebbb", // incorrect type
		"072665ee-a034-4cc3-a2e8-9f1822c4ebb",  // too short
		"072665ee-a034-4cc3-a2e8-9f1822c4ebbbb", // too long - GOTCHA valid subsitring, ensure whole string is matched
		"072665ee-a034-3cc3-82e8-9f1822c4ebbb",
	}

	for i, t := range test {
		if !uu.MatchString(t){
				fmt.Println(i, t, "\tfails")
		}
	}

	// Capture groups
	var ph = regexp.MustCompile(`\(([[:digit:]]{3})\) ([[:digit:]]{3})-([[:digit:]]{4})`)
	orig := "call me at (214) 514-9548 today"
	match := ph.FindStringSubmatch(orig)
	fmt.Printf("%q\n", match)

	if len(match) > 3 {
		// match[0] is the matched string, not the groups
		fmt.Printf("+1 %s-%s-%s\n", match[1], match[2], match[3])
	}

	intl := ph.ReplaceAllString(orig, "+1 ${1}-${2}-${3}")
	fmt.Println(intl)

}