# 고 명명 규칙
https://blog.billo.io/devposts/golang_naming_convention/

## Varaible, Function 
Public 은 대문자 시작   
Private 는 소문자 시작  
줄임말의 경우 다 대문자 작성.  
(URL, IP 등)

const 의 경우 다른 언어와 다르게 SOMETHING, SOME_THING 처럼 사용하지 않고  
일반 변수와 동일하게 작명한다.

```go
package something

const Name := "SOME"
const location := "thing"

var thing := "thing"
var requestIP := "is ip"

func something():
    something := "is something" 

func Anything():
    anything := "is anything"
```

## Struct 
Public 은 대문자 시작  
Private 는 소문자 시작
```go
type SampleStruct struct{
    PublicProp string
    privateProp string
}
```

## Package
Package 명은 소문자, 한단어로 지어야 한다.  
그리고 Package 명은 디렉토리 이름과 동일 해야 하며  
Package 안에 Package 경우 디렉토리명을 서로 다르게 해야 한다.  
또한 당연하게도 알려진 패키지와 이름이 동일하지 않도록 조심 해야 한다.


## Interface
Go의 Interface는 행동을 규약한다.  
그래서 Interface 명의 +er 을 붙여준다.  
(Reader, Closer 등)
```go
package main

import "fmt"

type Reader interface {
	Read() string
}

type MyFile struct {
	path string
}

func NewMyFile(path string) *MyFile {
	mf := MyFile{path: path}
	return &mf
}

func (mf *MyFile) Read() string {
	fmt.Println("file read result")
	return "file read result"
}

func main() {
	var myFile Reader
	myFile = NewMyFile("path")
	myFile.Read()
}
```

## Must Function
함수 이름 앞에 Must 가 붙은 경우에는  
함수 실행 동작과정에서 실패할 경우 panic을 내야 한다는 규칙.
(공식은 아니지만 많이 이렇게 함)