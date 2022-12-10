use std::{collections::HashMap, i32};

fn main() {
    let contents = include_str!("../input.txt");

    let instructions: Vec<(&str, i32)> = contents
        .split("\n")
        .map(|line| line.split(" ").collect::<Vec<&str>>())
        .map(|parts| (parts[0], parts[1].parse().unwrap()))
        .collect();

    let mut rope: Vec<(i32, i32)> = Vec::new();

    for _ in 0..10 {
        rope.push((50, 50));
    }

    let (part_one_result, part_two_result) = track_rope(instructions.clone(), &mut rope);

    println!("part_one_result: {}", part_one_result);
    println!("part_one_result: {}", part_two_result);

    assert_eq!(6266, part_one_result);
    assert_eq!(2369, part_two_result);
}

fn track_rope(instructions: Vec<(&str, i32)>, rope: &mut Vec<(i32, i32)>) -> (i32, i32) {
    assert!(rope.len() >= 2);
    let mut visited_by_second_knot: HashMap<_, _> = HashMap::<(i32, i32), i32>::new();
    let mut visited_by_tail: HashMap<_, _> = HashMap::<(i32, i32), i32>::new();

    for (direction, distance) in instructions {
        match direction {
            "R" => {
                for _ in 0..distance {
                    rope[0].0 += 1;
                    for n in 0..rope.len() - 1 {
                        update_knots(rope[n], &mut rope[n + 1]);
                    }
                    visited_by_second_knot.insert(rope[1], 1);
                    visited_by_tail.insert(rope[rope.len() - 1], 1);
                }
            }
            "L" => {
                for _ in 0..distance {
                    rope[0].0 -= 1;
                    for n in 0..rope.len() - 1 {
                        update_knots(rope[n], &mut rope[n + 1]);
                    }
                    visited_by_second_knot.insert(rope[1], 1);
                    visited_by_tail.insert(rope[rope.len() - 1], 1);
                }
            }
            "D" => {
                for _ in 0..distance {
                    rope[0].1 += 1;
                    for n in 0..rope.len() - 1 {
                        update_knots(rope[n], &mut rope[n + 1]);
                    }
                    visited_by_second_knot.insert(rope[1], 1);
                    visited_by_tail.insert(rope[rope.len() - 1], 1);
                }
            }
            "U" => {
                for _ in 0..distance {
                    rope[0].1 -= 1;
                    for n in 0..rope.len() - 1 {
                        update_knots(rope[n], &mut rope[n + 1]);
                    }
                    visited_by_second_knot.insert(rope[1], 1);
                    visited_by_tail.insert(rope[rope.len() - 1], 1);
                }
            }
            _ => {}
        }
    }
    return (
        visited_by_second_knot.len().try_into().unwrap(),
        visited_by_tail.len().try_into().unwrap(),
    );
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
