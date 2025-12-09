package io.github.piszmog.aoc

import java.time.Instant

fun main(args: Array<String>) {
    val start = Instant.now()

    val parser = getCSVParser(args[0], " ")
    val courses = parser.map { Course(Direction.valueOf(it[0].uppercase()), it[1].toInt()) }.toList()

    val part1Solution = day2Part1(courses)
    println("Part 1: $part1Solution")
    val part2Solution = day2Part2(courses)
    println("Part 2: $part2Solution")

    printElapsedTime(start)
}

fun day2Part1(input: List<Course>): Int {
    var horizontalSum = 0
    var depthSum = 0
    input.forEach {
        when (it.direction) {
            Direction.FORWARD -> horizontalSum += it.value
            Direction.DOWN -> depthSum += it.value
            Direction.UP -> depthSum -= it.value
        }
    }
    return horizontalSum * depthSum
}

fun day2Part2(input: List<Course>): Int {
    var horizontalSum = 0
    var depthSum = 0
    var aimSum = 0
    input.forEach {
        when (it.direction) {
            Direction.FORWARD -> {
                horizontalSum += it.value
                depthSum += it.value * aimSum
            }
            Direction.DOWN -> aimSum += it.value
            Direction.UP -> aimSum -= it.value
        }
    }
    return horizontalSum * depthSum
}

data class Course(val direction: Direction, val value: Int)

enum class Direction {
    DOWN, FORWARD, UP
}