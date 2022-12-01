use std::fs;

fn main() {
    let contents = fs::read_to_string("input.txt").expect("Something went wrong reading the file");

    let mut result: Vec<i32> = vec![];

    for b in contents.split("\n\n") {
        let mut total_for_elf = 0;
        for number in b.split("\n") {
            total_for_elf += number.parse::<i32>().unwrap();
        }
        result.push(total_for_elf)
    }

    result.sort();
    result.reverse();
    println!("Part One {}", result[0]);
    println!("Part Two {}", result[0] + result[1] + result[2]);
}
