use std::collections::{HashMap, VecDeque};

fn main() {
    let contents = include_str!("../input.txt");

    let mut instructions: VecDeque<i32> = contents
        .split("\n")
        .map(|line| {
            if line == "noop" {
                return 0;
            }
            return line.split(" ").collect::<Vec<&str>>()[1].parse().unwrap();
        })
        .collect();

    let mut x: i32 = 1;
    let mut current_cycle = 1;
    let mut values_in_cycle = HashMap::new();
    let mut executing_instruction: i32 = 0;
    let mut screen: Vec<&str> = vec![" "; 6 * 40];

    while current_cycle < 241 {
        values_in_cycle.insert(current_cycle, x);

        if executing_instruction == 0 {
            if !instructions.is_empty() {
                let instruction = instructions.pop_front().unwrap();
                if instruction != 0 {
                    executing_instruction = instruction
                }
            }
        } else {
            x += executing_instruction;
            executing_instruction = 0;
        }
        current_cycle += 1;
    }

    let mut part_one_result: i32 = 0;

    for v in [20, 60, 100, 140, 180, 220] {
        let multiplier: i32 = v.try_into().unwrap();
        part_one_result += values_in_cycle.get(&v).unwrap() * multiplier;
    }

    let mut row = 0;
    for x in 1..=240 {
        let m = values_in_cycle.get(&x).unwrap() + (row * 40);
        if [m - 1, m, m + 1].contains(&(&x - &1)) {
            screen[x as usize - 1] = "#";
        }
        if x % 40 == 0 {
            row += 1;
        }
    }

    for (i, p) in screen.iter().enumerate() {
        print!("{p}");
        if [39, 79, 119, 159, 199, 239].contains(&i) {
            print!("\n");
        }
    }

    print!("\n");
    assert_eq!(16020, part_one_result);
    println!("part_one_result: {}", part_one_result);
}
