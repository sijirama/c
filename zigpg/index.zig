const std = @import("std");

pub fn main() !void {
    const gpAllocator = std.heap.GeneralPurposeAllocator(.{}){};
    const allocator = gpAllocator.allocator();

    const arena = std.heap.ArenaAllocator.init(allocator);
    defer arena.deinit();

    const allocator2 = &arena.allocator;

    const varU8 = try allocator2.alloc(u8, 20);
    const tokens = std.ArrayList(u8).init(allocator);
}
