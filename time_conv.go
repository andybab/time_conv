package time_conv

import (
	"fmt"
	"strings"
	"strconv"
)

var MILLIS = 0
var SEC = 1
var MIN = 2
var HOUR = 3
var DAY = 4
var WEEK = 5

var STR_MATRIX = [6]string{"MILLIS", "SEC", "MIN", "HOUR", "DAY", "WEEK"}
var CONV_MATRIX = [6]int{1, 1000, 60, 60, 24, 7}

// Unknown instruction error
type UnknownInstructionError struct {
	What string
}

func (e *UnknownInstructionError) Error() string {
	return fmt.Sprint("Unknown program instruction: ", e.What)
}

// Function to generate conversion instructions and validate provided instructions
func getInstruction(rawInstruction string) (instruction int, err error) {
	lcInst := strings.ToLower(rawInstruction)
	//TODO: switch case....
	switch lcInst {
	case "millis":
		instruction = MILLIS
	case "seconds":
		instruction = SEC
	case "minutes":
		instruction = MIN
	case "hours":
		instruction = HOUR
	case "days":
		instruction = DAY
	case "weeks":
		instruction = WEEK
	default:
		instruction = 0
		err = &UnknownInstructionError{rawInstruction}
	}

	return
}

func Convert(args *[]string) (int, error) {
	fromInstruct, toInstruct := -1, -1
	var instruct *int = &fromInstruct
	requested_num := 1 //number to be converted

	for _, value := range *args {
		num, err := strconv.Atoi(value)

		if nil == err {
			requested_num = num
			continue
		}

		//Detect 'to' keyword, swap from/to instruct
		if "to" == value {
			instruct = &toInstruct
		} else if instr, err := getInstruction(value); nil == err {
			*instruct = instr
		} else {
			return 0, err
		}
	}

	//Resolve direction, (multiplication or division)
	direction := 0
	if fromInstruct < toInstruct {
		direction = 1
	} else {
		direction = -1
	}

	//Calculate result
	var result int = requested_num
	for i := fromInstruct; i != toInstruct; i = i + direction {
		if direction < 0 { //Use multiplication
			result = result * CONV_MATRIX[i]
		} else if direction > 0 { //Use division
			result = result / CONV_MATRIX[i+direction]
		}
	}

	return result, nil
}
