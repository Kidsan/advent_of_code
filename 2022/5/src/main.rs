use itertools::Itertools;
use scan_fmt::scan_fmt;
use std::collections::HashMap;

fn main() {
    let contents = include_str!("../input.txt");

    let parts: Vec<&str> = contents.split("\n\n").collect();

    let stacks = parts[0];
    let instructions = parts[1];

    let mut stacks = parse_stacks(stacks);
    let parsed_instructions = parse_instructions(instructions);

    for i in parsed_instructions.iter() {
        let source = stacks.get_mut(&i.source).unwrap();
        let mut to_add: Vec<&str> = Vec::new();
        print!("Instruction: {:#?} \n", i);

        let mut n = 0;

        while n < i.amount {
            if let Some(b) = source.pop() {
                to_add.push(b);
            }
            n += 1;
        }

        // to_add.reverse(); // this changes between p1 and p2
        print!(" to_add: {:#?} ", to_add);
        let dest = stacks.get_mut(&i.dest).unwrap();

        for v in to_add.iter() {
            dest.push(v);
        }
    }

    print!("ending stacks: {:#?}", stacks);

    let mut part_one_result = "".to_string();

    for (_, v) in stacks.iter().sorted() {
        // if v.len() > 0 {
        part_one_result += v[v.len() - 1];
        // }
    }

    print!("Part One: {}", part_one_result); // QPJPLMNNR
}

fn parse_stacks(stacks: &str) -> HashMap<i32, Vec<&str>> {
    let mut starting_stacks: HashMap<i32, Vec<&str>> = HashMap::new();
    let mut rows: Vec<Vec<&str>> = stacks
        .split("\n")
        .map(|row| row.split("").collect())
        .collect();

    let final_row = rows.pop().unwrap();
    let mut index = 2;

    while index < final_row.len() {
        let number: i32 = final_row[index].parse().unwrap();
        starting_stacks.insert(number, Vec::new());
        index += 4;
    }

    for row in rows {
        let mut n = 2;

        while n < row.len() {
            let c = row[n];
            if c != " " {
                let belongs_to: i32 = final_row[n].parse().unwrap();

                starting_stacks.get_mut(&belongs_to).unwrap().push(c);
            }

            n += 4;
        }
    }

    for (_, row) in starting_stacks.iter_mut() {
        row.reverse();
    }

    return starting_stacks;
}

#[derive(Debug)]
struct Instruction {
    amount: i32,
    source: i32,
    dest: i32,
}

fn parse_instructions(input: &str) -> Vec<Instruction> {
    // print!(" input: {:#?} ", input);
    return input
        .lines()
        .map(|line| scan_fmt!(line, "move {d} from {d} to {d}", i32, i32, i32).unwrap())
        .map(|(a, b, c)| Instruction {
            amount: a,
            source: b,
            dest: c,
        })
        .collect();
}
