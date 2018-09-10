package com.mynote.day01

fun main(args: Array<String>) {
    fun sayHello(): String {
        return "hello"
    }

    fun create(place :String):String{
        val result = "今天，天气晴朗，我们去$place，首先映入眼帘的是$place${place.length}个鎏金大字 ${sayHello()}"
        return result
    }
    println(create("颐和园"))

}