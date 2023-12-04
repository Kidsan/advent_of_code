fn main() {
    println!("{}", process(include_str!("./input.txt")))
}

#[derive(Clone, Copy, Debug)]
struct Game {
    matches: usize,
    count: i32,
}

impl Game {
    fn new(data: &str) -> Self {
        let parts: Vec<&str> = data.split(':').collect();
        let separated_numbers: Vec<&str> = parts[1].split('|').collect();
        Game {
            matches: count_matches(separated_numbers),
            count: 1,
        }
    }
}

fn count_matches(input: Vec<&str>) -> usize {
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

    let mut count = 0;
    for v in right.iter() {
        if left.contains(v) {
            count += 1
        }
    }
    count
}

fn process(input: &str) -> i32 {
    let mut games: Vec<Game> = input.lines().map(Game::new).collect();
    for i in 0..games.len() {
        for _ in 1..=games[i].count {
            if games[i].matches > 0 {
                for c in 1..=games[i].matches {
                    games[i + c].count += 1;
                }
            }
        }
    }
    games.iter().map(|v| v.count).sum()
}

#[cfg(test)]
mod test {

    use super::*;

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
            30
        )
    }
}
