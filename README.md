# pl-i18n
_pl-i18n_ is a toy programming language where all usages of native keywords such as `int`, `if`, `return`, etc. are in Korean.
I decided to make this for two reasons:
1. All programming languages are in english, and I wanted my parents to know what I do for my job. But it was a bit difficult to explain programming when there are no programming languages in korean.
2. I have always been interested in programming languages and wanted to learn more about them, and what better way to learn than to make one of my own.

## Trying it out
1. clone the repo
2. from the root directory, run `go run main.go`
3. Try the following
```
>> a = 3 // throws below error
no prefix parse function for = found
>> 정수 a = 3;
>> 정수 b = 5;
>> a + b
8
>> 정수 a = [1, 2, 3];
>> 프린트(a)
[1, 2, 3]
null
>> a[0]
1
```

## Documentation
Index of keywords:
- `int`: `정수`
- `string`: `문자`
- `float`: `부동소수`
- `booleans`: `불`
- `NULL`: `널`
- `if`: `이거면`
- `else`: `않으면`
- `else if`: `이거면 않으면`
- `func`: `펑크션`
- `return`: `산출`
- `true`: `진실`
- `false`: `거짓`
- `print`: `프린트`

### Integers
```
정수 x = 5;
```

### Floats
```
부동소수 다섯 = 5; // float five = 5;
부동소수 다섯 = 5.0; // float five = 5.0;
```

### Strings
```
문자 x = "안녕"; // string x = "hello";
```

### Booleans
```
불 x = 진실; // bool x = true;
불 y = 거짓; // bool y = false;
```

### Conditionals
```
이거면 (x) {
	정수 a = 1;
}
않으면 {
	정수 x = 2;
}

// if (x) {
//     int a = 1;
// } else {
//     int x = 2;
// }

이거면 (x == 1) {
	프린트("hello world");
}
않으면 이거면 (x == 2) {
	프린트("good bye world");
}
않으면 {
	프린트("hooray");
}

// if (x == 1) {
//     print("hello world");
// } else if (x == 2) {
//     print("good bye world");
// } else {
//     print("hooray");
// }
```

### Functions
```
펑크션 foo(문자 s, 정수 n) {
	이거면 (x == 1) {
		산출 s;
	}
	않으면 {
		산출 s + " world";
	}
}

// func foo(string s, int n) {
//     if (x == 1) {
//         return s;
//     }
//     else {
//         return s + " world";
//     }
// }
```

### Arrays
```
정수 a = [1, 2, 3];
a[0] // returns 1

// int a = [1, 2, 3];
```

### Hash Maps
```
문자 d = {"키": 5};

d["키"]; // returns 5
```

TODO: I haven't had time to create a proper identifier for dictionaries or arrays yet, so any identifier like `문자` or `정수` will do.