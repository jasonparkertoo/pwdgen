package main

import (
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

const MIN_PWD_LEN = 12

func Status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "up"})
}

func Pwd(c *gin.Context) {
	l := c.Query("len")
	p := c.Query("chars")

	chars := getReqChars(p)

	shuffle(chars)

	n, e := strconv.Atoi(l)
	if e != nil {
		n = MIN_PWD_LEN
	}

	var sb strings.Builder
	for i := 0; i < n; i++ {
		j := rand.Intn(len(chars))
		sb.WriteRune(chars[j])
	}

	c.String(http.StatusOK, sb.String())
}

func shuffle(r []rune) {
	for i, v := range r {
		j := rand.Intn(len(r))
		tmp := v
		r[i] = r[j]
		r[j] = tmp
	}
}

func getReqChars(p string) []rune {
	var chars []rune
	for _, param := range strings.Split(p, ",") {
		chars = append(chars, CharSet(param)...)
	}
	return chars
}


func CharSet(param string) []rune {
	switch param {
	case "num":
		return []rune{'0','1','2','3','4','5','6','7','8','9'}
	case "upp":
		return []rune{'A','B','C','D','E','F','G','H','I','J','K','L','M','N','O','P','Q','R','S','T','U','V','W','X','Y','Z'}
	case "low":
		return []rune{'a','b','c','d','e','f','g','h','i','j','k','l','m','n','o','p','q','r','s','t','u','v','w','x','y','z'}
	case "sym":
		return []rune{'#','$','%','&','*','+','<','=','>','@','[','\\','^','_','`','{','|','}','~'}
	case "pun":
		return []rune{'.' ,'?' ,',' ,'!' ,'\\' ,'"' ,':' ,';' ,'\'' ,'-' ,'/' ,'(' ,')'}
	default:
		return []rune("")
	}
}
