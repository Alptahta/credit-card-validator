package creditcardvalidator

var testCases = []struct {
	description string
	input       string
	ok          bool
}{
	{
		"valid credit card",
		"4539 1488 0343 6467",
		true,
	},
	{
		"valid credit card",
		"4539_1488_0343_6467",
		true,
	},
	{
		"valid credit card",
		"4539-1488-0343-6467",
		true,
	},
	{
		"invalid credit card",
		"a539-1488-0343-6467",
		false,
	},
	{
		"invalid credit card",
		"4539*1488-0343-6467",
		false,
	},
	{
		"invalid credit card",
		"0000 0000 0000 0000",
		false,
	},
	{
		"invalid credit card",
		"8273 1232 7352 0569",
		false,
	},
	{
		"single digit strings can not be valid",
		"1",
		false,
	},
	{
		"a single zero is invalid",
		"0",
		false,
	},

	{
		"strings with a non-digit added at the end become invalid",
		"059a",
		false,
	},
	{
		"valid strings with punctuation included become invalid",
		"055-444-285",
		false,
	},
	{
		"single zero with space is invalid",
		" 0",
		false,
	},
	{
		"using ascii value for non-doubled non-digit isn't allowed",
		"055b 444 285",
		false,
	},
	{
		"using ascii value for doubled non-digit isn't allowed",
		":9",
		false,
	},
}
