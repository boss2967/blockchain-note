package com.mynote.day01

fun main(args: Array<String>) {
    val str="   张三   "
    val newStr:String = str.trim()
    println(newStr)
    val str2="""
        /张三
        /李四
        /王五
    """.trimMargin("/")
    println(str2)
}