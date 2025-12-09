package io.github.piszmog.aoc

import kotlin.test.Test
import kotlin.test.assertEquals

class Day1Test {
    @Test
    fun `Day 1 Part 1`() {
        val actual = day1Part1(getInput())
        assertEquals(7, actual)
    }

    @Test
    fun `Day 1 Part 2`() {
        val actual = day1Part2(getInput())
        assertEquals(5, actual)
    }

    private fun getInput() =
        listOf(199, 200, 208, 210, 200, 207, 240, 269, 260, 263)
}
