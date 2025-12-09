package io.github.piszmog.aoc

import java.time.Instant

fun main(args: Array<String>) {
    val start = Instant.now()
    val reader = getFileReader(args[0])
    var row = 0
    var numLines = ""
    val boardLines = mutableListOf<String>()
    reader.lines().forEach {
        if (row == 0) {
            numLines = it
        } else {
            boardLines.add(it)
        }
        row++
    }
    val nums = getNums(numLines)
    val boards = getBoards(boardLines)

    val part1Solution = day4Part1(nums, boards.toMutableList())
    println("Part 1: $part1Solution")
    val part2Solution = day4Part2(nums, boards.toMutableList())
    println("Part 2: $part2Solution")
    printElapsedTime(start)
}

fun getNums(input: String): List<Int> {
    return input.split(",").map { it.toInt() }
}

fun getBoards(lines: List<String>): List<Board> {
    val boards = mutableListOf<Board>()
    val input = mutableListOf<List<Int>>()
    lines.forEach { line ->
        if (line.isBlank()) {
            if (input.isNotEmpty()) {
                boards.add(Board(input.map { it -> Row(it.map { Square(false, it) }) }, false))
                input.clear()
            }
        } else {
            input.add(line.split(" ").filter { it.isNotBlank() }.map { it.toInt() })
        }
    }
    if (input.isNotEmpty()) {
        boards.add(Board(input.map { it -> Row(it.map { Square(false, it) }) }, false))
    }
    return boards
}

fun day4Part1(callNums: List<Int>, boards: List<Board>): Int {
    var finalNum = 0
    var unmarkedSum = 0
    outer@ for (num in callNums) {
        finalNum = num
        for (board in boards) {
            board.markSquare(num)
            if (board.isWinner()) {
                unmarkedSum = board.getUnmarkedSum()
                break@outer
            }
        }
    }
    return finalNum * unmarkedSum
}

fun day4Part2(callNums: List<Int>, boards: List<Board>): Int {
    var finalNum = 0
    var unmarkedSum = 0
    outer@ for (num in callNums) {
        finalNum = num
        for (board in boards) {
            if (board.winner) {
                continue
            }
            board.markSquare(num)
            if (board.isWinner()) {
                board.winner = true
            }
            if (boards.all { it.winner }) {
                unmarkedSum = board.getUnmarkedSum()
                break@outer
            }
        }
    }
    return finalNum * unmarkedSum
}

data class Board(val rows: List<Row>, var winner: Boolean) {
    fun markSquare(num: Int) {
        for (row in rows) {
            for (square in row.squares) {
                if (square.value == num) {
                    square.marked = true
                    return
                }
            }
        }
    }

    fun getUnmarkedSum(): Int {
        var sum = 0
        rows.forEach { row -> row.squares.filter { !it.marked }.map { it.value }.forEach { sum += it } }
        return sum
    }

    fun isWinner(): Boolean {
        for (x in rows.indices) {
            if (isRowWinner(x) || isColumnWinner(x)) {
                return true
            }
        }
        return false
    }

    private fun isRowWinner(row: Int) = rows[row].squares.all { it.marked }

    private fun isColumnWinner(col: Int) = rows.map { it.squares[col].marked }.all { it }
}

data class Row(val squares: List<Square>)

data class Square(var marked: Boolean, val value: Int)