const std = @import("std");

pub fn main() !void {
    const stdout = std.io.getStdOut().writer();
    const allocator = std.heap.page_allocator;

    const file = try std.fs.cwd().openFile("input.txt", .{});
    defer file.close();

    var buf = std.io.bufferedReader(file.reader());

    while (true) {
        var line = try buf.readUntilDelimiterOrEofAlloc(allocator, '\n');
        if (line == null) break;
        try stdout.print("{}\n", .{line});
        allocator.free(line);
    }
}
