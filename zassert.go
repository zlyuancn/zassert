/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2019/5/1
   Description :
-------------------------------------------------
*/

package zassert

import "fmt"

type AssertError struct {
    Msg string
}

func (self AssertError) Error() string {
    return self.Msg
}

func PanicAssertError(msg string) {
    panic(AssertError{msg})
}

func True(a bool, msg ...interface{}) {
    if !a {
        raise("assert.True", msg)
    }
}

func False(a bool, msg ...interface{}) {
    if a {
        raise("assert.False", msg)
    }
}

func Equal(a, b interface{}, msg ...interface{}) {
    if a != b {
        raise("assert.Equal", msg)
    }
}

func NotEqual(a, b interface{}, msg ...interface{}) {
    if a == b {
        raise("assert.NotEqual", msg)
    }
}

func Nil(a interface{}, msg ...interface{}) {
    if a != nil {
        raise("assert.Nil", msg)
    }
}

func NotNil(a interface{}, msg ...interface{}) {
    if a == nil {
        raise("assert.NotNil", msg)
    }
}

func Zero(a int, msg ...interface{}) {
    if a != 0 {
        raise("assert.Zero", msg)
    }
}

func NotZero(a int, msg ...interface{}) {
    if a == 0 {
        raise("assert.NotZero", msg)
    }
}

func EmptyString(a string, msg ...interface{}) {
    if len(a) != 0 {
        raise("assert.EmptyString", msg)
    }
}

func NotEmptyString(a string, msg ...interface{}) {
    if len(a) == 0 {
        raise("assert.NotEmptyString", msg)
    }
}

func raise(def string, msg []interface{}) {
    if msg == nil || len(msg) == 0 {
        PanicAssertError(def)
    } else if len(msg) == 1 {
        PanicAssertError(msg[0].(string))
    } else {
        PanicAssertError(fmt.Sprintf(msg[0].(string), msg[1:]...))
    }
}
