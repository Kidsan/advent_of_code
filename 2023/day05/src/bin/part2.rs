use std::ops::Range;

fn main() {
    println!("solution: {}", process(include_str!("input.txt")))
}

#[derive(Debug, PartialEq)]
struct Translator {
    ranges: Vec<(std::ops::Range<i64>, i64)>,
}

impl Translator {
    fn new(input: Vec<Vec<i64>>) -> Self {
        let mut ranges = vec![];
        for i in input.iter() {
            ranges.push(((i[0]..i[0] + i[2]), i[1] - i[0]))
        }
        Self { ranges }
    }

    fn translate(&self, v: &i64) -> i64 {
        let mut res = *v;
        for r in self.ranges.iter() {
            if r.0.contains(v) {
                res = v + r.1;
                return res;
            }
        }
        res
    }
}

fn build_map(input: &str) -> Translator {
    let ranges: Vec<Vec<i64>> = input
        .trim()
        .split(":\n")
        .nth(1)
        .unwrap()
        .split('\n')
        .map(|v| {
            v.split(' ')
                .map(|i| i.parse::<i64>().unwrap())
                .collect::<Vec<i64>>()
        })
        .collect();
    Translator::new(ranges)
}

fn process(input: &str) -> i64 {
    let parts = input.split("\n\n").collect::<Vec<&str>>();
    let seed_numbers: Vec<i64> = parts[0]
        .split(": ")
        .nth(1)
        .unwrap()
        .split(' ')
        .map(|v| v.parse::<i64>().unwrap())
        .collect();

    let mut current_v = 0;

    let seeds: Vec<Range<i64>> = seed_numbers
        .chunks(2)
        .map(|x| {
            let start = x[0];
            let count = x[1];
            Range {
                start,
                end: start + count,
            }
        })
        .collect();

    let soil = build_map(parts[1]);
    let fert = build_map(parts[2]);
    let water = build_map(parts[3]);
    let light = build_map(parts[4]);
    let temp = build_map(parts[5]);
    let humidity = build_map(parts[6]);
    let loc = build_map(parts[7]);

    let mut res = 0;
    let mut found = false;

    while !found {
        let loc_v = loc.translate(&current_v);
        let h = humidity.translate(&loc_v);
        let temp = temp.translate(&h);
        let lv = light.translate(&temp);
        let wv = water.translate(&lv);
        let fv = fert.translate(&wv);
        let sv = soil.translate(&fv);
        for r in seeds.iter() {
            if r.contains(&sv) {
                found = true;
                res = current_v;
            }
        }
        current_v += 1
    }

    res
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_build_map() {
        assert_eq!(
            build_map(
                "seed-to-soil map:
50 98 2
52 50 48"
            ),
            Translator {
                ranges: vec![((50..52), 48), ((52..100), -2)]
            }
        )
    }

    #[test]
    fn test_process() {
        assert_eq!(
            process(
                "seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4"
            ),
            46
        )
    }
}
