use std::fs;

fn main() {
    let input: String = read("input.txt");
    println!("Part One: {}", calculate_floor(&input));
    println!("Part Two: {}", calculate_floor_part_two(&input));
}

fn read(filename: &str) -> String {
    fs::read_to_string(filename).expect("Something went wrong reading the file")
}

fn calculate_floor(input: &str) -> i32 {
    let mut result = 0;
    for letter in input.chars() {
        match letter {
            '(' => result += 1,
            ')' => result -= 1,
            _ => (),
        }
    }
    result
}

fn calculate_floor_part_two(input: &str) -> i32 {
    let mut floor = 0;
    for letter in input.chars().enumerate() {
        match letter.1 {
            '(' => floor += 1,
            ')' => floor -= 1,
            _ => (),
        }

        if floor < 0 {
            let pos: i32 = letter.0.try_into().unwrap();
            return pos + 1;
        }
    }
    floor
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn zero() {
        let contents = ["(())", "()()"];
        let expected = 0;

        for input in contents.iter() {
            assert_eq!(expected, calculate_floor(input));
        }
    }

    #[test]
    fn three() {
        let contents = ["(((", "(()(()(", "))((((("];
        let expected = 3;

        for input in contents.iter() {
            assert_eq!(expected, calculate_floor(input));
        }
    }

    #[test]
    fn negative_one() {
        let contents = ["())", "))("];
        let expected = -1;

        for input in contents.iter() {
            assert_eq!(expected, calculate_floor(input));
        }
    }

    #[test]
    fn negative_three() {
        let contents = [")))", ")())())"];
        let expected = -3;

        for input in contents.iter() {
            assert_eq!(expected, calculate_floor(input));
        }
    }

    #[test]
    fn test_calculate_floor_part_two() {
        let contents = [")"];
        let expected = 1;

        for input in contents.iter() {
            assert_eq!(expected, calculate_floor_part_two(input));
        }
    }

}
