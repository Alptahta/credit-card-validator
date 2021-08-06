# credit-card-validator

Credit Card Validation Library with Go programing language

# usage

go get example.com/mymodule@v0.1.0

# valid syntaxes

    {
    	"valid credit card",
    	"4539 1488 0343 6467",
    },
    {
    	"valid credit card",
    	"4539_1488_0343_6467",
    },
    {
    	"valid credit card",
    	"4539-1488-0343-6467",
    },
    {
    	"invalid credit card",
    	"a539-1488-0343-6467",
    },
    {
    	"invalid credit card",
    	"4539*1488-0343-6467",
    },
