package io.github.piszmog.aoc

import org.junit.jupiter.api.Test
import kotlin.test.assertEquals

class Day5Test {
    @Test
    fun `Day 5 Part 1`() {
        val input = getInput()
        assertEquals(5, getNumOverlap(input, false))
    }

    @Test
    fun `Day 5 Part 2`() {
        val input = getInput()
        assertEquals(12, getNumOverlap(input, true))
    }

    private fun getInput() =
        listOf(
            "0,9 -> 5,9",
            "8,0 -> 0,8",
            "9,4 -> 3,4",
            "2,2 -> 2,1",
            "7,0 -> 7,4",
            "6,4 -> 2,0",
            "0,9 -> 2,9",
            "3,4 -> 1,4",
            "0,0 -> 8,8",
            "5,5 -> 8,2",
        ).map { s ->
            val points = s.split("->").map { it.trim() }
            val start = points[0].split(",").map { it.trim().toInt() }
            val end = points[1].split(",").map { it.trim().toInt() }
            Line(Point(start[0], start[1]), Point(end[0], end[1]))
        }
}