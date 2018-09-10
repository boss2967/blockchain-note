package com.mynote.day01


fun main(args: Array<String>) {

    val a = 10
    val b = 20
    Max(a, b)
}
//kotlin没有三元运算符
//kotlin if when 都有返回值，函数式编程闭包的思想
fun Max(a: Int, b: Int): Int {
    return if (a > b) a else b
}
