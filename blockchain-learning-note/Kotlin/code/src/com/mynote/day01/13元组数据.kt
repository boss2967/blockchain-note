package com.mynote.day01

fun main(args: Array<String>) {
//    二元元组
//    val pair=Pair<String ,Int>("张三",20)
    var pair = "张三" to 20
    println(pair.first)
    println(pair.second)
//    三元元组
    val triple = Triple("李四",2,"12323")
    println(triple)
}

