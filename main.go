package main

type ThisIsGo {

}

func (t *ThisIsGo) GoStyle(style string) (string, err){
	if style == "go"{
		return "Go Style", nil
	} else {
		return "",errors.new("not go style")
	}
}

func main(){
	thisIsGo = ThisIsGo{}
	style,err := thisIsGo.GoStyle()
	if err != nil{
		panic("error occured")
	}
}