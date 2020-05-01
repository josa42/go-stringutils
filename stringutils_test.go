package stringutils_test

import (
	"testing"

	"github.com/josa42/go-stringutils"
)

func TestTrimPrefix(t *testing.T) {

	str1 := stringutils.TrimPrefix("")

	if str1 != "" {
		t.Error("TrimPrefix() should no nothing on empty string")
	}

	str2 := stringutils.TrimPrefix(`
		Foo
	`)

	if str2 != "Foo" {
		t.Error("TrimPrefix() should strip prefix")
	}

	str3 := stringutils.TrimPrefix(`
		{
		  "Foo": "bar"
		}
	`)

	if str3 != "{\n  \"Foo\": \"bar\"\n}" {
		t.Error("TrimPrefix() should strip prefix width diffrent prefixes")
	}

}

func TestTrimLeadingTabs(t *testing.T) {
	str1 := stringutils.TrimLeadingTabs(`
		{
		  "Foo": "bar"
		}
	`)

	if str1 != "{\n  \"Foo\": \"bar\"\n}" {
		t.Error("TrimLeadingTabs() should trim all leading tabs")
	}
}

func TestRemoveSurroundingEmptyLines(t *testing.T) {
	str1 := stringutils.RemoveSurroundingEmptyLines(`
		Foo
		Bar


	`)

	if str1 != "\t\tFoo\n\t\tBar" {
		t.Error("RemoveSurroundingEmptyLines() should remove all sourrounding empty lines")
	}
}
