use std::i32;

fn main() {
    let contents = include_str!("../input.txt");

    let grid = contents
        .split("\n")
        .map(|line| {
            line.chars()
                .map(|num| num.to_string().parse::<usize>().unwrap())
                .collect::<Vec<usize>>()
        })
        .collect::<Vec<Vec<usize>>>();

    let width = grid.len();
    let height = grid[0].len();

    let mut visible = (width * 2) + (height * 2) - 4;
    println!("{visible}");
    let X = 1;
    let Y = 1;

    let limit_x = width - 1;
    let limit_y = height - 1;

    println!("limits: {limit_x}, {limit_y}");

    for x in X..limit_x {
        for y in Y..limit_y {
            let current = grid[x][y];
            if search(&grid, current, x, y) {
                visible += 1;
            }
        }
    }

    // while X < limit_x && Y < limit_y {
    //     let current = grid[X][Y];
    //     println!("pos: {X}, {Y}");
    //     if search(&grid, current, X, Y) {
    //         visible += 1;
    //     }
    //     if X <= limit_x {
    //         println!("iter");
    //         X += 1;
    //     }

    //     if Y <= limit_y {
    //         println!("iter");
    //         Y += 1;
    //     }
    // }

    println!("Part One: {}", visible);
    // println!("Part Two: {}", p2);
}

fn search(grid: &Vec<Vec<usize>>, current: usize, x: usize, y: usize) -> bool {
    if search_up(grid, current, x, y)
        || search_left(grid, current, x, y)
        || search_down(grid, current, x, y)
        || search_right(grid, current, x, y)
    {
        println!("{x},{y}");
        return true;
    }
    return false;
}

fn search_up(grid: &Vec<Vec<usize>>, current: usize, x: usize, y: usize) -> bool {
    if y == 0 && current > grid[x][y] {
        return true;
    }
    if current > grid[x][y - 1] {
        return search_up(grid, current, x, y - 1);
    }
    false
}

fn search_left(grid: &Vec<Vec<usize>>, current: usize, x: usize, y: usize) -> bool {
    if x == 0 && current > grid[x][y] {
        return true;
    }
    if current > grid[x - 1][y] {
        return search_left(grid, current, x - 1, y);
    }
    false
}

fn search_down(grid: &Vec<Vec<usize>>, current: usize, x: usize, y: usize) -> bool {
    if y == grid[0].len() - 1 && current > grid[x][y] {
        return true;
    }
    if current > grid[x][y + 1] {
        return search_down(grid, current, x, y + 1);
    }
    false
}

fn search_right(grid: &Vec<Vec<usize>>, current: usize, x: usize, y: usize) -> bool {
    if x == grid.len() - 1 && current > grid[x][y] {
        return true;
    }
    if current > grid[x + 1][y] {
        return search_right(grid, current, x + 1, y);
    }
    false
}

#[cfg(test)]
mod tests {
    use super::*;

    // #[test]
    // fn test_parse_command() {
    //     assert!(parse_command("$ ls").is_ok());
    // }
}
