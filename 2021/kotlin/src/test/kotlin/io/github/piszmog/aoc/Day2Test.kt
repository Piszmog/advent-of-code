package io.github.piszmog.aoc

import org.junit.jupiter.api.Test
import kotlin.test.assertEquals

class Day2Test {
    @Test
    fun `Day 2 Part 1`() {
        val actual = day2Part1(getInput())
        assertEquals(150, actual)
    }

    @Test
    fun `Day 2 Part 2`() {
        val actual = day2Part2(getInput())
        assertEquals(900, actual)
    }

    private fun getInput() =
        listOf(
            "forward 5",
            "down 5",
            "forward 8",
            "up 3",
            "down 8",
            "forward 2",
        )
            .map { it.split(" ") }
            .map { Course(Direction.valueOf(it[0].uppercase()), it[1].toInt()) }
            .toList()
}