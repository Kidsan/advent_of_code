fn main() {
    println!(
        "ans: {}",
        process(include_str!("./input.txt").to_string().trim())
    );
}

fn next_number(nums: &mut Vec<Vec<i32>>, curr_y: usize, curr_x: usize) -> i32 {
    dbg!(curr_y, curr_x, &nums[curr_y][curr_x - 1]);
    if nums[curr_y].iter().all(|v| *v == 0) {
        return 0;
    }
    if nums[curr_y + 1].len() == nums[curr_y].len() {
        return nums[curr_y + 1][curr_x - 1] + nums[curr_y][curr_x - 1];
    } else {
        let next_num = next_number(nums, curr_y + 1, curr_x - 1);
        nums[curr_y + 1].push(next_num);
    }
    nums[curr_y + 1][curr_x - 1] + nums[curr_y][curr_x - 1]
}

fn process(input: &str) -> i32 {
    input
        .lines()
        .map(|line| {
            let mut vals = vec![];
            vals.push(
                line.split(' ')
                    .map(str::parse::<i32>)
                    .map(|v| v.unwrap())
                    .collect::<Vec<i32>>(),
            );
            let mut curr = 0;
            while !vals[curr].iter().all(|v| *v == 0) {
                let mut n = vec![];

                for (index, val) in vals[curr].iter().enumerate() {
                    if index + 1 < vals[curr].len() {
                        n.push(vals[curr][index + 1] - val);
                    }
                }
                vals.push(n);
                curr += 1;
            }
            dbg!(&vals);
            let target_index = vals[0].len();
            next_number(&mut vals, 0, target_index)
        })
        .sum()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_next_number() {
        let mut nums = vec![vec![3, 3, 3, 3, 3], vec![0, 0, 0, 0]];
        assert_eq!(next_number(&mut nums, 0, 5), 3)
    }

    #[test]
    fn test_process() {
        assert_eq!(
            process(
                "0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45"
            ),
            114
        )
    }
}
