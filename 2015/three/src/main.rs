use ::std::fs;
use std::collections::HashMap;

fn main() {
    println!("Hello, world!");
    let input: Vec<String> = read("input.txt");
    println!("Part One: {}", at_least_one_gift(&input));
    println!("Part Two: {}", two_santas(&input));
}

fn read(filename: &str) -> Vec<String> {
    let contents = fs::read_to_string(filename).expect("Something went wrong reading the file");

    let mut result: Vec<String> = vec![];

    for b in contents.split("") {
        result.push(b.to_string())
    }
    result
}

fn at_least_one_gift(input: &Vec<String>) -> i32 {
    let mut visited = HashMap::new();

    let mut x_position = 0;
    let mut y_position = 0;

    visited.insert(
        x_position.to_string() + &",".to_string() + &y_position.to_string(),
        1,
    );

    for direction in input {
        match direction.as_str() {
            "^" => y_position += 1,
            ">" => x_position += 1,
            "v" => y_position -= 1,
            "<" => x_position -= 1,
            _ => {}
        }
        visited.insert(
            x_position.to_string() + &",".to_string() + &y_position.to_string(),
            1,
        );
    }
    visited.len().try_into().unwrap()
}

fn two_santas(input: &Vec<String>) -> i32 {
    let mut visited = HashMap::new();

    let mut x_position = 0;
    let mut y_position = 0;

    let mut r_x_position = 0;
    let mut r_y_position = 0;

    visited.insert(
        x_position.to_string() + &",".to_string() + &y_position.to_string(),
        1,
    );

    for direction in input
    .into_iter()
    .enumerate()
    {
        // TODO: find how to loop and take 2 or some smarter way than this.
        if direction.0 % 2 == 0 {
            match direction.1.as_str() {
                "^" => y_position += 1,
                ">" => x_position += 1,
                "v" => y_position -= 1,
                "<" => x_position -= 1,
                _ => {}
            }
            visited.insert(
                x_position.to_string() + &",".to_string() + &y_position.to_string(),
                1,
            );
        } else {
            match direction.1.as_str() {
                "^" => r_y_position += 1,
                ">" => r_x_position += 1,
                "v" => r_y_position -= 1,
                "<" => r_x_position -= 1,
                _ => {}
            }
            visited.insert(
                r_x_position.to_string() + &",".to_string() + &r_y_position.to_string(),
                1,
            );
        }

    }
    visited.len().try_into().unwrap()
}
