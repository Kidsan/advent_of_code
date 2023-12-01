fn main() {
    let inp = include_str!("./input.txt");
    let res = process(inp);
    println!("solution: {}", res);
}

fn process(input: &str) -> u32 {
    input
        .lines()
        .map(|line| {
            line.replace("one", "o1e")
                .replace("two", "t2o")
                .replace("three", "t3e")
                .replace("four", "f4r")
                .replace("five", "f5e")
                .replace("six", "s6x")
                .replace("seven", "s7n")
                .replace("eight", "e8t")
                .replace("nine", "n9e")
        })
        .map(|line| {
            line.as_str()
                .chars()
                .filter(|d| d.is_numeric())
                .collect::<Vec<char>>()
        })
        .map(|line| {
            line.iter()
                .map(|c| c.to_digit(10).unwrap())
                .collect::<Vec<u32>>()
        })
        .map(|line| line.first().unwrap() * 10 + line.last().unwrap())
        .sum()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_process() {
        assert_eq!(
            process(
                "two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen"
            ),
            281
        );
    }
}
