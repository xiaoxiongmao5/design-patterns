package interfaceoptions

import (
	"fmt"
	"testing"
)

func equal[T string | int](t *testing.T, desc string, str1 T, str2 T) {
	fmt.Printf("======= %s =======\n", desc)
	if str1 == str2 {
		fmt.Printf("\n通过\n export\t%v\n got\t%v\n", str2, str1)
	} else {
		t.Errorf("\n未通过\n export\t%v\n got\t%v\n", str2, str1)
	}
	fmt.Println("")
}

func TestNewConfig(t *testing.T) {
	config := NewConfig(
		Option1{Value: "value1"},
		Option2{Value: 42},
	)
	fmt.Println(config)
	equal[string](t, "测试Option1", config.Option1, "value1")
	equal[int](t, "测试Option2", config.Option2, 42)
}
