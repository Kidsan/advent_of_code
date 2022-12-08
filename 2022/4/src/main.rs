use std::ops::RangeInclusive;

trait InclusiveRangeExt {
    fn contains_range(&self, other: &Self) -> bool;

    fn contains_or_is_contained(&self, other: &Self) -> bool {
        self.contains_range(other) || other.contains_range(self)
    }

    fn overlaps(&self, other: &Self) -> bool;

    fn overlaps_or_is_overlapped(&self, other: &Self) -> bool {
        self.overlaps(other) || other.overlaps(self)
    }
}

impl<T> InclusiveRangeExt for RangeInclusive<T>
where
    T: PartialOrd,
{
    fn contains_range(&self, other: &Self) -> bool {
        self.contains(other.start()) && self.contains(other.end())
    }

    fn overlaps(&self, other: &Self) -> bool {
        self.contains(other.start()) || self.contains(other.end())
    }
}

fn main() {
    let contents = include_str!("../input.txt").lines().map(|line| {
        let parts: Vec<&str> = line.split(",").collect();

        let left_contraints: Vec<&str> = parts[0].split("-").collect();
        let right_contraints: Vec<&str> = parts[1].split("-").collect();

        // left
        let parsed_start = left_contraints[0].parse::<i32>().unwrap();
        let parsed_end = left_contraints[1].parse::<i32>().unwrap();

        // right
        let right_parsed_start = right_contraints[0].parse::<i32>().unwrap();
        let right_parsed_end = right_contraints[1].parse::<i32>().unwrap();

        return (
            parsed_start..=parsed_end,
            right_parsed_start..=right_parsed_end,
        );
    });

    let result_one = contents
        .clone()
        .filter(|(a, b)| a.contains_or_is_contained(&b))
        .count();

    let result_two = contents
        .clone()
        .filter(|(a, b)| a.overlaps_or_is_overlapped(&b))
        .count();

    println!("Part One {}", result_one);
    println!("Part Two {}", result_two);
}
