package SPGoversion

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	t.Run("Lazy",TestGetLazyInstance)
	t.Run("Hunger",TestGetHungerInstance)
}

func TestGetLazyInstance(t *testing.T) {
	obj1 := GetLazyInstance()
	obj2 := GetLazyInstance()

	fmt.Println(obj1, obj2)

	if obj1 != obj2 {
		t.Errorf("不相等")
	}
	obj1.SetName("1")
	obj2.SetName("2")
	fmt.Println(obj1, obj2)
	if obj1 != obj2 {
		t.Errorf("不相等")
	}
}
func TestGetHungerInstance(t *testing.T) {
	obj1 := GetHungerInstance()
	obj2 := GetHungerInstance()
	fmt.Println(obj1, obj2)
	if obj1 != obj2 {
		t.Errorf("不相等")
	}
	obj1.SetName("1")
	obj2.SetName("2")
	fmt.Println(obj1, obj2)
	if obj1 != obj2 {
		t.Errorf("不相等")
	}
}

func BenchmarkGetInstanceParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetLazyInstance() != GetLazyInstance() {
				b.Errorf("test fail")
			}
		}
	})
}
func BenchmarkGetLazyInstanceParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetHungerInstance() != GetHungerInstance() {
				b.Errorf("test fail")
			}
		}
	})
}