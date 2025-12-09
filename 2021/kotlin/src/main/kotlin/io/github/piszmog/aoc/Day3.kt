package io.github.piszmog.aoc

import java.time.Instant
import kotlin.math.pow

fun main(args: Array<String>) {
    val start = Instant.now()
    val parser = getCSVParser(args[0])
    val report = parser.map { it -> it[0].toCharArray().map { it == '1' } }.toList()

    val part1Solution = day3Part1(report)
    println("Part 1: $part1Solution")
    val part2Solution = day3Part2(report)
    println("Part 2: $part2Solution")

    printElapsedTime(start)
}

fun day3Part1(input: List<List<Boolean>>): Int {
    var gammaRate = IntArray(input[0].size) { 0 }
    input.forEach { row ->
        var col = 0
        row.forEach {
            when (it) {
                true -> gammaRate[col]++
                false -> gammaRate[col]--
            }
            col++
        }
    }
    gammaRate = gammaRate.map { if (it > 0) 1 else 0 }.toIntArray()
    val epsilonRate = gammaRate.map { if (it == 1) 0 else 1 }.toIntArray()
    return binaryToDecimal(gammaRate) * binaryToDecimal(epsilonRate)
}

fun day3Part2(input: List<List<Boolean>>) =
    getRating(input, true) * getRating(input, false)

private fun getRating(input: List<List<Boolean>>, bitType: Boolean): Int {
    var data = input
    var binaryBool = listOf<Boolean>()
    var mostCommon = 0
    for (col in 0..input[0].size) {
        if (data.size == 1) {
            binaryBool = data.first()
            break
        }
        if (data.size == 2) {
            binaryBool = data.first { it[col] == bitType }
            break
        }
        data.forEach { row ->
            when (row[col]) {
                true -> mostCommon++
                false -> mostCommon--
            }
        }
        val mostCommonBit = when (bitType) {
            true -> if (mostCommon > 0) true else mostCommon >= 0
            false -> if (mostCommon > 0) false else mostCommon < 0
        }
        mostCommon = 0
        data = data.filter { it[col] == mostCommonBit }.toList()
    }
    val binary = binaryBool.map { if (it) 1 else 0 }.toIntArray()
    return binaryToDecimal(binary)
}

private fun binaryToDecimal(input: IntArray): Int {
    var result = 0
    for ((power, i) in input.reversed().withIndex()) {
        result += i * 2.0.pow(power.toDouble()).toInt()
    }
    return result
}
