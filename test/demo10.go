package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"reflect"
	"sync"
	"time"
	"unsafe"
)

var block = "package"

const (
	one = iota + 1
	two
	three
	four
)

func main() {
	//test001()
	//testSwitch01()
	//testSwitch02()
	//testFor01()
	//testArray001()
	//testSlice001()
	//testMap001()
	//testStringtoByte()

	//测试返回值
	//result,e:= sum(-1,2)
	//if e != nil{
	//	fmt.Println("this is error")
	//	return
	//}
	//fmt.Println(result)

	//测试多个变量
	//result := sum2(1,2,3,4)
	//fmt.Println(result)

	//测试函数
	//cal := colsure()
	//fmt.Println(cal())
	//fmt.Println(cal())
	//fmt.Println(cal())

	//测试闭包
	//age := Age(25)
	//sm := Age.String
	//sm(age)

	//测试struct
	//testStruct()

	//测试接口 interface
	//p:=person{name:"wyh",age:20,addr:address{province: "浙江",city:"台州"}}
	//printString(p)
	//printString(p.addr)

	//测试自定义错误
	//sum,err:= add2(-1,0)
	//if cm,ok := err.(*commonError);ok {
	//	fmt.Println("错误代码是:",cm.errorCode,"错误信息是:",cm.errorMsg)
	//}else{
	//	fmt.Println("sum:",sum)
	//}

	//测试panic
	//defer func(){
	//	if p:=recover();p!=nil{
	//		fmt.Println(p)
	//	}
	//}()
	//connectMysql("","root","root")

	//测试go
	//testGo001()
	//testGo002()

	//测试select
	//testSelect001()

	//测试waitGroup
	//testWaitGroup001()

	//测试once
	//doOnce()

	//测试 协程中途退出 chan传输变量
	//testGoExits001()

	//测试context 多个协程的变量传输 或者停止
	//testContext001()

	//测试超时
	//testTimeout001()

	//测试point
	//testPoint001()

	//测试new函数
	//testNew001()
	//testNew002()

	//测试reflect 反射拿到value 和 type
	//testReflect001()

	//测试reflect 反射的重新设置值
	//testReflect002()

	//测试reflect 反射的接口设置
	//testReflect003()

	//测试 反射方法调用
	//testMethod001()

	//测试unsafe 指针转换
	//testUnsafe001()

	//测试 uintptr 指针转换 主要为了地址相加使用
	testUintptr001()
}

func test001() {
	fmt.Println(one, two, three, four)
}

func testSwitch01() {
	switch i := 6; {
	case i > 10:
		fmt.Println("i > 10: =", i)
	case i > 5 && i < 10:
		fmt.Println("5 < i < 10: =", i)
	default:
		fmt.Println(" i < 5: =", i)
	}
}

func testSwitch02() {
	switch 1 > 2 {
	case true:
		fmt.Println("true")
	case false:
		fmt.Println("false")
	default:
		fmt.Println("default")
	}
}

func testFor01() {
	sum := 0
	for i := 0; i <= 100; i++ {
		sum += i
	}
	fmt.Println("sum:", sum)
}

func testArray001() {
	arr := [5]string{"a", "b", "c", "d", "e"}
	for i, v := range arr {
		fmt.Printf("index:%d,value:%s\n", i, v)
	}
}

//切片
func testSlice001() {
	arr := [5]string{"a", "b", "c", "d", "e"}
	slice := arr[2:5]
	slice[1] = "f"
	fmt.Println(arr)
}

//map
func testMap001() {
	nameAgeMap := make(map[string]int)
	nameAgeMap["飞雪无情"] = 20
	age, ok := nameAgeMap["飞雪无情"]
	if ok {
		fmt.Println(age)
	}
}

func testStringtoByte() {
	s := "Hello飞雪无情"
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Println(s[0], s[5], s[15])
	for i, r := range s {
		fmt.Println(i, r)
	}
}

func sum(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("a或b不能为负数")
	}
	return a + b, nil
}

func sum2(params ...int) (sum int) {
	for _, v := range params {
		sum += v
	}
	return
}

func colsure() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func (p person) String() string {
	return fmt.Sprintf("this person name %s,age %d", p.name, p.age)
}

func printString(s fmt.Stringer) {
	fmt.Println(s.String())
}

type Age uint

func (age Age) String() {
	fmt.Printf("this is age:", age)
}

type person struct {
	name string
	age  uint
	addr address
}

type address struct {
	province string
	city     string
}

func testStruct() {
	p := person{name: "飞雪无情", age: 18, addr: address{province: "浙江", city: "台州"}}
	fmt.Println(p.name, p.age)
	fmt.Println(p.addr.province)
}

func (addr address) String() string {
	return fmt.Sprintf("this is address province %s,city:%s", addr.province, addr.city)
}

type commonError struct {
	errorCode int
	errorMsg  string
}

func (error *commonError) Error() string {
	return error.errorMsg
}

func add2(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, &commonError{errorCode: 1, errorMsg: "a或b位不能为负数"}
	}
	return a + b, nil
}

func connectMysql(ip, username, password string) {
	if ip == "" {
		panic("ip不能为空")
	}
}

func testGo001() {
	go fmt.Println("飞雪无情")
	fmt.Println("this is main")
	time.Sleep(time.Second)
}

func testGo002() {
	ch := make(chan string)
	fmt.Println("this is main")
	go func() {
		fmt.Println("this is go")
		ch <- "go 结束"
	}()
	v := <-ch
	fmt.Println("this is from go ", v)
}

func testSelect001() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)

	go func() {
		ch1 <- downLoad("first")
	}()

	go func() {
		ch2 <- downLoad("second")
	}()

	go func() {
		ch3 <- downLoad("third")
	}()

	select {
	case file := <-ch1:
		fmt.Println(file)
	case file := <-ch2:
		fmt.Println(file)
	case file := <-ch3:
		fmt.Println(file)
	}
}

func downLoad(file string) string {
	r := rand.Intn(10)
	time.Sleep(time.Duration(r) * time.Second)
	return file + "---"
}

var sum1 int
var rw sync.RWMutex

func testWaitGroup001() {
	var wg sync.WaitGroup
	wg.Add(110)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			add003(10)
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			readSum()
		}()
	}
	wg.Wait()
	fmt.Println(sum1)
}

func add003(i int) {
	rw.Lock()
	defer rw.Unlock()
	sum1 += i
	fmt.Println("this is write", sum1)
}

func readSum() {
	rw.RLock()
	defer rw.RUnlock()
	fmt.Println("this is read", sum1)
}

func doOnce() {
	var once sync.Once

	do := func() {
		fmt.Println("do once")
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(do)
			done <- true
		}()
	}

	for i := 0; i < 10; i++ {
		fmt.Println(<-done)
	}
}

func testGoExits001() {
	var wg sync.WaitGroup
	wg.Add(1)
	ch := make(chan bool)
	go func() {
		defer wg.Done()
		watchDog(ch, "dog")
	}()
	time.Sleep(4 * time.Second)
	ch <- true
	wg.Wait()
}

func watchDog(ch chan bool, name string) {
	for {
		select {
		case <-ch:
			fmt.Println("this is done")
			return
		default:
			fmt.Println("this is ", name, " 正在监控")
		}
	}
	time.Sleep(3 * time.Second)
}

func watchDog002(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("this is stop ", name)
			return
		default:
			fmt.Println("this is default ", name)
			time.Sleep(time.Second)
		}
	}
}

func testContext001() {
	ctx, stop := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		watchDog002(ctx, "watchDog002")
	}()
	time.Sleep(5 * time.Second)
	stop()
	wg.Wait()
}

func testTimeout001() {
	ch := make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		ch <- "the end"
	}()

	select {
	case v := <-ch:
		fmt.Println("this value is ", v)
	case <-time.After(5 * time.Second):
		fmt.Println("网络超时")
	}
}

func testPoint001() {
	age := 18
	modify(&age)
	fmt.Println(age)
}

func modify(age *int) {
	*age = 20
}

func testNew001() {
	var s string
	s = "wyh"
	fmt.Println(s)
}

func testNew002() {
	var sp *string
	sp = new(string)
	*sp = "wyh"
	fmt.Println(*sp)
}

func testReflect001() {
	i := 3
	iv := reflect.ValueOf(i)
	i1 := iv.Interface().(int)
	fmt.Println(i1)
}

func testReflect002() {
	i := 3
	iv := reflect.ValueOf(&i)
	iv.Elem().SetInt(4)
	fmt.Println(i)
}

type person1 struct {
	Name string
	Age  int
}

func testReflect003() {
	p := person1{Name: "飞雪无情", Age: 20}
	pt := reflect.TypeOf(p)
	stringType := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	writeType := reflect.TypeOf((*io.Writer)(nil)).Elem()
	fmt.Println("is string ", pt.Implements(stringType))
	fmt.Println("is io ", pt.Implements(writeType))
}

func (p person1) String() string {
	return fmt.Sprintf("this is person name %s,age %d", p.Name, p.Age)
}

func (p person1) Print(s string) {
	fmt.Printf("%s, name %s,age %d", s, p.Name, p.Age)
}

func testMethod001() {
	p := person1{Name: "wyh", Age: 18}
	pv := reflect.ValueOf(p)
	mPrint := pv.MethodByName("Print")
	args := []reflect.Value{reflect.ValueOf("登录")}
	mPrint.Call(args)
}

func testUnsafe001() {
	i := 1
	ip := &i
	fp := (*float64)(unsafe.Pointer(ip))
	*fp = *fp * 3
	fmt.Println(i)
}

func testUintptr001() {
	p := new(person)
	pName := (*string)(unsafe.Pointer(p))
	*pName = "wyh"
	pAge := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + unsafe.Offsetof(p.age)))
	*pAge = 18
	fmt.Println(*p)
}
