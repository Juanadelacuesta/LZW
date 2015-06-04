package main

import "fmt"

type dictionary map[string]int


func (d *dictionary) Init( ) {
	
	for i:= 0; i< 27; i++ {
		d[string(i + 'a')]=i
	}
}

func (d *dictionary) Add(substring string, int code) {
	
	d[substring]= len(d)
}

func (d *dictionary) Check(substring string)(code int, ok bool) {
	
	return int,ok = d[substring]
}

func lzw_compressor (data string) (coded_data string, codex dictionary){
	
	substring := ""
	d := &codex
	d.Init()

	for i:= range data {
		substring += data[i] 
		code, ok := d.Check(substring)
		if ok {
			code_data[i] = code

		} else {
			code:=d.Add(substring)
			code_data[i] = code
			substring := ""
		}
	}

}	

func main() {
	
	




}
