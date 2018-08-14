package main

import (
	"io"
	"time"
	"math/rand"

	"github.com/qingche123/go-ipfs-api"
	u "github.com/ipfs/go-ipfs-util"
	"fmt"
	"os"
	"bytes"
)

var sh *shell.Shell
var ncalls int

var _ = time.ANSIC

func sleep() {
	ncalls++
	//time.Sleep(time.Millisecond * 5)
}

func randString() string {
	alpha := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	l := rand.Intn(10) + 2

	var s string
	for i := 0; i < l; i++ {
		s += string([]byte{alpha[rand.Intn(len(alpha))]})
	}
	return s
}

func makeRandomObject() (string, error) {
	// do some math to make a size
	x := rand.Intn(120) + 1
	y := rand.Intn(120) + 1
	z := rand.Intn(120) + 1
	size := x * y * z

	r := io.LimitReader(u.NewTimeSeededRand(), int64(size))
	sleep()
	return sh.Add(r)
}

func makeRandomDir(depth int) (string, error) {
	if depth <= 0 {
		return makeRandomObject()
	}
	sleep()
	empty, err := sh.NewObject("unixfs-dir")
	if err != nil {
		return "", err
	}

	curdir := empty
	for i := 0; i < rand.Intn(8)+2; i++ {
		var obj string
		if rand.Intn(2) == 1 {
			obj, err = makeRandomObject()
			if err != nil {
				return "", err
			}
		} else {
			obj, err = makeRandomDir(depth - 1)
			if err != nil {
				return "", err
			}
		}

		name := randString()
		sleep()
		nobj, err := sh.PatchLink(curdir, name, obj, true)
		if err != nil {
			return "", err
		}
		curdir = nobj
	}

	return curdir, nil
}

func localCopyTest(){
	fmt.Println("LocalTest")
	sh = shell.NewShell("127.0.0.1:5001")
	f, err := os.Open("test")
	if err != nil {
		fmt.Println(err.Error())
	}
	copyNodes := make([]string, 1)
	copyNodes[0] = "127.0.0.1:5002"
	hash, err := sh.AddAndCopy(f, 1, copyNodes)
	if err!= nil {
		fmt.Println(err.Error())
	}
	fmt.Println(hash)
}


func remoteCopyTest(){
	fmt.Println("RemoteTest")
	sh = shell.NewShell("10.0.1.128:5001")
	f, err := os.Open("test")
	if err != nil {
		fmt.Println(err.Error())
	}
	copyNodes := make([]string, 3)
	copyNodes[0] = "10.0.1.106:5001"
	copyNodes[1] = "10.0.1.107:5001"
	copyNodes[2] = "10.0.1.108:5001"
	hash, err := sh.AddAndCopy(f, 1, copyNodes)
	if err!= nil {
		fmt.Println(err.Error())
	}
	fmt.Println(hash)
}

func cryptTest(){
	fmt.Println("CryptoTest")
	sh = shell.NewShell("10.0.1.128:5001")

	for i := 0; i < 10; i++  {
		data := []byte(randString())
		hash, err := sh.EncryptAndAdd(data, "test11test11test11test11test11", shell.AES)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("File Hash: %s\n", hash)

		decData, err := sh.GetAndDecrypt(hash, "test11test11test11test11test11")
		if err != nil {
			fmt.Println(err.Error())
		}
		if err != nil {
			fmt.Println(err.Error())
		}
		if bytes.Compare(decData, data) == 0 {
			fmt.Println("GetAndDecrypt Success")
		} else {
			fmt.Println("GetAndDecrypt Failed")
		}
	}
}

func randomTest(){
	fmt.Println("RandomTest")
	sh = shell.NewShell("10.0.1.128:5001")

	for i := 0; i < 20; i++ {
		_, err := makeRandomObject()
		if err != nil {
			fmt.Println("err: ", err)
			return
		}
	}
	fmt.Println("we're okay")

	out, err := makeRandomDir(10)
	fmt.Printf("%d calls\n", ncalls)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(out)
	for {
		time.Sleep(time.Second * 1000)
	}

}
func main() {
	localCopyTest()
	//remoteCopyTest()
	//cryptTest()
	//randomTest()
}
