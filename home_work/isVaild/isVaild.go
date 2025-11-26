package main

import "fmt"

func isValidDetailed(s string) bool {
	// 边界情况处理
	if s == "" {
		return true
	}

	if len(s)%2 != 0 {
		return false
	}

	// 定义括号配对关系
	//keytype -> byte
	//valuetype -> byte
	//初始化时，同时插入键值对
	bracketPairs := map[byte]byte{
		')': '(', // 右括号 -> 对应的左括号
		']': '[',
		'}': '{',
	}

	// 使用切片模拟栈
	// 初始化一个空切片， []byte{} 创建的是长度和容量都为0的切片
	stack := []byte{}

	for i := range s {
		//    for i := 0; i < len(s); i++ {
		currentChar := s[i]

		/*
		* value, exists := map[key]
		* value = 如果key存在，返回对应的值；如果key不存在，返回value类型的零值
		* exists = 布尔值，true表示key存在，false表示key不存在
		 */

		// 检查当前字符是否是右括号
		if expectedLeft, isRightBracket := bracketPairs[currentChar]; isRightBracket {
			// 如果是右括号，检查栈顶元素是否匹配

			// 情况1: 栈为空，没有左括号与之匹配
			if len(stack) == 0 {
				return false
			}

			// 情况2: 栈顶元素不匹配
			top := stack[len(stack)-1]
			if top != expectedLeft {
				return false
			}

			// 匹配成功，弹出栈顶元素
			stack = stack[:len(stack)-1]
		} else {
			// 当前字符是左括号，压入栈中
			stack = append(stack, currentChar)
		}
	}

	// 所有字符处理完毕后，栈应该为空
	// 如果栈不为空，说明有未匹配的左括号
	return len(stack) == 0
}

func main() {
	// 测试代码
	s := "()"
	fmt.Println("()字符串是否有效:", isValidDetailed(s))

	s = "()[]{}"
	fmt.Println("()[]{}字符串是否有效:", isValidDetailed(s))

	s = "(]"
	fmt.Println("(]字符串是否有效:", isValidDetailed(s))

	s = "([])"
	fmt.Println("([])字符串是否有效:", isValidDetailed(s))

	s = "([)]"
	fmt.Println("([)]字符串是否有效:", isValidDetailed(s))
}
