package main

import (
	"fmt"
	"os"
	"strings"
)

/*
环境变量 是一种向 Unix 程序传递配置信息的常见方式。
让我们来看看如何设置、获取以及列出环境变量。*/

func main() {
	/*使用 os.Setenv 来设置一个键值对。
	使用 os.Getenv获取一个键对应的值。
	如果键不存在，将会返回一个空字符串。*/
	os.Setenv("foo", "1")
	fmt.Println("FOO:", os.Getenv("foo"))
	fmt.Println("BAR:", os.Getenv("bar"))
	/*PS D:\GolandProjects\gobyexample> go run .\environment-variables.go
	FOO: 1
	BAR: */

	/*使用 os.Environ 来列出所有环境变量键值对。
	这个函数会返回一个 KEY=value 形式的字符串切片。
	你可以使用 strings.SplitN 来得到键和值。
	这里我们打印所有的键。*/
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}

	//如果我们在运行前设置了 BAR 的值，那么运行程序将会获取到这个值。 windows下会报错  改用git bash 就能成功执行了
	//BAR=2 go run environment-variables.go
	/*PS D:\GolandProjects\gobyexample> BAR=123 go run .\environment-variables.go
	BAR=123 : 无法将“BAR=123”项识别为 cmdlet、函数、脚本文件或可运行程序的名称。请检查名称的拼写，如果包括路径，请确保路径正确，然后再试一次。
	所在位置 行:1 字符: 1
	+ BAR=123 go run .\environment-variables.go
	+ ~~~~~~~
	    + CategoryInfo          : ObjectNotFound: (BAR=123:String) [], CommandNotFoundException
	    + FullyQualifiedErrorId : CommandNotFoundException

	*/
}

/*foo
TERM_SESSION_ID
UATDATA
ZES_ENABLE_SYSMAN
ProgramW6432
DATAGRIP_VM_OPTIONS
CLION_VM_OPTIONS
JETBRAINS_CLIENT_VM_OPTIONS
CommonProgramW6432
GOPATH
USERNAME
ALLUSERSPROFILE
USERPROFILE
RUBYMINE_VM_OPTIONS
_INTELLIJ_FORCE_SET_GOROOT
PROCESSOR_REVISION
IDEA_INITIAL_DIRECTORY
FPS_BROWSER_APP_PROFILE_STRING
PUBLIC
PYCHARM_VM_OPTIONS
Path
GOROOT
DriverData
JETBRAINSCLIENT_VM_OPTIONS
IDEA_VM_OPTIONS
SESSIONNAME
LOGONSERVER
HOMEDRIVE
APPCODE_VM_OPTIONS
TERMINAL_EMULATOR
HOMEPATH
SystemRoot
STUDIO_VM_OPTIONS
VS120COMNTOOLS
ChocolateyInstall
_INTELLIJ_FORCE_PREPEND_PATH
LOCALAPPDATA
APPDATA
PROCESSOR_IDENTIFIER
GoLand
PATHEXT
PSModulePath
ProgramFiles(x86)
_INTELLIJ_FORCE_SET_GO111MODULE
WEBIDE_VM_OPTIONS
PHPSTORM_VM_OPTIONS
M2_HOME
RIDER_VM_OPTIONS
OS
USERDNSDOMAIN
PROCESSOR_ARCHITECTURE
WEBSTORM_VM_OPTIONS
NUMBER_OF_PROCESSORS
JAVA_HOME
ComSpec
PROCESSOR_LEVEL
DEVECOSTUDIO_VM_OPTIONS
USERDOMAIN_ROAMINGPROFILE
PyCharm Community Edition
ProgramFiles
windir
DATASPELL_VM_OPTIONS
_INTELLIJ_FORCE_SET_GOPATH
TEMP
TMP
GATEWAY_VM_OPTIONS
OneDrive
CommonProgramFiles(x86)
USERDOMAIN
SystemDrive
GO111MODULE
COMPUTERNAME
ProgramData
GOLAND_VM_OPTIONS
ChocolateyLastPathUpdate
FPS_BROWSER_USER_PROFILE_STRING

CommonProgramFiles
SSLKEYLOGFILE
__INTELLIJ_COMMAND_HISTFILE__
*/
