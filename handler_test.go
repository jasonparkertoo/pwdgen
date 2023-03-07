package main

import (
	"regexp"
	"strconv"
	"strings"
	"testing"
)

// test requested characters are returned
func TestGetReqCharsParamNum(t *testing.T) {
	param := "num"

	exp := regexp.MustCompile(`[0-9]`)
	act := getReqChars(param)

	failEmptyRuneSlice(act, param, t)

	for _, v := range act {
		if !exp.MatchString(string(v)) {
			t.Fatalf(`getReqChar("num") = returned non-numeric value %v`, strconv.QuoteRune(v))
		}
	}
}

// test requested characters are returned
func TestGetReqCharsParamUpp(t *testing.T) {
	param := "upp"

	exp := regexp.MustCompile(`[A-Z]+`)
	act := getReqChars(param);

	failEmptyRuneSlice(act, param, t)

	for _, v := range act {
		if !exp.MatchString(string(v)) {
			t.Fatalf(`getReqChar("` + param + `") = returned non-uppercase letter: %v`, strconv.QuoteRune(v))
		}
	}
}

func TestGetReqCharsParamLow(t *testing.T) {
	param := "low"

	exp := regexp.MustCompile(`[a-z]+`)
	act := getReqChars(param)

	failEmptyRuneSlice(act, param, t)

	for _, v := range act {
		if !exp.MatchString(string(v)) {
			t.Fatalf(`getReqChar("` + param + `") = returned non-lowercase letter: %v`, strconv.QuoteRune(v))
		}
	}
}

func TestGetReqCharsParamSym(t *testing.T) {
	param := "sym"

	exp := string(CharSet(param))
	act := getReqChars(param)

	failEmptyRuneSlice(act, param, t)

	for _, v := range act {
		if !strings.ContainsRune(exp, v) {
			t.Fatalf(`getReqChar("` + param + `") = returned symbol letter: %v`, strconv.QuoteRune(v))
		}
	}
}

func TestGetReqCharsParamPun(t *testing.T) {
	param := "pun"
	
	exp := string(CharSet(param))
	act := getReqChars(param)

	failEmptyRuneSlice(act, param, t)

	for _, v := range act {
		if !strings.ContainsRune(exp, v) {
			t.Fatalf(`getReqChar("` + param + `") = returned non-uppercase letter: %v`, strconv.QuoteRune(v))
		}
	}
}

func failEmptyRuneSlice(s []rune, param string, t *testing.T) {
	if len(s) < 1 || s == nil {
		t.Fatal(`getReqChar("` + param + `") = returned and empty slice`)
	}
}
