use std::cmp::Ordering;
use std::num::ParseIntError;
use std::str::FromStr;

use thiserror::Error;

#[derive(Debug, Clone, Eq)]
pub enum Packet {
    Int(u32),
    List(Vec<Self>),
}

impl Packet {
    pub fn as_slice(&self) -> &[Self] {
        if let Self::List(list) = self {
            list.as_slice()
        } else {
            std::slice::from_ref(self)
        }
    }

    fn parse_one(s: &str) -> Result<(Self, &str), ParsePacketError> {
        use ParsePacketError::*;
        if let Some(mut s) = s.strip_prefix('[') {
            let mut list = vec![];
            if let Some(trailing) = s.strip_prefix(']') {
                return Ok((Self::List(list), trailing));
            }
            if s.is_empty() {
                return Err(UnclosedList);
            }
            loop {
                let (value, trailing) = Self::parse_one(s)?;
                list.push(value);
                let (c, trailing) = {
                    let mut chars = trailing.chars();
                    (chars.next(), chars.as_str())
                };
                match c {
                    Some(',') => (),
                    Some(']') => return Ok((Self::List(list), trailing)),
                    Some(c) => return Err(InvalidSeparator(c)),
                    None => return Err(UnclosedList),
                }
                s = trailing;
            }
        } else {
            let terminator = s.find([',', ']']).unwrap_or(s.len());
            let (s, trailing) = s.split_at(terminator);
            Ok((Self::Int(s.parse()?), trailing))
        }
    }
}

impl Ord for Packet {
    fn cmp(&self, other: &Self) -> Ordering {
        if let (Self::Int(a), Self::Int(b)) = (self, other) {
            a.cmp(b)
        } else {
            self.as_slice().cmp(other.as_slice())
        }
    }
}

impl PartialOrd for Packet {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        Some(self.cmp(other))
    }
}

impl PartialEq for Packet {
    fn eq(&self, other: &Self) -> bool {
        self.cmp(other).is_eq()
    }
}

#[derive(Debug, Clone, Error)]
pub enum ParsePacketError {
    #[error("invalid integer value: {0}")]
    InvalidInt(#[from] ParseIntError),
    #[error("invalid list separator: {0:?}")]
    InvalidSeparator(char),
    #[error("unclosed list value")]
    UnclosedList,
    #[error("trailing data after the packet")]
    TrailingData,
}

impl FromStr for Packet {
    type Err = ParsePacketError;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let (value, trailing) = Self::parse_one(s)?;
        if trailing.is_empty() {
            Ok(value)
        } else {
            Err(Self::Err::TrailingData)
        }
    }
}

#[macro_export]
macro_rules! packet {
    ($n:literal) => {
        $crate::Packet::Int($n)
    };
    ([$($i:tt),*]) => {
        $crate::Packet::List(vec![
            $(
                $crate::packet!($i)
            ),*
        ])
    };
}

fn main() {
    let divider_a = packet!([[2]]);
    let divider_b = packet!([[6]]);
    let mut all_packets = vec![divider_a.clone(), divider_b.clone()];
    let mut right_order_pairs = 0usize;

    let mut lines = include_str!("../input.txt").lines();
    for pair_num in 1.. {
        let a = lines.next().unwrap().parse().unwrap();
        let b = lines.next().unwrap().parse().unwrap();

        if a <= b {
            right_order_pairs += pair_num;
        }
        all_packets.push(a);
        all_packets.push(b);

        match lines.next() {
            None => break,
            Some("") => continue,
            Some(other) => panic!("Expected blank line, not {other:?}"),
        }
    }

    println!("Correctly ordered pairs: {right_order_pairs}");

    all_packets.sort();
    let mut decoder_key = 1;
    for (i, packet) in all_packets.iter().enumerate() {
        if packet == &divider_a || packet == &divider_b {
            decoder_key *= i + 1;
        }
    }

    println!("Decoder key: {decoder_key}");
}
