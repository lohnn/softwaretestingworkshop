package ssn

import (
	"errors"
	"strconv"
	"strings"
)

/**
 * A class that represents a Swedish Social Security Number (SSN)
 *
 * Provides simple API to extract the different parts of the number
 *
 * Usage Example:
 *
 * SwedishSocialSecurityNumber ssn = new SwedishSocialSecurityNumber("060901-2829");
 * int year = ssn.getYear();
 * String month = ssn.getMonth();
 *
 */
type SwedishSocialSecurityNumber struct {
	year, day, month int
	fourLast         string
}

/**
* constructor that takes a string representation of a Swedish Social security number.
*
* The format is YYMMDD[-|+]XXXX
* YY is year from 00 to 99.
* MM is month from 01 to 12.
* DD is day from 01 to 31
*
* [-] is normally used as separator
* @Example: "060901-2829" for someone who is born 2006-09-01
* [+] is used as separator from the year a person turns 100 years old.
* @Example: "060901+2829" for someone who is born 1906-09-01
*
*  XXXX is a four character control-number
*  Note that we do not use the Luhn algorithm in this task
*
* @param String socialSecurityNumber on the format "YYMMDD[-|+]XXXX" Y,M,D must be integer numbers
*
* @throws IllegalArgumentException if string is not of the correct format "YYMMDD[-|+]XXXX"
* @throws IllegalArgumentException if month is not in the scope [1-12]
* @throws IllegalArgumentException if day is not in the scope [1-31]
*
 */
func NewSSN(socialSecurityNumber string) (ssn SwedishSocialSecurityNumber, err error) {
	err = nil
	ssn = SwedishSocialSecurityNumber{}

	if len(socialSecurityNumber) != 11 {
		return ssn, errors.New("Social security has to be 11 characters")
	}
	ssnSplit := strings.Split(socialSecurityNumber, "")

	separator := ssnSplit[6]
	if separator != "-" && separator != "+" {
		return ssn, errors.New("Need separator to be - or +")
	}

	year, err := strconv.Atoi(ssnSplit[0] + ssnSplit[1])
	if year < 0 {
		return ssn, errors.New("No negative years please...")
	}
	if separator == "-" {
		if year > 14 {
			year += 1900
		} else {
			year += 2000
		}
	} else { //if "+"
		if year > 14 {
			year += 1800
		} else {
			year += 1900
		}
	}

	month, err := strconv.Atoi(ssnSplit[2] + ssnSplit[3])
	if month < 1 || month > 12 {
		return ssn, errors.New("Month between 1 and 12")
	}
	day, err := strconv.Atoi(ssnSplit[4] + ssnSplit[5])
	if day < 1 || day > 31 {
		return ssn, errors.New("Month between 1 and 12")
	}

	fourLast := ssnSplit[7] + ssnSplit[8] + ssnSplit[9] + ssnSplit[10]

	if err != nil {
		return ssn, errors.New("The inputted social security number is not correct")
	}
	return SwedishSocialSecurityNumber{year: year, month: month, day: day, fourLast: fourLast}, nil
}

/**
 * A class that represents a Swedish Social Security Number (SSN)
 *
 * Provides simple API to extract the different parts of the number
 *
 * Usage Example:
 *
 * SwedishSocialSecurityNumber ssn = new SwedishSocialSecurityNumber("060901-2829");
 * int year = ssn.getYear();
 * String month = ssn.getMonth();
 *
 */
func (ssn SwedishSocialSecurityNumber) GetYear() int {
	return ssn.year
}

/**
 * gets the month of birth as a string
 *
 * The months are January, February, Mars, April, May,
 * June, July, August, September, October, November, December
 *
 * @Example:For "800901-2829" it will return "September"
 *
 * @return A string representation of the month of birth
 */
func (ssn SwedishSocialSecurityNumber) GetMonth() string {
	month := []string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}
	return month[ssn.month-1]
}

/**
 * gets the day of birth as a number
 *
 * @Example: For "800901-2829" it will return 1
 *
 * @return int day of month [1-31]
 */
func (ssn SwedishSocialSecurityNumber) GetDay() int {
	return ssn.day
}

/**
 * gets the last four letters in the Swedish Security Number (SSN)
 *
 * @example: for the SSN "800909-2829" this method will return "2829"
 *
 * @return String returns the last 4 characters in a SSN
 */
func (ssn SwedishSocialSecurityNumber) GetCode() string {
	return ssn.fourLast
}
