use std::fmt::Error;

use nom::branch::alt;
use nom::character::complete::{line_ending, u64};
use nom::combinator::{all_consuming, map};
use nom::multi::separated_list0;
use nom::sequence::{delimited, preceded, separated_pair, terminated, tuple};
use nom::{bytes::complete::tag, multi::separated_list1};
use nom::{Finish, IResult, Parser};

fn main() {
    let (_, packets) = all_consuming(parse_all_packets)(include_str!("../input.txt"))
        .finish()
        .unwrap();

    println!("packets: {:?}", packets);
}

fn int_list_parser(input: &str) -> IResult<&str, Vec<PacketContent>> {
    delimited(
        tag("["),
        separated_list0(
            tag(","),
            map(int_parser, PacketContent::Integer)
                .or(map(int_list_parser, PacketContent::Sublist)),
        ),
        tag("]"),
    )(input)
}

fn int_parser(input: &str) -> IResult<&str, u64> {
    u64(input)
}

fn packet_parser(input: &str) -> IResult<&str, Packet> {
    map(int_list_parser, |contents| Packet { list: contents })(input)
}

fn parse_packet_pair(i: &str) -> IResult<&str, PacketPair> {
    map(
        terminated(
            separated_pair(packet_parser, line_ending, packet_parser),
            line_ending,
        ),
        |(first, second)| PacketPair {
            0: first,
            1: second,
        },
    )(i)
}
#[derive(Debug)]
struct PacketPair(Packet, Packet);

#[derive(Debug, Clone)]
enum PacketContent {
    Integer(u64),
    Sublist(Vec<PacketContent>),
}

#[derive(Debug, Clone)]
struct Packet {
    list: Vec<PacketContent>,
}

fn parse_all_packets(input: &str) -> IResult<&str, Vec<PacketPair>> {
    separated_list1(tag("\n"), parse_packet_pair)(input)
}

// fn parse_packet(input: &str) -> IResult<&str, PacketPair> {
//     println!("input: {}", input);
//     let (_, packets) = packet_parser(input).unwrap();
//     println!("{:?}", packets);
//     Ok(("", PacketPair(vec![], vec![])))
// }

// #[cfg(test)]
// mod tests {
//     use super::*;

//     #[test]
//     fn test_int_list_parser() {
//         let (_, (_, result, _)) = int_list_parser("[1,3,5,7,9]").unwrap();
//         assert_eq!(result, vec![1, 3, 5, 7, 9]);

//         let (_, (_, result, _)) = int_list_parser("[1]").unwrap();
//         assert_eq!(result, vec![1]);
//     }

//     #[test]
//     fn test_int_parser() {
//         let (_, (_, result, _)) = int_parser("4").unwrap();
//         assert_eq!(result, vec![4]);
//     }

//     #[test]
//     fn test_packet_parser() {
//         let (_, (_, result, _)) = packet_parser("4").unwrap();
//         assert_eq!(result, vec![4]);
//         let (_, (_, result, _)) = packet_parser("[4]").unwrap();
//         assert_eq!(result, vec![4]);

//         let (_, (_, result, _)) = packet_parser("[4,5,6]").unwrap();
//         assert_eq!(result, vec![4, 5, 6]);

//         let (_, (_, result, _)) = packet_parser("[[4],[5,6]]").unwrap();
//         assert_eq!(result, vec![vec![4], vec![5, 6]]);
//     }
// }
