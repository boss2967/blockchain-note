package com.mynote.day01

fun main(args: Array<String>) {
    val str = "zhangsan.lisi-wangwu"
    val result: List<String> = str.split(".", "-")
    println(result)
}