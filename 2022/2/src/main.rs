fn main() {
    let contents = include_str!("../input.txt").lines();

    let result_one: i32 = contents
        .clone()
        .map(|positions| match positions {
            "A X" => 4,
            "B X" => 1,
            "C X" => 7,
            "A Y" => 8,
            "B Y" => 5,
            "C Y" => 2,
            "A Z" => 3,
            "B Z" => 9,
            "C Z" => 6,
            _ => 0,
        })
        .sum();

    let result_two: i32 = contents
        .clone()
        .map(|positions| match positions {
            "A X" => 3,
            "B X" => 1,
            "C X" => 2,
            "A Y" => 4,
            "B Y" => 5,
            "C Y" => 6,
            "A Z" => 8,
            "B Z" => 9,
            "C Z" => 7,
            _ => 0,
        })
        .sum();

    println!("Part One {}", result_one);
    println!("Part Two {}", result_two);
}
