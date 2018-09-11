package com.mynote.day01

fun main(args: Array<String>) {
    val str1 = "abc"
    val str2 = String(charArrayOf('a', 'b', 'c'))
    println(str1.equals(str2))//true
    println(str1 == str2)//true
    println(str1 === str2)//false
}