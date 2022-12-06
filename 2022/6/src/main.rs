fn main() {
    let contents = include_str!("../input.txt");

    println!("Part One: {}", first_word(contents, 4));
    println!("Part Two: {}", first_word(contents, 14));
}

fn first_word(contents: &str, word_size: usize) -> i32 {
    for (i, v) in contents.chars().enumerate() {
        let mut tmp: String = String::from(v);

        let chars = contents.chars().collect::<Vec<char>>();

        let mut n: usize = 1;

        while n < word_size - 1 {
            let next = chars[n + i];
            if tmp.contains(next) {
                break;
            }

            tmp += &String::from(next);
            n += 1;
        }

        if n == word_size - 1 {
            let next = chars[n + i];
            if tmp.contains(next) {
                continue;
            }

            return (word_size + i).try_into().unwrap();
        }
    }
    0
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_first_packet() {
        assert_eq!(first_word("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 4), 7);
    }

    #[test]
    fn test_first_messsage() {
        assert_eq!(first_word("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 14), 19);
        assert_eq!(first_word("bvwbjplbgvbhsrlpgdmjqwftvncz", 14), 23);
        assert_eq!(first_word("nppdvjthqldpwncqszvftbrmjlhg", 14), 23);
        assert_eq!(first_word("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 14), 29);
        assert_eq!(first_word("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 14), 26);
    }
}
