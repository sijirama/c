const std = @import("std");

pub fn main() !void {
    const val_1: i32 = 50;
    const val_2: f64 = 3.14;

    const result = val_1 + val_2;
    std.debug.print("{}\n", .{result});
}
