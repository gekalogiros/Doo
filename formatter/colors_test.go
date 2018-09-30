package formatter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const content = "Text"

func TestDisplay(t *testing.T) {
	actual := display(yellow, 0, content)
	expected := "\033\\[0;33mText\033\\[0m"
	assert.Regexp(t, expected, actual)
}

func TestRed(t *testing.T) {
	testNormalTextWithColor(t, 31, Red)
}

func TestGreen(t *testing.T) {
	testNormalTextWithColor(t, 32, Green)
}

func TestYellow(t *testing.T) {
	testNormalTextWithColor(t, 33, Yellow)
}

func TestBRed(t *testing.T) {
	testBoldTextWithColor(t, 31, 1, BRed)
}

func TestBGreen(t *testing.T) {
	testBoldTextWithColor(t, 32, 1, BGreen)
}

func TestBYellow(t *testing.T) {
	testBoldTextWithColor(t, 33, 1, BYellow)
}

func testBoldTextWithColor(t *testing.T, code int, strength int, underTest func(string, int) string) {

	actualColorCodePointer, actualFontSizePointer, actualContentPointer := verifyDisplay()

	underTest(content, strength)

	assert.Equal(t, code, *actualColorCodePointer)
	assert.Equal(t, strength, *actualFontSizePointer)
	assert.Equal(t, content, *actualContentPointer)
}

func testNormalTextWithColor(t *testing.T, code int, underTest func(string) string) {

	actualColorCodePointer, actualFontSizePointer, actualContentPointer := verifyDisplay()

	underTest(content)

	assert.Equal(t, code, *actualColorCodePointer)
	assert.Equal(t, 0, *actualFontSizePointer)
	assert.Equal(t, content, *actualContentPointer)
}

func verifyDisplay() (*int, *int, *string){

	actualColorCode, actualFontSize := 0, 0
	actualContent := ""
	output := ""

	display = func(colorCode int, fontSize int, content string) string {
		actualColorCode = colorCode
		actualFontSize = fontSize
		actualContent = content
		return output
	}

	return &actualColorCode, &actualFontSize, &actualContent
}
