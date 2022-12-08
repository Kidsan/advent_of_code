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

    let position_x = 1;
    let position_y = 1;

    let limit_x = width - 1;
    let limit_y = height - 1;

    let mut scores: Vec<i32> = Vec::new();

    for x in position_x..limit_x {
        for y in position_y..limit_y {
            let current = grid[x][y];
            if search(&grid, current, x, y) {
                visible += 1;
            }
            scores.push(get_view_score(&grid, current, x, y));
        }
    }

    scores.sort();

    println!("Part One: {}", visible);
    println!("Part Two: {}", scores.last().unwrap());
}

fn search(grid: &Vec<Vec<usize>>, current: usize, x: usize, y: usize) -> bool {
    if search_up(grid, current, x, y)
        || search_left(grid, current, x, y)
        || search_down(grid, current, x, y)
        || search_right(grid, current, x, y)
    {
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

fn get_view_score(grid: &Vec<Vec<usize>>, current: usize, x: usize, y: usize) -> i32 {
    let score_up = visible_trees_north(grid, &0, current, x, y);
    let score_left = visible_trees_west(grid, &0, current, x, y);
    let score_down = visible_trees_south(grid, &0, current, x, y);
    let score_right = visible_trees_east(grid, &0, current, x, y);

    return score_down * score_left * score_right * score_up;
}

fn visible_trees_north(
    grid: &Vec<Vec<usize>>,
    count: &i32,
    current: usize,
    x: usize,
    y: usize,
) -> i32 {
    if y == 0 && current > grid[x][y] {
        return *count;
    }
    if current > grid[x][y - 1] {
        return visible_trees_north(grid, &(count + 1), current, x, y - 1);
    }
    *count + 1
}

fn visible_trees_west(
    grid: &Vec<Vec<usize>>,
    count: &i32,
    current: usize,
    x: usize,
    y: usize,
) -> i32 {
    if x == 0 && current > grid[x][y] {
        return *count;
    }
    if current > grid[x - 1][y] {
        return visible_trees_west(grid, &(count + 1), current, x - 1, y);
    }
    *count + 1
}

fn visible_trees_south(
    grid: &Vec<Vec<usize>>,
    count: &i32,
    current: usize,
    x: usize,
    y: usize,
) -> i32 {
    if y == grid[0].len() - 1 && current > grid[x][y] {
        return *count;
    }
    if current > grid[x][y + 1] {
        return visible_trees_south(grid, &(count + 1), current, x, y + 1);
    }
    *count + 1
}

fn visible_trees_east(
    grid: &Vec<Vec<usize>>,
    count: &i32,
    current: usize,
    x: usize,
    y: usize,
) -> i32 {
    if x == grid.len() - 1 && current > grid[x][y] {
        return *count;
    }
    if current > grid[x + 1][y] {
        return visible_trees_east(grid, &(count + 1), current, x + 1, y);
    }
    *count + 1
}
