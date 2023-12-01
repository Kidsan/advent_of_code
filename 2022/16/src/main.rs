use std::{cell::RefCell, cmp, collections::HashMap};

use nom::{
    branch::alt,
    bytes::complete::tag,
    character::complete::{alpha0, i32},
    combinator::{all_consuming, map},
    multi::separated_list0,
    sequence::tuple,
    Finish, IResult,
};
#[derive(PartialEq, Debug, Clone)]
struct Valve<'a> {
    name: &'a str,
    flow_rate: i32,
    tunnels: Vec<&'a str>,
    mask: i64,
}

fn main() {
    let (_, valves): (&str, Vec<Valve>) =
        all_consuming(parse_all_valves)(include_str!("../input.txt"))
            .finish()
            .unwrap();

    let mut valve_map = HashMap::new();

    for (i, valve) in valves.iter().enumerate() {
        let mut v = valve.clone();
        v.mask = i64::pow(2, i as u32);
        valve_map.insert(valve.name, v);
    }

    let distances = calc_distances(&valve_map);
    let mut part_one_answer: HashMap<i64, i64> = HashMap::new();
    let mut part_two_answer: HashMap<i64, i64> = HashMap::new();

    let final_answer = visit(
        String::from("AA"),
        30,
        0,
        &valve_map,
        &distances,
        0,
        &mut part_one_answer,
    );

    let _part_two = visit(
        String::from("AA"),
        26,
        0,
        &valve_map,
        &distances,
        0,
        &mut part_two_answer,
    );
    let part_one = final_answer.values().max().unwrap();

    let mut part_two_result = 0;

    for (k1, v1) in part_two_answer.iter() {
        for (k2, v2) in part_two_answer.iter() {
            if (k1 & k2) == 0 && (v1 + v2) > part_two_result {
                part_two_result = v1 + v2;
            }
        }
    }

    println!("partone = {}", part_one);
    println!("part_two = {}", part_two_result);
}

fn visit<'a>(
    valve: String,
    budget: i64,
    state: i64,
    cave: &HashMap<&str, Valve>,
    distances: &HashMap<(String, String), RefCell<i64>>,
    flow: i64,
    answer: &'a mut HashMap<i64, i64>,
) -> &'a mut HashMap<i64, i64> {
    let n = if !answer.contains_key(&state) {
        0
    } else {
        *answer.get(&state).unwrap()
    };
    answer.insert(state, cmp::max(n, flow));
    for k in cave
        .iter()
        .filter(|(_, cv)| cv.flow_rate > 0)
        .map(|(ck, _)| ck)
    {
        let dist: i64;
        {
            dist = *distances
                .get(&(valve.clone(), k.to_string()))
                .unwrap()
                .borrow();
        }
        let new_budget = budget - dist - 1;
        let mask = cave.get(k).unwrap().mask;
        if (state & mask) != 0 || new_budget < 0 {
            continue;
        } else {
            let flow_here = cave.get(k).unwrap().flow_rate as i64;
            let _ = visit(
                k.to_string(),
                new_budget,
                state | mask,
                cave,
                distances,
                flow + (new_budget * flow_here),
                answer,
            );
        }
    }
    answer
}

fn calc_distances(cave: &HashMap<&str, Valve>) -> HashMap<(String, String), RefCell<i64>> {
    let mut distances: HashMap<(String, String), RefCell<i64>> = HashMap::new();
    cave.keys().for_each(|x| {
        cave.keys().for_each(|y| {
            if cave.get(x).unwrap().tunnels.contains(y) {
                distances
                    .entry((x.to_string(), y.to_string()))
                    .or_insert(RefCell::new(1));
            } else {
                distances
                    .entry((x.to_string(), y.to_string()))
                    .or_insert(RefCell::new(i64::MAX));
            }
        });
    });
    cave.keys().for_each(|k| {
        cave.keys().for_each(|i| {
            cave.keys().for_each(|j| {
                let ij: i64;
                let ik: i64;
                let kj: i64;
                let tmp: i64;
                {
                    ij = *distances
                        .get(&(i.to_string(), j.to_string()))
                        .unwrap()
                        .borrow();
                }
                {
                    ik = *distances
                        .get(&(i.to_string(), k.to_string()))
                        .unwrap()
                        .borrow();
                }
                {
                    kj = *distances
                        .get(&(k.to_string(), j.to_string()))
                        .unwrap()
                        .borrow();
                }
                if ik == i64::MAX || kj == i64::MAX {
                    // workaround to avoid overflow on addition
                    tmp = cmp::min(ij, i64::MAX);
                } else {
                    tmp = cmp::min(ij, ik + kj);
                }
                {
                    distances.insert((i.to_string(), j.to_string()), RefCell::new(tmp));
                }
            });
        });
    });
    return distances;
}

fn parse_valve(input: &str) -> IResult<&str, Valve> {
    let (i, (_0, _1, _2, _3, _4, _5)) = tuple((
        tag("Valve "),
        nom::character::complete::alpha0::<&str, nom::error::Error<&str>>,
        tag(" has flow rate="),
        i32,
        alt((
            tag("; tunnels lead to valves "),
            tag("; tunnel leads to valve "),
        )),
        alt((
            separated_list0(tag(", "), alpha0),
            map(alpha0, |d: &str| vec![d]),
        )),
    ))(input)
    .unwrap();

    Ok((
        i,
        Valve {
            name: _1,
            flow_rate: _3,
            tunnels: _5,
            mask: 1,
        },
    ))
}

fn parse_all_valves(i: &str) -> IResult<&str, Vec<Valve>> {
    separated_list0(tag("\n"), parse_valve)(i)
}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn test_parse_valve() {
        assert_eq!(
            (
                "",
                Valve {
                    name: "AA",
                    flow_rate: 0,
                    tunnels: vec!["DD", "II", "BB"],
                    mask: 1,
                }
            ),
            parse_valve("Valve AA has flow rate=0; tunnels lead to valves DD, II, BB").unwrap()
        );

        assert_eq!(
            (
                "",
                Valve {
                    name: "AA",
                    flow_rate: 100,
                    tunnels: vec!["DD"],
                    mask: 1
                }
            ),
            parse_valve("Valve AA has flow rate=100; tunnel leads to valve DD").unwrap()
        );
    }
}
