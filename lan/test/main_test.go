package test

import (
	"fmt"
	"os"
	"testing"
)

func setup() {
	fmt.Println("before all tests")
}

func teardown() {
	fmt.Println("after all tests")
}

func Test1(t *testing.T) {
	fmt.Println("i am test1")
}

func Test2(t *testing.T) {
	fmt.Println("i am test2")
}

// 若测试中包含 TestMain, 那么运行 go test 将调用 TestMain 而非直接运行测试
func TestMain(m *testing.M) {
	setup()
	// m.Run 会触发所有测试用例的执行
	code := m.Run()
	teardown()
	//  处理返回的状态码，如果不为0，说明有用例失败。
	os.Exit(code)
}
