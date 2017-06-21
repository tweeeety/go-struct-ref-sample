package main

import "log"

type Data struct {
	Name string
	Age  int64
}

// #1.
//   arg :値渡し
//   処理:値のままプロパティアクセスして代入
func changeNameArgNormalAccessNormal(d Data) {
	d.Name = "fuga"
}

// #2.
//   arg :値渡し
//   処理:ポインタにしてからプロパティアクセスして代入
func changeNameArgNormalAccessPointer(d Data) {
	dPointer := &d
	dPointer.Name = "fuga"
}

// #3.
//   arg :ポインタ渡し
//   処理:値にしてからプロパティアクセスして代入
func changeNameArgPointerAccessNormal(dPointer *Data) {
	d := *dPointer
	d.Name = "fuga"
}

// #4.
//   arg :ポインタ渡し
//   処理:ポインタのままプロパティアクセスして代入
func changeNameArgPointerAccessPointer(dPointer *Data) {
	dPointer.Name = "fuga"
}

func main() {
	// #1
	d1 := Data{Name: "hoge", Age: 20}
	changeNameArgNormalAccessNormal(d1)
	log.Printf("d1: %+v", d1)

	// #2
	d2 := Data{Name: "hoge", Age: 20}
	changeNameArgNormalAccessPointer(d2)
	log.Printf("d2: %+v", d2)

	// #3
	d3 := Data{Name: "hoge", Age: 20}
	changeNameArgPointerAccessNormal(&d3)
	log.Printf("d3: %+v", d3)

	// #4
	d4 := Data{Name: "hoge", Age: 20}
	changeNameArgPointerAccessPointer(&d4)
	log.Printf("d4: %+v", d4)
}
