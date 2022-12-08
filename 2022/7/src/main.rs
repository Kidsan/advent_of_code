use std::{
    collections::{HashMap, HashSet},
    path::PathBuf,
};

// got help on this one, should try rewrite it using nom to parse commands maybe
fn main() {
    let contents = include_str!("../input.txt");

    let fs = parse_file(contents);

    let mut sizes = HashMap::new();
    for k in fs.keys() {
        compute_dir_size(&fs, &mut sizes, k);
    }
    let total_size = sizes[&PathBuf::from("/")];
    let part_one = sizes.values().filter(|&&s| s <= 100000).sum::<i64>();
    let part_two = sizes
        .values()
        .filter(|&&s| 40000000 + s >= total_size)
        .min()
        .copied()
        .unwrap();

    println!("Part One: {:#?}", part_one);
    println!("Part Two: {}", part_two);
}

fn parse_file(input: &str) -> HashMap<PathBuf, HashSet<(i64, &str)>> {
    let mut fs = HashMap::new();
    let mut pwd = PathBuf::new();

    for l in input.split('$').skip(1) {
        match l.trim().lines().next().unwrap() {
            "ls" => {
                let entries = l.lines().skip(1).map(|output| {
                    let (size, f) = output.split_once(' ').unwrap();
                    (size.parse::<i64>().unwrap_or(-1), f)
                });
                fs.entry(pwd.clone())
                    .or_insert(HashSet::new())
                    .extend(entries);
            }
            "cd .." => {
                pwd.pop();
            }
            cd_dir => {
                pwd.push(cd_dir.split_once(' ').unwrap().1);
            }
        }
    }
    return fs;
}

fn compute_dir_size(
    fs: &HashMap<PathBuf, HashSet<(i64, &str)>>,
    sizes: &mut HashMap<PathBuf, i64>,
    dir: &PathBuf,
) {
    if sizes.contains_key(dir) {
        return;
    }
    let size = fs[dir]
        .iter()
        .map(|&(s, d)| match s {
            -1 => {
                let dir = dir.join(d);
                compute_dir_size(fs, sizes, &dir);
                sizes[&dir]
            }
            s => s,
        })
        .sum();
    sizes.insert(dir.clone(), size);
}
