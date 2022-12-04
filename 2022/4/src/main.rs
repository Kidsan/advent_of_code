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
            (parsed_start, parsed_end),
            (right_parsed_start, right_parsed_end),
        );
    });

    let result_one = contents
        .clone()
        .map(|parsed_line| entire_overlap(parsed_line.0, parsed_line.1))
        .filter(|has_overlap| *has_overlap)
        .count();

    let result_two = contents
        .clone()
        .map(|parsed_line| partial_overlap(parsed_line.0, parsed_line.1))
        .filter(|has_overlap| *has_overlap)
        .count();

    println!("Part One {}", result_one);
    println!("Part Two {}", result_two);
}

fn entire_overlap(left: (i32, i32), right: (i32, i32)) -> bool {
    // left contains right
    if left.0 <= right.0 && left.1 >= right.1 {
        return true;
    }

    // right contains left
    if right.0 <= left.0 && right.1 >= left.1 {
        return true;
    }
    false
}

fn partial_overlap(left: (i32, i32), right: (i32, i32)) -> bool {
    for pos in left.0..=left.1 {
        if pos >= right.0 && pos <= right.1 {
            return true;
        }
    }
    false
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_entire_overlap() {
        assert_eq!(entire_overlap((2, 8), (3, 7)), true);
        assert_eq!(entire_overlap((6, 6), (4, 6)), true);
        assert_eq!(entire_overlap((2, 3), (4, 6)), false);
    }

    #[test]
    fn test_partial_overlap() {
        assert_eq!(partial_overlap((2, 8), (3, 7)), true);
        assert_eq!(partial_overlap((6, 6), (4, 6)), true);
        assert_eq!(partial_overlap((2, 3), (4, 6)), false);

        assert_eq!(partial_overlap((2, 6), (4, 8)), true);
    }
}
