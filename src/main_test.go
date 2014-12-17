package main

import (
	"github.com/stretchr/testify/assert"
	"ssn"
	"testing"
)

var sut, _ = ssn.NewSSN("810504-8303")
var bornAfter2000, _ = ssn.NewSSN("140202-0166")
var bornBefore2000, _ = ssn.NewSSN("141130+2951")

func TestShouldReturnCorrectYear(t *testing.T) {
	actual := sut.GetYear()
	assert.Equal(t, 1981, actual)

	actual = bornAfter2000.GetYear()
	assert.Equal(t, 2014, actual)

	actual = bornBefore2000.GetYear()
	assert.Equal(t, 1914, actual)
}

func TestShouldReturnCorrectMonth(t *testing.T) {
	actual := sut.GetMonth()
	assert.Equal(t, "May", actual)

	actual = bornAfter2000.GetMonth()
	assert.Equal(t, "February", actual)

	actual = bornBefore2000.GetMonth()
	assert.Equal(t, "November", actual)
}

func TestShouldReturnCorrectDay(t *testing.T) {
	actual := sut.GetDay()
	assert.Equal(t, 4, actual)

	actual = bornAfter2000.GetDay()
	assert.Equal(t, 2, actual)

	actual = bornBefore2000.GetDay()
	assert.Equal(t, 30, actual)
}

func TestShouldReturnCorrectCode(t *testing.T) {
	actual := sut.GetCode()
	assert.Equal(t, "8303", actual)

	actual = bornAfter2000.GetCode()
	assert.Equal(t, "0166", actual)

	actual = bornBefore2000.GetCode()
	assert.Equal(t, "2951", actual)
}

func TestErrOnEmptyInput(t *testing.T) {
	_, err := ssn.NewSSN("")
	assert.NotEqual(t, nil, err)
}

func TestErrOnNoSeparator(t *testing.T) {
	_, err := ssn.NewSSN("8105048303")
	assert.NotEqual(t, err, nil)
}

func TestErrIfThoustandIsUsed(t *testing.T) {
	_, err := ssn.NewSSN("19810504-8303")
	assert.NotEqual(t, err, nil)
}

func TestErrOnTooShort(t *testing.T) {
	_, err := ssn.NewSSN("81050-8303")
	assert.NotEqual(t, err, nil)
}

func TestErrOnWrongSeparator(t *testing.T) {
	_, err := ssn.NewSSN("810504_8303")
	assert.NotEqual(t, err, nil)
}

func TestErrOnMinusAYear(t *testing.T) {
	_, err := ssn.NewSSN("-20101-2564")
	assert.NotEqual(t, err, nil)
}

func TestErrOnMonth0(t *testing.T) {
	_, err := ssn.NewSSN("120003-2564")
	assert.NotEqual(t, err, nil)
}

func TestErrOnMonth14(t *testing.T) {
	_, err := ssn.NewSSN("121403-25644")
	assert.NotEqual(t, err, nil)
}

func TestErrOnDay0(t *testing.T) {
	_, err := ssn.NewSSN("121200-2564")
	assert.NotEqual(t, err, nil)
}

func TestErrOnDay33(t *testing.T) {
	_, err := ssn.NewSSN("121233-2564")
	assert.NotEqual(t, err, nil)
}
