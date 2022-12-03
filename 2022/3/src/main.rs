use std::{fmt::Error, fs};

fn main() {
    let contents = fs::read_to_string("input.txt").expect("Something went wrong reading the file");

    let input: Vec<&str> = contents.split("\n").collect();
    let result_one: usize = part_one(input.clone());
    let result_two: usize = part_two(input.clone());

    println!("Part One {}", result_one);
    println!("Part Two {}", result_two);
}

fn find_duplicated(input: &str) -> Result<char, Error> {
    let (first, last) = input.split_at(input.len() / 2);

    for letter in first.chars() {
        if last.contains(letter) {
            return Ok(letter);
        }
    }

    for letter in last.chars() {
        if first.contains(letter) {
            return Ok(letter);
        }
    }

    return Err(Error);
}

fn find_duplicated_p2(input: Vec<&str>) -> Result<char, Error> {
    let (first, second, third) = (input[0], input[1], input[2]);

    for letter in first.chars() {
        if second.contains(letter) && third.contains(letter) {
            return Ok(letter);
        }
    }

    for letter in second.chars() {
        if first.contains(letter) && third.contains(letter) {
            return Ok(letter);
        }
    }

    for letter in third.chars() {
        if first.contains(letter) && second.contains(letter) {
            return Ok(letter);
        }
    }

    return Err(Error);
}

fn find_priority(input: char) -> usize {
    " abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
        .find(input)
        .unwrap()
}

fn part_one(input: Vec<&str>) -> usize {
    input
        .iter()
        .map(|bag| match find_duplicated(bag) {
            Ok(d) => find_priority(d),
            Err(_) => 0,
        })
        .sum()
}

fn part_two(input: Vec<&str>) -> usize {
    let mut result = 0;

    let mut current = 0;

    while current < input.len() {
        let mut group_priority = 0;

        let group = [input[current], input[current + 1], input[current + 2]];

        group_priority += match find_duplicated_p2(group.to_vec()) {
            Ok(d) => find_priority(d),
            Err(_) => 0,
        };

        result += group_priority;
        current += 3;
    }

    result
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_find_duplicated() {
        assert!(find_duplicated("abca").is_ok());
        assert_eq!(find_duplicated("abca"), Ok('a'));

        assert!(!find_duplicated("abcd").is_ok());
        assert_eq!(find_duplicated("abcd"), Err(Error));

        assert!(find_duplicated("vJrwpWtwJgWrhcsFMMfFFhFp").is_ok());
        assert_eq!(find_duplicated("vJrwpWtwJgWrhcsFMMfFFhFp"), Ok('p'));
    }

    #[test]
    fn test_find_priority() {
        let expected: usize = 2;
        assert_eq!(find_priority('b'), expected);

        let expected: usize = 52;
        assert_eq!(find_priority('Z'), expected);
    }

    #[test]
    fn test_part_one() {
        let input = [
            "vJrwpWtwJgWrhcsFMMfFFhFp",
            "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
            "PmmdzqPrVvPwwTWBwg",
            "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
            "ttgJtRGJQctTZtZT",
            "CrZsJsPPZsGzwwsLwLmpwMDw",
        ]
        .to_vec();
        assert_eq!(part_one(input), 157);
    }

    #[test]
    fn test_part_two() {
        let input = [
            "vJrwpWtwJgWrhcsFMMfFFhFp",
            "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
            "PmmdzqPrVvPwwTWBwg",
            "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
            "ttgJtRGJQctTZtZT",
            "CrZsJsPPZsGzwwsLwLmpwMDw",
        ]
        .to_vec();
        assert_eq!(part_two(input), 70);
    }
}
