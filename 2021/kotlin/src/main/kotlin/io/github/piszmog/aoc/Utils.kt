package io.github.piszmog.aoc

import org.apache.commons.csv.CSVFormat
import org.apache.commons.csv.CSVParser
import java.io.BufferedReader
import java.io.FileReader
import java.time.Duration
import java.time.Instant

fun printElapsedTime(start: Instant) {
    println("Duration: ${Duration.between(start, Instant.now()).toMillis()}ms")
}

fun getCSVParser(filePath: String, delimiter: String = ","): CSVParser {
    val bufferedReader = BufferedReader(FileReader(filePath))
    val csvFormat = CSVFormat.Builder.create(CSVFormat.DEFAULT).setDelimiter(delimiter[0]).build()
    return CSVParser(bufferedReader, csvFormat)
}

fun getFileReader(filePath: String): BufferedReader {
    return BufferedReader(FileReader(filePath))
}