fn main() {
    println!("{}", process(include_str!("./input.txt")))
}

struct Game {
    points: i32,
}

impl Game {
    fn new(data: &str) -> Self {
        let parts: Vec<&str> = data.split(':').collect();
        let separated_numbers: Vec<&str> = parts[1].split('|').collect();
        Game {
            points: calculate_score(separated_numbers),
        }
    }
}

fn calculate_score(input: Vec<&str>) -> i32 {
    let left: Vec<i32> = input[0]
        .trim()
        .replace("  ", " ")
        .split(' ')
        .map(|v| v.parse::<i32>().unwrap())
        .collect();
    let right: Vec<i32> = input[1]
        .trim()
        .replace("  ", " ")
        .split(' ')
        .map(|v| v.parse::<i32>().unwrap())
        .collect();

    let mut score = 0;
    let mut count = 0;
    for v in right.iter() {
        if left.contains(v) {
            if count == 0 {
                count = 1;
                score = 1;
            } else {
                score *= 2;
            }
        }
    }
    score
}

fn process(input: &str) -> i32 {
    input.lines().map(Game::new).map(|g| g.points).sum()
}

#[cfg(test)]
mod test {

    use super::*;

    #[test]
    fn test_calculate_score() {
        assert_eq!(
            calculate_score(vec!["41 48 83 86 17", "83 86  6 31 17  9 48 53"]),
            8
        );
        assert_eq!(
            calculate_score(vec!["13 32 20 16 61", "61 30 68 82 17 32 24 1"]),
            2
        );
        assert_eq!(
            calculate_score(vec!["1 21 53 59 44", "69 82 63 72 16 21 14  1"]),
            2
        );
    }

    #[test]
    fn test_process() {
        assert_eq!(
            process(
                "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11"
            ),
            13
        )
    }
}
