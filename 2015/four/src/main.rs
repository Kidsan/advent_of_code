fn main() {
    println!("Part One: {}", solve(&"iwrupvqb".to_string(), "00000"));
    println!("Part Two: {}", solve(&"iwrupvqb".to_string(), "000000"));
}

fn build_input(input: &String, ind: i32) -> String {
    input.to_string() + &ind.to_string()
}

fn solve(input: &String, desired_prefix: &str) -> i32 {
    let mut i = 0;
    
    loop {
        let v = build_input(input, i);
        let digest = md5::compute(v);
        let f = format!("{:x}", digest);
        if f.starts_with(desired_prefix) {
            break;
        }
        i += 1
    }
    return i
}

