package where

type Build struct {
}

// Hold 保存已经注入的Where条件
type Hold struct {
	Where string
	Args  interface{}
}
