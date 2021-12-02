# Advent of Code 2021 - Day 1
#
# godot --no-window -s day1.gd

extends SceneTree

func read_ints(file: String) -> Array:
    var ints = []
    var f = File.new()
    f.open(file, File.READ)
    if f.is_open():
        while not f.eof_reached():
            ints.append(int(f.get_line()))
    return ints

func count_increments(data: Array) -> int:
    var count = 0
    for i in len(data) - 1:
        if data[i] < data[i + 1]:
            count += 1
    return count

func merge_window(data: Array, width: int) -> Array:
    var merged = []
    for i in len(data) - width + 1:
        var sum = 0
        for j in range(width):
            sum += data[i + j]
        merged.append(sum)
    return merged

func _init() -> void:
    var data_sample: Array = read_ints("./data/sample-1.txt")
    var data_input: Array = read_ints("./data/input-1.txt")

    print("Solution 1 (sample): %d" % count_increments(data_sample))
    print("Solution 1 (input): %d" % count_increments(data_input))

    print("Solution 2 (sample): %d" % count_increments(merge_window(data_sample, 3)))
    print("Solution 2 (input): %d" % count_increments(merge_window(data_input, 3)))

    quit()
