# get-generator

This tool is inspired by the [easyjson](https://github.com/mailru/easyjson) generation logic

> note: it's in WIP interface{}, enums are not supported

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

### Installation
```
go get github.com/abba5/get-generator
go install github.com/abba5/get-generator/...@master
```

### Run
```
get-generator -all file-name.go
```


### Output

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
