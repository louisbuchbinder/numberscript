# Numberscript
A stack based esoteric programming language.

|character|meaning|
|---------|-------|
|0 1 2 3 4 5 6 7 8 9|number characters are interpreted literally, e.g. 123|
|+ - / * %|add, subtract, divide, multiply, modulo the following value with the stack value|
|>|increment the stack by one|
|<|decrement the stack by one|
|π|print the stack value to standard out|
|( )|do while stack value != 0 loop code inside parenthesis
|? , .|conditional if/else block. Think of the operators like "? == if", ", == else", and ". == end". The condition will check if the stack value is nonzero and if true execute the truth block otherwise execute the false block. e.g. "?>π,>1π." or "if stack_val != 0 then increment the stack pointer and print. else increment the stack pointer increment the stack value and print"|
|=|comment marker, everything to the right of = will be ignored. All whitespace is ignored|

### hello world

hello.ns
```
>10>33>100>108>114>111>87>32>111>108>108>101>72(π<)
```
output
```
Hello World!
```

### install
```bash
go install github.com/louisbuchbinder/numberscript
```

### usage
```bash
numberscript hello.ns
```

### more [examples](/test)

