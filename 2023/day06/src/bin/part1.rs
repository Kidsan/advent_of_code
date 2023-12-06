use std::vec;

fn main() {
    println!("solution: {}", process(include_str!("input.txt")))
}

fn re(i: Vec<(i64, i64)>) -> i64 {
    let mut v = vec![];

    for race in i.iter() {
        let mut win_window_start = 0;
        let mut win_window_end = 0;
        for i in 1..race.0 {
            if win_window_start == 0 && i * (race.0 - i) > race.1 {
                win_window_start = i;
            }

            if win_window_end == 0 {
                let speed = race.0 - i;
                let time_holding = race.0 - i;
                let time_remaining = race.0 - time_holding;
                if speed * time_remaining > race.1 {
                    win_window_end = race.0 - i;
                }
            }
            if win_window_start != 0 && win_window_end != 0 || i > race.0 - i {
                dbg!(win_window_start, win_window_end);
                break;
            }
        }
        v.push((win_window_end - win_window_start) + 1)
    }
    v.iter().product()
}

fn process(_input: &str) -> i64 {
    // i,j = 1,race_len -1
    // would i win? would j win?
    // if yes, found start, if no increment, same for j but decrement
    // if found start + end, ans for race is end - start
    288
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_re() {
        assert_eq!(re(vec![(7, 9), (15, 40), (30, 200)]), 288);
        assert_eq!(
            re(vec![(57, 291), (72, 1172), (69, 1176), (92, 2026)]),
            160816
        );
        assert_eq!(re(vec![(57726992, 291117211762026)]), 46561107)
    }

    #[test]
    fn test_process() {
        assert_eq!(
            process(
                "Time:      7  15   30
Distance:  9  40  200"
            ),
            288
        )
    }
}
