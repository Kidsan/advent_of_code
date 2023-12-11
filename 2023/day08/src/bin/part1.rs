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
                "RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)"
                    .to_string()
            ),
            2
        );
        assert_eq!(
            process(
                "LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)"
                    .to_string()
            ),
            6
        )
    }
}
