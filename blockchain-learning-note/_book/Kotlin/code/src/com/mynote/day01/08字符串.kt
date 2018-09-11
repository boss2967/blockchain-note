package com.mynote.day01

fun main(args: Array<String>) {
    val place = "广东省深圳市保安区"
    println(place)
    val place2 = "广东省\n深圳市\n保安区"
    println(place2)
//  原样输出字符串
    val place3 = """
        广东省
        深圳市
        保安区
""".trimIndent()
    println(place3)
}