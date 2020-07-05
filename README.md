# MIT 6.824: Distributed Systems (2020) Starter Code

This repository contains a better version of the code from [MIT's 6.824](https://pdos.csail.mit.edu/6.824/index.html) intended for non-MIT students who just want to dive into the code.

The modifications include:
1. Restructuring the code to use [Go modules](https://blog.golang.org/using-go-modules) and to remove errors. Some of these errors include duplicate methods on the same package, when really they should belong to different folders/packages since they should not be compiled separately anyway (see the [mr folder](mr)).
2. Editing some scripts to account for the restructuring.
3. Removing files that required Athena to use. The presence of these files will cause errors/warnings in your Go code otherwise. 
4. Removing golint warnings/errors. Most of these changes are cosmetic (changing `x += 1` to `x++`, using camelCase instead of snake_case, adding documentation for exported functions/variables). The only non-cosmetic change is replacing calls to `t.Fatalf` from non-test goroutines in [labrpc/test_test.go](labrpc/test_test.go).

## Why these changes?

The result of these changes is that __you can immediately clone this repo and use all your IDE features without noise__. Otherwise, your IDE will complain of errors from the original code (though running the terminal commands will work fine).

### Source

The original code can be found in

```
git://g.csail.mit.edu/6.824-golabs-2020 6.824
```
