package io.github.piszmog.aoc

import java.time.Instant

fun main(args: Array<String>) {
    val start = Instant.now()
    val reader = getFileReader(args[0])
    val lines = reader.lines().map { s ->
        val points = s.split("->").map { it.trim() }
        val startPoint = points[0].split(",").map { it.trim().toInt() }
        val end = points[1].split(",").map { it.trim().toInt() }
        Line(Point(startPoint[0], startPoint[1]), Point(end[0], end[1]))
    }.toList()

    val part1Solution = getNumOverlap(lines, false)
    println("Part 1: $part1Solution")
    val part2Solution = getNumOverlap(lines, true)
    println("Part 2: $part2Solution")
    printElapsedTime(start)
}

fun getNumOverlap(input: List<Line>, diagonal: Boolean): Int {
    val map = mutableMapOf<Point, Int>()
    var lines = input
    if (!diagonal) {
        lines = lines.filter { !it.isDiagonal() }
    }
    lines.forEach { line ->
        line.getPath().forEach { point ->
            if (map.containsKey(point)) {
                map[point] = map[point]!! + 1
            } else {
                map[point] = 1
            }
        }
    }
    return map.filter { it.value > 1 }.size
}

data class Line(val start: Point, val end: Point) {
    fun getPath(): List<Point> {
        val path = mutableListOf<Point>()
        val current = start.copy()
        while (current != end) {
            path.add(current.copy())
            if (current.x > end.x) {
                current.x--
            } else if (current.x < end.x) {
                current.x++
            }
            if (current.y > end.y) {
                current.y--
            } else if (current.y < end.y) {
                current.y++
            }
        }
        path.add(end)
        return path
    }

    fun isDiagonal() = start.x != end.x && start.y != end.y

}

data class Point(var x: Int, var y: Int)