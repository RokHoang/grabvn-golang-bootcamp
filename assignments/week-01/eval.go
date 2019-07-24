package main
import (
	"fmt"
	"strconv"
	"strings"
	"errors"
	"regexp"
)

func eval(text string) error {

	s := strings.Split(trim(text)," ")
	if (len(s)<3){
		return errors.New("Too few arguments")
	}

	if (len(s)>3){
		return errors.New("Too many arguments")
	}

	t1,op,t2 := s[0], s[1], s[2]

	n1, errN1 := strconv.Atoi(t1)
	if (errN1 != nil){
		return errN1
	}

	n2, errN2 := strconv.Atoi(t2)
	if (errN2 != nil){
		return errN2
	}

	switch op{
	case "+": fmt.Println(text,"=",n1+n2)
	case "-": fmt.Println(text,"=",n1-n2)
	case "*": fmt.Println(text,"=",n1*n2)
	case "/": {
		result, errDivide := divide(n1,n2)
		if (errDivide == nil){
			fmt.Println(text,"=",result)
		} else {
			return errDivide
		}
	}
	default: return errors.New("Wrong operator")
	}
	return nil
}

func trim(text string) string {
	space := regexp.MustCompile(`\s+`)
	return space.ReplaceAllString(text, " ")
}

func divide(n1 int, n2 int) (float64, error){
	if (n2 == 0){
		return 0,errors.New("Cannot be divided by 0 (zero)")
	}

	return float64(n1)/float64(n2),nil
}