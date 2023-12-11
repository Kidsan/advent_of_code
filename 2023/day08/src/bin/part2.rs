use std::collections::HashMap;

fn main() {
    println!(
        "ans: {}",
        process(include_str!("./input.txt").to_string().trim().to_string())
    );
}

fn process(inp: String) -> usize {
    let spl: Vec<&str> = inp.split("\n\n").collect();
    let directions = spl[0].as_bytes();
    let m = spl[1];

    let path = m
        .split('\n')
        .map(|line| line.split(" = ").collect::<Vec<&str>>())
        .map(|parts| {
            let source = parts[0];

            let options: Vec<&str> = parts[1]
                .trim_matches(|c| c == '(' || c == ')')
                .split(", ")
                .collect();

            (source, (options[0], options[1]))
        })
        .collect::<HashMap<_, _>>();

    let mut count = 0;
    let mut n = "AAA";

    loop {
        n = if directions[count % directions.len()] == b'L' {
            path[n].0
        } else {
            path[n].1
        };
        if n == "ZZZ" {
            return count + 1;
        }
        count += 1;
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_process() {
        assert_eq!(
            process(
                "LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)"
                    .to_string()
            ),
            6
        );
    }
}
