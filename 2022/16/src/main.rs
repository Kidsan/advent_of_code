use std::collections::HashMap;

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
}

fn main() {
    let (_, valves): (&str, Vec<Valve>) =
        all_consuming(parse_all_valves)(include_str!("../input.txt"))
            .finish()
            .unwrap();

    let mut valveMap = HashMap::new();

    for valve in &valves {
        valveMap.insert(valve.name, valve);
    }

    let mut current_pos = "AA";
    let mut time_left = 30;

    let mut q = vec![valveMap.get(current_pos).unwrap()];

    while let Some(next) = q.pop() {
        for path in &next.tunnels {
            let potential_next_valve = valveMap.get(path).unwrap();
            let potential_extra_flow = potential_next_valve.flow_rate * (time_left - 2);
            // work backwards?
        }
    }

    println!("{:?}", valves);
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
                    tunnels: vec!["DD", "II", "BB"]
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
                    tunnels: vec!["DD"]
                }
            ),
            parse_valve("Valve AA has flow rate=100; tunnel leads to valve DD").unwrap()
        );
    }
}
