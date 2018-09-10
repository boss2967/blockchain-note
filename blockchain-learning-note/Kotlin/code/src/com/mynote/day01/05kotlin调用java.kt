package com.mynote.day01

import java.math.BigDecimal

fun main(args: Array<String>) {
    //存钱
    var money:Double  = 1.234567898465423454365254365
    //取钱
    println(money)//1.2345678984654234 有误差
    var bigDe:BigDecimal = BigDecimal(1.234567898465423454365254365)//省略new关键字
    println(bigDe)
}