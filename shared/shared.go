package shared

import (
	"encoding/base64"
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/willy182/goshare"
)

const (
	// CHARS for setting short random string
	CHARS = "abcdefghijklmnopqrstuvwxyz0123456789"
	// NUMBERS for setting short random number
	NUMBERS = "0123456789"

	// this block is for validating URL format
	ip           string = `(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))`
	urlSchema    string = `((ftp|sftp|tcp|udp|wss?|https?):\/\/)`
	urlUsername  string = `(\S+(:\S*)?@)`
	urlPath      string = `((\/|\?|#)[^\s]*)`
	urlPort      string = `(:(\d{1,5}))`
	urlIP        string = `([1-9]\d?|1\d\d|2[01]\d|22[0-3])(\.(1?\d{1,2}|2[0-4]\d|25[0-5])){2}(?:\.([0-9]\d?|1\d\d|2[0-4]\d|25[0-4]))`
	urlSubdomain string = `((www\.)|([a-zA-Z0-9]([-\.][-\._a-zA-Z0-9]+)*))`
	urlPattern   string = `^` + urlSchema + `?` + urlUsername + `?` + `((` + urlIP + `|(\[` + ip + `\])|(([a-zA-Z0-9]([a-zA-Z0-9-_]+)?[a-zA-Z0-9]([-\.][a-zA-Z0-9]+)*)|(` + urlSubdomain + `?))?(([a-zA-Z\x{00a1}-\x{ffff}0-9]+-?-?)*[a-zA-Z\x{00a1}-\x{ffff}0-9]+)(?:\.([a-zA-Z\x{00a1}-\x{ffff}]{1,}))?))\.?` + urlPort + `?` + urlPath + `?$`
	area         string = `^\+\d{1,5}$`
	phone        string = `^\d{5,}$`
)

var (
	// ErrBadFormatURL variable for error of url format
	ErrBadFormatURL = errors.New("invalid url format")
	// ErrBadFormatMail variable for error of email format
	ErrBadFormatMail = errors.New("invalid email format")
	// ErrBadFormatPhoneNumber variable for error of email format
	ErrBadFormatPhoneNumber = errors.New("invalid phone format")

	// urlRegexp regex for validate URL
	urlRegexp = regexp.MustCompile(urlPattern)
	// areaRegexp  regex for phone area number using +
	areaRegexp = regexp.MustCompile(area)
	// telpRegexp regex for phone number
	phoneRegexp = regexp.MustCompile(phone)

	// domains for list domain validate
	domains = new(collection)
)

type collection struct {
	items map[string]struct{}
	err   error
	once  sync.Once
}

// ValidateURL function for validating url
func ValidateURL(str string) error {
	if !urlRegexp.MatchString(str) {
		return ErrBadFormatURL
	}
	return nil
}

// ValidatePhoneNumber function for validating phone number only
func ValidatePhoneNumber(str string) error {
	if !phoneRegexp.MatchString(str) {
		return ErrBadFormatPhoneNumber
	}
	return nil
}

// ValidatePhoneAreaNumber function for validating area phone number
func ValidatePhoneAreaNumber(str string) error {
	if !areaRegexp.MatchString(str) {
		return ErrBadFormatPhoneNumber
	}
	return nil
}

// StringArrayReplace function for replacing whether string in array
// str string searched string
// list []string array
func StringArrayReplace(str string, listFind, listReplace []string) string {
	for i, v := range listFind {
		if strings.Contains(str, v) {
			str = strings.Replace(str, v, listReplace[i], -1)
		}
	}
	return str
}

// ValidateNumeric function for check valid numeric
func ValidateNumeric(str string) bool {
	var num, symbol int
	for _, r := range str {
		if r >= 48 && r <= 57 { //code ascii for [0-9]
			num = +1
		} else {
			symbol = +1
		}
	}

	if symbol > 0 {
		return false
	}

	return num >= 1
}

// ValidateAlphabet function for check alphabet
func ValidateAlphabet(str string) bool {
	var uppercase, lowercase, symbol int
	for _, r := range str {
		if IsUppercase(r) {
			uppercase = +1
		} else if IsLowercase(r) {
			lowercase = +1
		} else { //except alphabet
			symbol = +1
		}
	}

	if symbol > 0 {
		return false
	}
	return uppercase >= 1 || lowercase >= 1
}

// ValidateAlphabetWithSpace function for check alphabet with space
func ValidateAlphabetWithSpace(str string) bool {
	var uppercase, lowercase, space, symbol int
	for _, r := range str {
		if IsUppercase(r) { //code ascii for [A-Z]
			uppercase = +1
		} else if IsLowercase(r) { //code ascii for [a-z]
			lowercase = +1
		} else if r == 32 { //code ascii for space
			space = +1
		} else { //except alphabet
			symbol = +1
		}
	}

	if symbol > 0 {
		return false
	}
	return uppercase >= 1 || lowercase >= 1 || space >= 1
}

// ValidateAlphanumeric function for check valid alphanumeric
func ValidateAlphanumeric(str string, must bool) bool {
	var uppercase, lowercase, num, symbol int
	for _, r := range str {
		if IsUppercase(r) {
			uppercase = +1
		} else if IsLowercase(r) {
			lowercase = +1
		} else if IsNumeric(r) {
			num = +1
		} else {
			symbol = +1
		}
	}

	if symbol > 0 {
		return false
	}

	if must { //must alphanumeric
		return (uppercase >= 1 || lowercase >= 1) && num >= 1
	}

	return uppercase >= 1 || lowercase >= 1 || num >= 1
}

// ValidateAlphanumericWithSpace function for validating string to alpha numeric with space
func ValidateAlphanumericWithSpace(str string, must bool) bool {
	var uppercase, lowercase, num, space, symbol int
	for _, r := range str {
		if IsUppercase(r) { //code ascii for [A-Z]
			uppercase = +1
		} else if IsLowercase(r) { //code ascii for [a-z]
			lowercase = +1
		} else if IsNumeric(r) { //code ascii for [0-9]
			num = +1
		} else if r == 32 { //code ascii for space
			space = +1
		} else {
			symbol = +1
		}
	}

	if symbol > 0 {
		return false
	}

	if must { //must alphanumeric
		return (uppercase >= 1 || lowercase >= 1) && num >= 1 && space >= 1
	}

	return (uppercase >= 1 || lowercase >= 1 || num >= 1) || space >= 1
}

// GenerateRandomID function for generating shipping ID
func GenerateRandomID(length int, prefix ...string) string {
	var strPrefix string

	if len(prefix) > 0 {
		strPrefix = prefix[0]
	}

	yearNow, monthNow, _ := time.Now().Date()
	year := strconv.Itoa(yearNow)[2:len(strconv.Itoa(yearNow))]
	month := int(monthNow)
	RandomString := RandomString(length)

	id := fmt.Sprintf("%s%s%d%s", strPrefix, year, month, RandomString)
	return id
}

// RandomString function for random string
func RandomString(length int) string {
	rand.Seed(time.Now().UTC().UnixNano())

	charsLength := len(CHARS)
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = CHARS[rand.Intn(charsLength)]
	}
	return string(result)
}

// RandomNumber function for random number
func RandomNumber(length int) string {
	rand.Seed(time.Now().UTC().UnixNano())

	charsLength := len(NUMBERS)
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = NUMBERS[rand.Intn(charsLength)]
	}
	return string(result)
}

// StringInSlice function for checking whether string in slice
// str string searched string
// list []string slice
func StringInSlice(str string, list []string, caseSensitive ...bool) bool {
	isCaseSensitive := true
	if len(caseSensitive) > 0 {
		isCaseSensitive = caseSensitive[0]
	}

	if isCaseSensitive {
		for _, v := range list {
			if v == str {
				return true
			}
		}
	} else {
		for _, v := range list {
			if strings.ToLower(v) == strings.ToLower(str) {
				return true
			}
		}
	}

	return false
}

// IsUppercase reusable rune check if char is uppercase
func IsUppercase(r rune) bool {
	return int(r) >= 65 && int(r) <= 90
}

// IsLowercase reusable rune check if char is lowercase
func IsLowercase(r rune) bool {
	return int(r) >= 97 && int(r) <= 122
}

// IsNumeric reusable rune check if char is numeric
func IsNumeric(r rune) bool {
	return int(r) >= 48 && int(r) <= 57
}

// IsAllowedSymbol check if rune is any of
// [space, coma, ., !, ", #, $, %, &, ', (, ), *, +, -, /, :, ;, <, =, >, ?, @, [, \, ], ^, _, `, {, |, }, ~]
func IsAllowedSymbol(r rune) bool {
	m := int(r)
	return m >= 32 && m <= 47 || m >= 58 && m <= 64 || m >= 91 && m <= 96 || m >= 123 && m <= 126
}

// ValidateLatinOnly func for check valid latin only
func ValidateLatinOnly(str string) bool {
	var uppercase, lowercase, num, allowed, symbol int
	for _, r := range str {
		if IsUppercase(r) {
			uppercase = +1
		} else if IsLowercase(r) {
			lowercase = +1
		} else if IsNumeric(r) {
			num = +1
		} else if IsAllowedSymbol(r) {
			allowed = +1
		} else {
			symbol = +1
		}
	}

	if symbol > 0 {
		return false
	}

	return uppercase >= 1 || lowercase >= 1 || num >= 1 || allowed >= 0
}

// AnalyzeData function
// separation of data unique and duplicates
func AnalyzeData(array ...[]string) ([]string, []string) {
	temp := make(map[string]int)

	for _, arrs := range array {
		for _, val := range arrs {
			if _, ok := temp[val]; !ok {
				temp[val] = 1
			} else {
				temp[val] += 1
			}
		}
	}

	unique := make([]string, 0)
	duplicate := make([]string, 0)

	for name, isDuplicate := range temp {
		if isDuplicate > 1 {
			duplicate = append(duplicate, name)
		} else {
			unique = append(unique, name)
		}
	}

	return unique, duplicate
}

// IsDisabledEmail for split and validate email domain
func IsDisabledEmail(email string) bool {
	parts := strings.SplitN(email, "@", 2)
	if len(parts) != 2 {
		return false
	}
	return IsDisabledDomain(parts[1])
}

// IsDisabledDomain for validate domain
func IsDisabledDomain(domain string) bool {
	domains.once.Do(func() { domains.loadDomainList() })
	if domains.err != nil {
		return false
	}
	domain = strings.TrimSpace(domain)
	return domains.hasValidDomain(strings.ToLower(domain))
}

func (c *collection) hasValidDomain(item string) bool {
	_, ok := c.items[item]
	return ok
}

func (c *collection) loadDomainList() {
	c.items = make(map[string]struct{})
	for _, value := range goshare.DisposableDomains {
		c.items[value] = struct{}{}
	}
}

// RandomStringBase64 function for random string and base64 encoded
func RandomStringBase64(length int) string {
	rb := make([]byte, length)
	_, err := rand.Read(rb)

	if err != nil {
		return ""
	}
	rs := base64.URLEncoding.EncodeToString(rb)

	reg, _ := regexp.Compile("[^A-Za-z0-9]+")

	return reg.ReplaceAllString(rs, "")
}
