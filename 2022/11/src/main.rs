use std::collections::VecDeque;

fn main() {
    let contents = include_str!("../input.txt")
        .split("\n\n")
        .collect::<Vec<&str>>();

    // println!("{contents:#?}");

    let mut parsed_monkeys = parse_input(contents);
    let mut part_one_inspections = simulate_throws(&mut parsed_monkeys, 20);

    part_one_inspections.sort();
    let part_one_result = part_one_inspections[part_one_inspections.len() - 1]
        * part_one_inspections[part_one_inspections.len() - 2];

    // println!("{:?}", monkeys);

    // assert_eq!(16020, part_one_result);
    println!("part_one_result: {}", part_one_result);
}

#[derive(Debug, Clone)]
struct Monkey<'a> {
    items: VecDeque<i64>,
    operation: &'a str,
    operator_value: i64,
    divisor: i64,
    pass_target: i64,
    fail_target: i64,
}

fn parse_input(monkey_list: Vec<&str>) -> Vec<Monkey> {
    monkey_list
        .iter()
        .map(|monkey| monkey.split("\n").collect::<Vec<&str>>())
        .map(|monkey_lines| {
            // println!("{:?}", monkey_lines);
            let mut items = VecDeque::new();
            let mut operation = "";
            let mut operator_value = 0;
            let mut divisor: i64 = 0;
            let mut pass_target = 0;
            let mut fail_target = 0;

            monkey_lines.iter().for_each(|line| {
                // println!("foo: {line}");
                if line.starts_with("  Starting items: ") {
                    let list = line.trim_start_matches("  Starting items: ");
                    items = list
                        .split(", ")
                        .map(|n| n.parse::<i64>().unwrap())
                        .collect::<VecDeque<i64>>()
                }

                if line.starts_with("  Operation: new = old") {
                    let op = line.trim_start_matches("  Operation: new = old ");
                    let parts = op.split(" ").collect::<Vec<&str>>();
                    operation = parts[0];

                    if parts[1] == "old" {
                        operation = "sq";
                        operator_value = 0;
                    } else {
                        operator_value = parts[1].parse().unwrap();
                    }
                }

                if line.starts_with("  Test: divisible by ") {
                    let m = line.trim_start_matches("  Test: divisible by ");
                    divisor = m.parse().unwrap();
                }

                if line.starts_with("    If true: throw to monkey ") {
                    pass_target = line
                        .trim_start_matches("    If true: throw to monkey ")
                        .parse()
                        .unwrap();
                }

                if line.starts_with("    If false: throw to monkey ") {
                    fail_target = line
                        .trim_start_matches("    If false: throw to monkey ")
                        .parse()
                        .unwrap();
                }
            });

            return Monkey {
                items,
                operation,
                operator_value,
                divisor,
                pass_target,
                fail_target,
            };
        })
        .collect::<Vec<Monkey>>()
}

fn simulate_throws(monkey_list: &mut Vec<Monkey>, iterations: i64) -> Vec<i64> {
    let mut monkey_items: Vec<VecDeque<i64>> = Vec::new();

    for monkey in monkey_list.clone() {
        monkey_items.push(monkey.items.clone());
    }

    let first_position = monkey_items.clone();

    let mut monkey_inspections = vec![0; monkey_list.len()];

    for _ in 0..iterations {
        for (index, monkey) in monkey_list.iter().enumerate() {
            let count = monkey_items[index].len();

            for _ in 0..count {
                let mut i = monkey_items[index].pop_front().unwrap();
                i = apply_operation(monkey.operation, monkey.operator_value, i);
                i = i / &3;
                if i % monkey.divisor == 0 {
                    monkey_items[monkey.pass_target as usize].push_back(i);
                } else {
                    monkey_items[monkey.fail_target as usize].push_back(i);
                }
                monkey_inspections[index] += 1;
            }
        }

        if monkey_items == first_position {
            println!("ay");
        }
    }
    monkey_inspections
}

fn apply_operation(operator: &str, operator_value: i64, item: i64) -> i64 {
    match operator {
        "sq" => item * item,
        "+" => item + operator_value,
        "-" => item - operator_value,
        "*" => item * operator_value,
        _ => {
            println!("{operator}");
            panic!("Unknown Operator")
        }
    }
}
