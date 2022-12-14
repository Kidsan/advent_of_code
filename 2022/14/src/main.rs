use nom::{
    bytes::complete::tag, character::complete::u64, combinator::all_consuming,
    multi::separated_list0, sequence::tuple, Finish, IResult,
};

fn main() {
    let (_, mut lines): (&str, Vec<Vec<(u64, u64)>>) =
        all_consuming(parse_all_coordinates)(include_str!("../input.txt"))
            .finish()
            .unwrap();

    for n in 0..lines.len() {
        fill_missing_points(&mut lines[n]);
    }

    let mut cave: Vec<Vec<u64>> = vec![vec![0; 1000]; 1000]; // try do it without simulating the grid -> use a hashset of points
    let mut lowest_rock = 0;

    for line in lines {
        for point in line {
            cave[point.0 as usize][point.1 as usize] = 1;
            if point.1 > lowest_rock {
                lowest_rock = point.1;
            }
        }
    }

    let part_one_units = part_one(&mut cave.clone(), lowest_rock);
    let part_two_units = part_two(&mut cave.clone(), lowest_rock);

    println!("part_one: {}", part_one_units);
    println!("part_two: {}", part_two_units);
}

fn part_one(cave: &mut Vec<Vec<u64>>, lowest_rock: u64) -> u64 {
    const STARTING_POINT: (usize, usize) = (500, 0);

    let mut units: u64 = 0;

    'outer: loop {
        units += 1;
        let mut pos = STARTING_POINT;

        'inner: loop {
            if pos.1 > lowest_rock as usize {
                break 'outer;
            }
            if cave[pos.0][pos.1 + 1] == 0 {
                pos.1 += 1;
                continue 'inner;
            }

            if cave[pos.0][pos.1 + 1] == 1 {
                if cave[pos.0 - 1][pos.1 + 1] == 0 {
                    pos.0 -= 1;
                    pos.1 += 1;
                    continue 'inner;
                } else if cave[pos.0 + 1][pos.1 + 1] == 0 {
                    pos.0 += 1;
                    pos.1 += 1;
                    continue 'inner;
                } else {
                    cave[pos.0][pos.1] = 1;
                    continue 'outer;
                }
            }
        }
    }
    units - 1
}

fn part_two(cave: &mut Vec<Vec<u64>>, lowest_rock: u64) -> u64 {
    const STARTING_POINT: (usize, usize) = (500, 0);
    let floor = lowest_rock + 2;

    for x in 0..cave.len() {
        cave[x][floor as usize] = 1;
    }

    let mut units: u64 = 0;

    'outer: loop {
        units += 1;
        let mut pos = STARTING_POINT;

        'inner: loop {
            if cave[pos.0][pos.1] == 1
                && cave[pos.0 - 1][pos.1 + 1] == 1
                && cave[pos.0 + 1][pos.1 + 1] == 1
            {
                break 'outer;
            }
            if cave[pos.0][pos.1 + 1] == 0 {
                pos.1 += 1;
                continue 'inner;
            }

            if cave[pos.0][pos.1 + 1] == 1 {
                if cave[pos.0 - 1][pos.1 + 1] == 0 {
                    pos.0 -= 1;
                    pos.1 += 1;
                    continue 'inner;
                } else if cave[pos.0 + 1][pos.1 + 1] == 0 {
                    pos.0 += 1;
                    pos.1 += 1;
                    continue 'inner;
                } else {
                    cave[pos.0][pos.1] = 1;
                    continue 'outer;
                }
            }
        }
    }
    units - 1
}

fn parse_line(input: &str) -> IResult<&str, Vec<(u64, u64)>> {
    separated_list0(tag(" -> "), point_coordinates)(input)
}

fn point_coordinates(input: &str) -> IResult<&str, (u64, u64)> {
    let (i, (x, _, y)) =
        tuple((u64::<&str, nom::error::Error<&str>>, tag(","), u64))(input).unwrap(); //::<&str, (u64, &str, u64), nom::error::Error<&str>, _>
    Ok((i, (x, y)))
}

pub fn parse_all_coordinates(i: &str) -> IResult<&str, Vec<Vec<(u64, u64)>>> {
    separated_list0(tag("\n"), parse_line)(i)
}

fn fill_missing_points(lines: &mut Vec<(u64, u64)>) {
    for n in 0..lines.len() - 1 {
        let point = lines[n];
        let next_point = lines[n + 1];

        if point.0 == next_point.0 && point.1 < next_point.1 {
            let mut y = point.1 + 1;
            while y < next_point.1 {
                lines.push((point.0, y));
                y += 1;
            }
        }

        if point.0 == next_point.0 && point.1 > next_point.1 {
            let mut y = point.1 - 1;

            while y > next_point.1 {
                lines.push((point.0, y));
                y -= 1;
            }
        }

        if point.1 == next_point.1 {
            if point.0 < next_point.0 {
                let mut x = point.0 + 1;
                while x < next_point.0 {
                    lines.push((x, point.1));
                    x += 1;
                }
            }
            if point.0 > next_point.0 {
                let mut x = point.0 - 1;
                while x > next_point.0 {
                    lines.push((x, next_point.1));
                    x -= 1;
                }
            }
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_point_coordinates() {
        assert_eq!(("", (1, 1)), point_coordinates("1,1").unwrap());
        assert_eq!(("", (498, 4)), point_coordinates("498,4").unwrap());
    }

    #[test]
    fn test_parse_line() {
        assert_eq!(
            ("", vec![(498, 4), (498, 6), (496, 6)]),
            parse_line("498,4 -> 498,6 -> 496,6").unwrap()
        );
    }

    #[test]
    fn test_fill_missing_points() {
        let mut input = vec![(1, 1), (1, 2), (1, 3)];
        fill_missing_points(&mut input);
        assert_eq!(vec![(1, 1), (1, 2), (1, 3)], input);

        let mut input = vec![(1, 1), (1, 3), (1, 5), (5, 5)];
        fill_missing_points(&mut input);
        assert_eq!(
            vec![
                (1, 1),
                (1, 3),
                (1, 5),
                (5, 5),
                (1, 2),
                (1, 4),
                (2, 5),
                (3, 5),
                (4, 5),
            ],
            input
        );
    }
}
