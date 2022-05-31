use std::{char, fs};

fn main() {
    let input: Vec<String> = read("input.txt");
    println!("Part One: {}", count_nice(&input));
    // println!("Part Two: {}", calculate_ribbon_size(&input));
}

fn read(filename: &str) -> Vec<String> {
    let contents = fs::read_to_string(filename).expect("Something went wrong reading the file");

    let mut result: Vec<String> = vec![];

    for b in contents.split("\n") {
        result.push(b.to_string())
    }
    result
}

fn count_nice(input: &Vec<String>) -> i32 {
    let mut result = 0;

    for v in input {
        if validate(v) {
            result += 1
        }
    }

    result
}

fn validate(input: &String) -> bool {
    let mut count_of_vowels = 0;
    let mut has_double = false;
    let mut prev: Option<char> = None;

    for letter in input.chars().enumerate() {
        let v = letter.1;

        match v {
            'a' | 'e' | 'i' | 'o' | 'u' => count_of_vowels += 1,
            _ => {}
        }

        if Some(v) == prev {
            has_double = true
        } else {
            prev = Some(v);
        }
    }

    if input.contains("ab") || input.contains("cd") || input.contains("pq") || input.contains("xy")
    {
        return false;
    }

    if count_of_vowels >= 3 && has_double {
        return true;
    }

    false
}

#[cfg(test)]
mod tests {
    use super::*;

    struct TestCase {
        input: String,
        expected: bool,
    }

    #[test]
    fn test_validate() {
        let tests = vec![
            TestCase {
                input: "ugknbfddgicrmopn".to_string(),
                expected: true,
            },
            TestCase {
                input: "aaa".to_string(),
                expected: true,
            },
            TestCase {
                input: "jchzalrnumimnmhp".to_string(),
                expected: false,
            },
            TestCase {
                input: "haegwjzuvuyypxyu".to_string(),
                expected: false,
            },
            TestCase {
                input: "dvszwmarrgswjxmb".to_string(),
                expected: false,
            },
        ];

        for tc in tests {
            let res = validate(&tc.input);
            assert_eq!(
                tc.expected, res,
                "using: {}, expected {} and got {}",
                tc.input, tc.expected, res
            );
        }
    }
}
