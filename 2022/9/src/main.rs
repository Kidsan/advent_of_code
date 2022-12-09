use std::{collections::HashMap, i32};

fn main() {
    let contents = include_str!("../input.txt");

    let mut grid: Vec<_> = Vec::<Vec<i32>>::new();

    for _ in 0..9 {
        grid.push(Vec::from([0, 0, 0, 0, 0, 0, 0, 0, 0, 0]));
    }

    let instructions: Vec<(&str, i32)> = contents
        .split("\n")
        .map(|line| line.split(" ").collect::<Vec<&str>>())
        .map(|parts| (parts[0], parts[1].parse().unwrap()))
        .collect();

    let mut part_one_rope = Vec::from([(0, 9), (0, 9)]);

    let mut part_two_rope: Vec<(i32, i32)> = Vec::new();

    for _ in 0..10 {
        part_two_rope.push((50, 50));
    }

    let part_one_result = track_rope(instructions.clone(), &mut part_one_rope);
    let part_two_result = track_rope(instructions, &mut part_two_rope);
    println!("part_one_result: {}", part_one_result);
    println!("part_one_result: {}", part_two_result);
}

fn track_rope(instructions: Vec<(&str, i32)>, rope: &mut Vec<(i32, i32)>) -> i32 {
    let mut visited_by_tail: HashMap<_, _> = HashMap::<(i32, i32), i32>::new();

    for (direction, distance) in instructions {
        match direction {
            "R" => {
                for _ in 0..distance {
                    rope[0].0 += 1;
                    for n in 0..rope.len() - 1 {
                        update_knots(rope[n], &mut rope[n + 1]);
                    }

                    visited_by_tail.insert(rope[rope.len() - 1], 1);
                }
            }
            "L" => {
                for _ in 0..distance {
                    rope[0].0 -= 1;
                    for n in 0..rope.len() - 1 {
                        update_knots(rope[n], &mut rope[n + 1]);
                    }

                    visited_by_tail.insert(rope[rope.len() - 1], 1);
                }
            }
            "D" => {
                for _ in 0..distance {
                    rope[0].1 += 1;
                    for n in 0..rope.len() - 1 {
                        update_knots(rope[n], &mut rope[n + 1]);
                    }

                    visited_by_tail.insert(rope[rope.len() - 1], 1);
                }
            }
            "U" => {
                for _ in 0..distance {
                    rope[0].1 -= 1;
                    for n in 0..rope.len() - 1 {
                        update_knots(rope[n], &mut rope[n + 1]);
                    }

                    visited_by_tail.insert(rope[rope.len() - 1], 1);
                }
            }
            _ => {}
        }
    }
    visited_by_tail.len().try_into().unwrap()
}

fn update_knots(knot: (i32, i32), mut next_knot: &mut (i32, i32)) {
    if knot.0 - next_knot.0 > 1 {
        if knot.1 > next_knot.1 {
            next_knot.1 += 1;
        } else if knot.1 < next_knot.1 {
            next_knot.1 -= 1;
        }
        next_knot.0 += 1;
    } else if next_knot.0 - knot.0 > 1 {
        if knot.1 > next_knot.1 {
            next_knot.1 += 1;
        } else if knot.1 < next_knot.1 {
            next_knot.1 -= 1;
        }

        next_knot.0 -= 1;
    } else if knot.1 - next_knot.1 > 1 {
        if next_knot.0 > knot.0 {
            next_knot.0 -= 1;
        } else if knot.0 > next_knot.0 {
            next_knot.0 += 1;
        }

        next_knot.1 += 1;
    } else if next_knot.1 - knot.1 > 1 {
        if knot.0 > next_knot.0 {
            next_knot.0 += 1;
        } else if next_knot.0 > knot.0 {
            next_knot.0 -= 1;
        }
        next_knot.1 -= 1;
    }
}
