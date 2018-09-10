package com.mynote.day01

fun main(args: Array<String>) {
    val path = "/user/dic/kotilin/haha.doc"
    //获取前6个
    println(path.substring(0..5))
    println(path.substringBefore("r"))
    println(path.substringAfterLast("r"))
    println(path.substringBeforeLast("r"))
    println(path.substringAfter("f"))
}