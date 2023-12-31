# modscanner
modified bufio/scanner

modscanner is a wrapper of bufio/scanner in the go standard library, but it provides additional feature to show current line number of the file.

```
s := scanner.NewScanner(r)

s.Scan()

// line number 
n := s.LineNumber()
```