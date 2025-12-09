package io.github.piszmog.aoc

import kotlin.test.Test
import kotlin.test.assertEquals

class Day6Test {
    @Test
    fun `Day 6 Part 1`() {
        val input = getInput()
        val actual = getFishPop(input, 80)
        assertEquals(5934, actual)
    }

    @Test
    fun `Day 6 Part 2`() {
        val input = getInput()
        val actual = getFishPop(input, 256)
        assertEquals(26984457539, actual)
    }

    private fun getInput() = getFish("3,4,3,1,2")
}