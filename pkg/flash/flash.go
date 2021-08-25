package flash

import (
	"encoding/gob"
	"github.com/attitudefx7/goblog/pkg/session"
)

type Flashes map[string]interface{}

var flashKey = "_flashes"

func init()  {
	gob.Register(Flashes{})
}

func Info(message string)  {
	addFlash("info", message)
}

// Warning 添加 Warning 类型的消息提示
func Warning(message string) {
	addFlash("warning", message)
}

// Success 添加 Success 类型的消息提示
func Success(message string) {
	addFlash("success", message)
}

// Danger 添加 Danger 类型的消息提示
func Danger(message string) {
	addFlash("danger", message)
}

// All 获取所有消息
func All() Flashes {
	val := session.Get(flashKey)
	// 读取是必须做类型检测
	flashMessages, ok := val.(Flashes)
	if !ok {
		return nil
	}
	// 读取即销毁，直接删除
	session.Forget(flashKey)
	return flashMessages
}

// 私有方法，新增一条提示
func addFlash(key string, message string) {
	flashes := Flashes{}
	flashes[key] = message
	session.Put(flashKey, flashes)
	session.Save()
}
