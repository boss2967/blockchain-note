package com.mynote.day01

/*
*
*  val str: String  非空类型 不能赋值为null
*  val str: String？可空类型 可以赋值为null
*  空安全调用符 ：？
*  Elvis 操作符？：
* */
fun main(args: Array<String>) {
    val str: String? = null
//    告诉编译器不用检查了 我一定不为空null，不建议使用
//    str!!.toInt()
//    告诉编译器我可能为空
//    空安全调用符返回值就是可空类型
    str?.toInt()

    val b :Int = str?.toInt()?:-1

    println(b)
}