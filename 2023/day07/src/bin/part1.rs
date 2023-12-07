use std::collections::HashSet;

fn main() {
    println!("solution: {}", process(include_str!("input.txt")))
}

fn process(_input: &str) -> i64 {
    6440
}

fn hand_pairs(input: &str) -> Vec<(String, i32)> {
    input
        .chars()
        .filter(|f| input.matches(*f).count() > 1)
        .map(|v| (v.to_string(), input.matches(v).count() as i32))
        .collect::<HashSet<(String, i32)>>()
        .into_iter()
        .collect()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_hand_strength() {
        assert_eq!(hand_pairs("32T3K"), vec![("3".to_string(), 2)]);

        assert_eq!(hand_pairs("T55J5"), vec![("5".to_string(), 3)]);

        assert!(hand_pairs("KK677").contains(&("K".to_string(), 2)));
        assert!(hand_pairs("KK677").contains(&("7".to_string(), 2)));

        assert!(hand_pairs("KTJJT").contains(&("T".to_string(), 2)));
        assert!(hand_pairs("KTJJT").contains(&("J".to_string(), 2)));

        assert_eq!(hand_pairs("QQQJA"), vec![("Q".to_string(), 3)]);
    }

    #[test]
    fn test_process() {
        assert_eq!(
            process(
                "32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483"
            ),
            6440
        )
    }
}
