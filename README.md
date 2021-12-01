# Advent of Code 2021

This repository includes my [Go](https://golang.org/) code for the [Advent of Code 2021](https://adventofcode.com/2021/about).

*If you want to solve the puzzles yourself first, do so, **then** look for the solution to compare.* Do **not** cheat yourself out of a great time and learning experience by just copying code!

**Please keep in mind most of this is hacked together. I don't claim it's perfect code or even in all cases 100% valid/defined/proper code or whatever. It's for fun and it got the job done in a short amount of time.** 😀

I'll update this repository now and then to include all my solutions whenever I've got time to properly comment and upload it.

This is mostly for those that are curious. If you want to look at the code, experiment with it, change it, etc. be my guest.

## Running from command line

For this year I decided to try something new and use [Go](https://golang.org/). We'll see how it goes…

To run the code for any given day, use the following line (replacing `1` with the number of the desired day):

```bash
go run ./cmd/day1
```

As an alternative, you can also compile the executables explicitly, using `build` rather than `run`:

```bash
go build ./cmd/day1
./day1
```

Inputs are expected to be in the `data` sub folder relative to the current working directory.

## Running/debugging in VS Code

I'm using [Visual Studio Code](https://code.visualstudio.com/) for writing and debugging. This should work out of the box, if you open the repository directory directly.

You should be able to pick the source code file of any day and simply hit <kbd>F5</kbd> to run/debug it after creating/adjusting your local `launch.json` configuration file to change the working directory (`cwd`) to the workspace folder (`${workspaceFolder}`) as seen in the following example:

```json
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch Package",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${fileDirname}",
            "cwd": "${workspaceFolder}"
        }
    ]
}
```
