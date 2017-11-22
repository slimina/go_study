package main

import (
	"fmt"
	"os"
)

//文件操作的大多数函数都是在os包里面
/*
目录操作
func Mkdir(name string, perm FileMode) error
创建名称为name的目录，权限设置是perm，例如0777
func MkdirAll(path string, perm FileMode) error
根据path创建多级子目录，例如astaxie/test1/test2。
func Remove(name string) error
删除名称为name的目录，当目录下有文件或者其他目录时会出错
func RemoveAll(path string) error
根据path删除多级子目录，如果path是单个名称，那么该目录下的子目录全部删除。
*/
func main() {
	/*
			第一位 含义：第一位换成二进制也分成三部分，abc
			a - setuid位, 如果该位为1, 则表示设置setuid   它允许用户以可执行文件owner或group的权限来
						运行这个可执行文件。它们经常适用于：为了运行特定的任务，可以允许用户暂时的提高权限。
			b - setgid位, 如果该位为1, 则表示设置setgid  该权限只对目录有效. 目录被设置该位后,
						  任何用户在此目录下创建的文件都具有和该目录所属的组相同的组.
			c - sticky位, 如果该位为1, 则表示设置sticky 对文件使用sticky bit位. 设置该位后,
							就算用户对目录具有写权限, 也不能删除该文件.
		   第2、3、4位  所有者权限，同组用户权限，其它用户权限
		   4代表读权限，2代表写权限，1代表执行权限
	*/
	err := os.Mkdir("test", 0777)
	if err != nil {
		fmt.Println(err)
	}
	defer os.Remove("test")
	os.MkdirAll("ok/a/b", 0777)
	defer os.RemoveAll("ok")
}
