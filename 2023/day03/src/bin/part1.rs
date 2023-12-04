fn main() {
    println!("Hello, world!");
}

fn to_grid(input: &str) -> Vec<Vec<String>> {
    input
        .lines()
        .map(|line| {
            line.split_inclusive('.')
                .flat_map(|w| w.split_inclusive(char::is_numeric).collect::<Vec<_>>())
                .collect()
        })
        .map(|line: Vec<&str>| {
            let mut res: Vec<String> = vec![];
            let mut num = String::new();
            for v in line.clone().into_iter() {
                match v {
                    _x if v.parse::<i32>().is_ok() => {
                        num.push(v.chars().next().unwrap());
                    }
                    _ => {
                        if !num.is_empty() {
                            let r = num.clone();
                            res.push(r);
                            num = String::new();
                        }
                        if v.contains('.') && v != "." {
                            for ch in v.split('.') {
                                if !ch.is_empty() {
                                    res.push(ch.to_string())
                                }
                            }
                        } else {
                            res.push(v.to_string());
                        }
                    }
                }
            }
            res
        })
        .collect::<Vec<Vec<String>>>()
}

fn process(input: &str) -> i32 {
    let grid = to_grid(input);
    let pos: Vec<(i32, i32)> = vec![
        (0, 1),
        (1, 0),
        (1, 1),
        (-1, 0),
        (-1, -1),
        (0, -1),
        (1, -1),
        (-1, 1),
    ];
    let mut res: i32 = 0;
    let mut values: Vec<i32> = vec![];

    for (y, line) in grid.iter().enumerate() {
        for (x, v) in line.iter().enumerate() {
            if v.parse::<i32>().is_ok() {
                dbg!("current value is", v);
                for (mx, my) in pos.iter() {
                    let nx: i32 = x as i32 + *mx;
                    let ny: i32 = y as i32 + *my;
                    if nx < v.len() as i32 && nx >= 0 && ny < line.len() as i32 && ny >= 0 {
                        dbg!(true);
                        let inspection_value = grid[ny as usize][nx as usize].clone();
                        if inspection_value == "." {
                            dbg!(inspection_value);
                        } else if inspection_value.parse::<i32>().is_ok() {
                        } else {
                            // res += v.parse::<i32>().unwrap();
                            values.push(v.parse::<i32>().unwrap())
                        }
                    }
                }
            }
        }
    }
    // 4361
    dbg!(&values);
    res = values.iter().sum();
    res
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_process() {
        assert_eq!(
            process(
                "467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598.."
            ),
            4361
        )
    }
    #[test]
    fn test_to_grid() {
        assert_eq!(
            to_grid(
                "467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598.."
            ),
            vec![
                vec!["467", ".", ".", "114", ".", "."],
                vec![".", ".", ".", "*", ".", ".", ".", ".", "."],
                vec![".", ".", "35", ".", ".", "633", "."],
                vec![".", ".", ".", ".", ".", ".", "#", ".", "."],
                vec!["617", "*", ".", ".", ".", ".", "."],
                vec![".", ".", ".", ".", ".", "+", "58", "."],
                vec![".", ".", "592", ".", ".", ".", ".", "."],
                vec![".", ".", ".", ".", ".", ".", "755", "."],
                vec![".", ".", ".", "$", "*", ".", ".", "."],
                vec![".", "664", ".", "598", ".", ".",],
            ]
        )
    }
}
