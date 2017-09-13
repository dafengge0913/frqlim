# frqlim
Frequency limiter  
Limit the number of access over a period of time  
# usage
```go
func main()  {
	lim := frqlim.New(4, time.Second)
	if lim.Do() {
		// do something
	} else {
		// deny
	}
}
```
