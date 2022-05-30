# get-generator

### WIP: 

pointer cause panic in golang, 
if it's not handle correctly

aim to generate get methods for the structure fields.

example

```
type Student struct {
	Name string
	Id   int
}
```

then generated method will be 

```
func (s *Student) GetName() string {
	if s != nil {
		return s.Name
	}
	return ""
}

func (s *Student) GetId() int {
	if s != nil {
		return s.Id
	}
	return 0
}
```
