use std::collections::VecDeque;

use itertools::Itertools;

fn main() {
    let mut grid = include_str!("../input.txt")
        .lines()
        .map(|line| line.as_bytes().iter().copied().collect::<Vec<_>>())
        .collect::<Vec<_>>();

    let start_coordinates = (0..grid.len())
        .cartesian_product(0..grid[0].len())
        .find(|&(x, y)| grid[x][y] == b'S')
        .unwrap();
    let target_coordinates = (0..grid.len())
        .cartesian_product(0..grid[0].len())
        .find(|&(x, y)| grid[x][y] == b'E')
        .unwrap();

    grid[start_coordinates.0][start_coordinates.1] = b'a';
    grid[target_coordinates.0][target_coordinates.1] = b'z';

    let part_one_result = bfs(&grid, start_coordinates, target_coordinates).unwrap();

    let part_two_result = (0..grid.len())
        .cartesian_product(0..grid[0].len())
        .filter(|&(x, y)| grid[x][y] == b'a')
        .filter_map(|start| bfs(&grid, start, target_coordinates))
        .min()
        .unwrap();

    println!("part_one_result: {}", part_one_result);
    println!("part_two_result: {}", part_two_result);
}

fn bfs(grid: &[Vec<u8>], start: (usize, usize), goal: (usize, usize)) -> Option<usize> {
    let mut visited = vec![vec![false; grid[0].len()]; grid.len()];
    let mut queue = VecDeque::new();
    queue.push_back((start, 0));
    while let Some(((x, y), len)) = queue.pop_front() {
        if (x, y) == goal {
            return Some(len);
        }
        for (dx, dy) in [(0, -1), (-1, 0), (0, 1), (1, 0)] {
            let (nx, ny) = ((x as isize + dx) as usize, (y as isize + dy) as usize);

            let square = match grid.get(nx).and_then(|row| row.get(ny)) {
                Some(&square) => square,
                None => {
                    continue;
                }
            };

            if grid[x][y] + 1 >= square && !visited[nx][ny] {
                visited[nx][ny] = true;
                queue.push_back(((nx, ny), len + 1));
            }
        }
    }
    None
}
