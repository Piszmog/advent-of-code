package io.github.piszmog.aoc

import java.time.Instant

fun main(args: Array<String>) {
    val start = Instant.now()
    val parser = getCSVParser(args[0])
    val depths = parser.map { it[0].toInt() }.toList()

    val part1Solution = day1Part1(depths)
    println("Part 1: $part1Solution")
    val part2Solution = day1Part2(depths)
    println("Part 2: $part2Solution")

    printElapsedTime(start)
}

fun day1Part1(input: List<Int>): Int {
    var increases = 0
    var previous = Int.MAX_VALUE
    input.forEach {
        if (it > previous) {
            increases++
        }
        previous = it
    }
    return increases
}

fun day1Part2(input: List<Int>): Int {
    val sumMeasurements = mutableListOf<Int>()
    var increases = 0
    var previous = Int.MAX_VALUE
    input.forEach {
        if (sumMeasurements.sum() > previous) {
            increases++
            previous = sumMeasurements.sum()
        }
        if (sumMeasurements.size == 3) {
            previous = sumMeasurements.sum()
            sumMeasurements.removeFirst()
        }
        sumMeasurements.add(it)
    }
    if (sumMeasurements.sum() > previous) {
        increases++
    }
    return increases
}