fn main() {
    let inp = include_str!("./input.txt");
    let res = process(inp);
    println!("solution: {}", res);
}

fn process(input: &str) -> u32 {
    input
        .lines()
        .map(|line| line.chars().filter(|d| d.is_numeric()))
        .map(|line| line.map(|c| c.to_digit(10).unwrap()).collect::<Vec<u32>>())
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
                "1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet"
            ),
            142
        );
    }
}
