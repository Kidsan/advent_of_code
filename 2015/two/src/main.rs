use ::std::fs;

struct BoxDimension {
    length: i32,
    width: i32,
    height: i32,
}

fn main() {
    println!("Hello, world!");
    let input: Vec<BoxDimension> = read("input.txt");
    println!("Part One: {}", calculate_paper_size(&input));
    println!("Part Two: {}", calculate_ribbon_size(&input));
}

fn read(filename: &str) -> Vec<BoxDimension> {
    let contents = fs::read_to_string(filename).expect("Something went wrong reading the file");

    let mut result: Vec<BoxDimension> = vec![];

    for b in contents.split("\n") {
        let sizes: Vec<&str> = b.split("x").collect();

        result.push(BoxDimension {
            length: sizes[0].parse::<i32>().unwrap(),
            width: sizes[1].parse::<i32>().unwrap(),
            height: sizes[2].parse::<i32>().unwrap(),
        })
    }
    result
}

fn calculate_paper_size(input: &Vec<BoxDimension>) -> i32 {
    let mut result = 0;

    for box_size in input {
        let mut smallest_area = box_size.length * box_size.width;
        if (box_size.width * box_size.height) <= smallest_area {
            smallest_area = box_size.width * box_size.height;
        }

        if (box_size.length * box_size.height) <= smallest_area {
            smallest_area = box_size.length * box_size.height;
        }

        let paper_needed = (2 * box_size.length * box_size.width)
            + (2 * box_size.width * box_size.height)
            + (2 * box_size.height * box_size.length);
        result += paper_needed + smallest_area
    }
    result
}

fn calculate_ribbon_size(input: &Vec<BoxDimension>) -> i32 {
    let mut result = 0;

    for box_size in input {
        let bow_size = box_size.length * box_size.width * box_size.height;

        let around_vertically = (box_size.width * 2) + (box_size.height * 2);
        let around_horizontally = (box_size.length * 2) + (box_size.width * 2);
        let around_other = (box_size.height * 2) + (box_size.length * 2);

        let mut smallest_area = around_vertically;

        if around_horizontally <= smallest_area {
            smallest_area = around_horizontally
        }

        if around_other <= smallest_area {
            smallest_area = around_other
        }

        result += bow_size + smallest_area
    }
    result
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn example() {
        let boxes = vec![BoxDimension {
            length: 2,
            width: 3,
            height: 4,
        }];
        let expected = 58;

        assert_eq!(expected, calculate_paper_size(boxes));
    }

    #[test]
    fn example_two() {
        let boxes = vec![BoxDimension {
            length: 1,
            width: 1,
            height: 10,
        }];
        let expected = 43;

        assert_eq!(expected, calculate_paper_size(boxes));
    }

    #[test]
    fn example_three() {
        let boxes = vec![
            BoxDimension {
                length: 2,
                width: 3,
                height: 4,
            },
            BoxDimension {
                length: 1,
                width: 1,
                height: 10,
            },
        ];
        let expected = 101;

        assert_eq!(expected, calculate_paper_size(boxes));
    }

    #[test]
    fn ribbon_example_one() {
        let boxes = vec![
            BoxDimension {
                length: 2,
                width: 3,
                height: 4,
            }
        ];
        let expected = 34;

        assert_eq!(expected, calculate_ribbon_size(boxes));
    }

    #[test]
    fn ribbon_example_two() {
        let boxes = vec![
            BoxDimension {
                length: 1,
                width: 1,
                height: 10,
            }
        ];
        let expected = 14;

        assert_eq!(expected, calculate_ribbon_size(boxes));
    }
}
