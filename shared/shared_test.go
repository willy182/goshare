package shared

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomString(t *testing.T) {
	length := 10
	rs := RandomString(length)

	// set invalid string that
	// should not be contained in random string
	invalidString := `!@#$%^&*()_+`

	if len(rs) != length {
		t.Errorf("length of random string is not equal %d", length)
	}

	if strings.Contains(rs, invalidString) {
		t.Fatal("random string contains symbols")
	}
}

func TestValidateURL(t *testing.T) {
	var (
		url string
		err error
	)

	url = "http://www.google.com"
	if err = ValidateURL(url); err != nil {
		t.Fatal("testing 1st valid URL is not valid")
	}

	url = "www.google.com"
	if err = ValidateURL(url); err != nil {
		t.Fatal("testing 2nd valid URL is not valid")
	}

	url = "ftp://www.google.com"
	if err = ValidateURL(url); err != nil {
		t.Fatal("testing 3rd valid URL is not valid")
	}

	url = "https:///www.google.com"
	if err = ValidateURL(url); err == nil {
		t.Fatal("testing invalid URL is not valid")
	}
}

func TestValidatePhoneNumber(t *testing.T) {
	var (
		tel string
		err error
	)

	tel = "08119889788"
	if err = ValidatePhoneNumber(tel); err != nil {
		t.Fatal("testing valid phone number is not valid")
	}

	tel = "081-1988-9788"
	if err = ValidatePhoneNumber(tel); err == nil {
		t.Fatal("testing 1st invalid phone number is not valid")
	}

	tel = "0811"
	if err = ValidatePhoneNumber(tel); err == nil {
		t.Fatal("testing 2nd invalid phone number - not greater than 5 chars is not valid")
	}
}

func TestValidatePhoneAreaNumber(t *testing.T) {
	var (
		area string
		err  error
	)

	area = "+62"
	if err = ValidatePhoneAreaNumber(area); err != nil {
		t.Fatal("testing valid area number is not valid")
	}

	area = "+6 2"
	if err = ValidatePhoneAreaNumber(area); err == nil {
		t.Fatal("testing 1st invalid area number is not valid")
	}

	area = "+"
	if err = ValidatePhoneAreaNumber(area); err == nil {
		t.Fatal("testing 2nd invalid area number is not valid")
	}
}

func TestValidateNumeric(t *testing.T) {
	t.Run("Test Validate Numeric", func(t *testing.T) {
		boolFalse := ValidateNumeric("1.0.1")
		assert.False(t, boolFalse)

		boolTrue := ValidateNumeric("0123456789")
		assert.True(t, boolTrue)
	})
}

func TestValidateAlphabet(t *testing.T) {
	t.Run("Test Validate Alphabet", func(t *testing.T) {
		boolTrue := ValidateAlphabet("huFtBanGeT")
		assert.True(t, boolTrue)

		boolFalse := ValidateAlphabet("1FgH^*")
		assert.False(t, boolFalse)
	})
}

func TestValidateAlphabetWithSpace(t *testing.T) {
	t.Run("Test Validate Alphabet With Space", func(t *testing.T) {
		boolFalse := ValidateAlphabetWithSpace("huFtBanGeT*")
		assert.False(t, boolFalse)

		boolTrue := ValidateAlphabetWithSpace("huFt BanGeT")
		assert.True(t, boolTrue)
	})
}

func TestValidateAlphanumeric(t *testing.T) {
	t.Run("Test Validate Alphabet Numeric", func(t *testing.T) {
		boolTrue := ValidateAlphanumeric("okesip12", true)
		assert.True(t, boolTrue)

		boolTrue = ValidateAlphanumeric("okesip", false)
		assert.True(t, boolTrue)

		boolFalse := ValidateAlphanumeric("1FgH^*", false)
		assert.False(t, boolFalse)
	})
}

func TestValidateAlphanumericWithSpace(t *testing.T) {
	t.Run("Test Validate Alphabet Numeric With Space", func(t *testing.T) {
		boolTrue := ValidateAlphanumericWithSpace("oke sip1", false)
		assert.True(t, boolTrue)

		boolTrue = ValidateAlphanumericWithSpace("OKE sip1", false)
		assert.True(t, boolTrue)

		boolFalse := ValidateAlphanumericWithSpace("okesip1", true)
		assert.False(t, boolFalse)

		boolFalse = ValidateAlphanumericWithSpace("okesip1@", true)
		assert.False(t, boolFalse)
	})
}

func TestGenerateRandomID(t *testing.T) {
	t.Run("Test Generate Random ID", func(t *testing.T) {
		var res string
		randomID := GenerateRandomID(5)
		assert.IsType(t, res, randomID)

		randomID = GenerateRandomID(5, "00")
		assert.IsType(t, res, randomID)
	})
}

func TestRandomNumber(t *testing.T) {
	t.Run("Test Generate Random Number", func(t *testing.T) {
		var res string
		randomNumber := RandomNumber(5)

		assert.IsType(t, res, randomNumber)
	})
}

func TestStringInSlice(t *testing.T) {
	var positiveStr = "mantab"
	var positiveStrCheck = "mantul"
	type args struct {
		str           string
		list          []string
		caseSensitive []bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Testcase #1: Positive",
			args: args{str: positiveStr, list: []string{positiveStr, positiveStrCheck}},
			want: true,
		},
		{
			name: "Testcase #2: Positive",
			args: args{str: positiveStr, list: []string{positiveStr, positiveStrCheck}, caseSensitive: []bool{false}},
			want: true,
		},
		{
			name: "Testcase #3: Negative",
			args: args{str: positiveStr, list: []string{"mantap", positiveStrCheck}, caseSensitive: []bool{false}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringInSlice(tt.args.str, tt.args.list, tt.args.caseSensitive...); got != tt.want {
				t.Errorf("StringInSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateLatinOnly(t *testing.T) {
	t.Run("Test Validate Latin Only", func(t *testing.T) {
		boolFalse := ValidateLatinOnly("스칼 k4nj1 k0r34")
		assert.False(t, boolFalse)

		boolTrue := ValidateLatinOnly("okeAJ 123 ~!@#")
		assert.True(t, boolTrue)

		boolTrue = ValidateLatinOnly("okeAJ")
		assert.True(t, boolTrue)
	})
}

func TestStringArrayReplace(t *testing.T) {
	t.Run("Test String Array Replace", func(t *testing.T) {
		find := []string{"##YEAR##", "##FULLNAME##", "##URL##"}
		replacer := []string{"2012", "member", "http://asd.co"}
		content := StringArrayReplace("asdsad", find, replacer)
		assert.Equal(t, "asdsad", content)

		content2 := StringArrayReplace("##YEAR## asdsad", find, replacer)
		assert.Equal(t, "2012 asdsad", content2)
	})
}

func TestAnalyzeData(t *testing.T) {
	data1 := []string{"satu", "dua", "tiga"}
	data2 := []string{"dua", "tiga", "empat"}
	type args struct {
		array [][]string
	}
	tests := []struct {
		name  string
		args  args
		want  []string
		want1 []string
	}{
		// TODO: Add test cases.
		{
			name: "case 1",
			args: args{
				[][]string{
					data1,
					data2,
				},
			},
			want:  []string{"satu", "empat"},
			want1: []string{"dua", "tiga"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := AnalyzeData(tt.args.array...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AnalyzeData() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("AnalyzeData() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
