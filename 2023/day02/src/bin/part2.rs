use nom::{
    bytes::complete::tag,
    character::complete::{alpha1, i32},
    multi::separated_list0,
    sequence::tuple,
    IResult,
};

fn main() {
    let input = include_str!("./input.txt").trim();
    println!("solution: {}", process(input))
}

#[derive(Debug, PartialEq)]
struct Game {
    pub id: i32,
    pub draws: Vec<Draw>,
}

#[derive(Debug, PartialEq)]
struct Draw {
    pub blue: i32,
    pub red: i32,
    pub green: i32,
}

fn parse_draw(input: &str) -> IResult<&str, Draw> {
    let (i, info) = separated_list0(
        tag(", "),
        tuple((i32::<&str, nom::error::Error<&str>>, tag(" "), alpha1)),
    )(input)
    .unwrap();
    let mut blue = 0;
    let mut red = 0;
    let mut green = 0;

    for (count, _, colour) in info {
        match colour {
            "blue" => blue = count,
            "red" => red = count,
            "green" => green = count,
            _ => (),
        }
    }
    Ok((i, Draw { blue, red, green }))
}

fn parse_line(input: &str) -> IResult<&str, Game> {
    let (i, (_game, id, _colon, draws)) = tuple((
        tag("Game "),
        i32,
        tag(": "),
        separated_list0(tag("; "), parse_draw),
    ))(input)
    .unwrap();
    Ok((i, Game { id, draws }))
}

fn parse_games(input: &str) -> IResult<&str, Vec<Game>> {
    separated_list0(tag("\n"), parse_line)(input)
}

fn process(input: &str) -> i32 {
    let mut res = 0;
    dbg!(input);
    let (_, games) = parse_games(input).unwrap();

    for game in games {
        let mut known_red = 0;
        let mut known_blue = 0;
        let mut known_green = 0;
        for draw in game.draws {
            if draw.red > known_red {
                known_red = draw.red
            }
            if draw.green > known_green {
                known_green = draw.green
            }
            if draw.blue > known_blue {
                known_blue = draw.blue
            }
        }
        res += known_blue * known_red * known_green
    }
    res
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_parse_draw() {
        assert_eq!(
            parse_draw("4 blue, 5 red, 3 green"),
            Ok((
                "",
                Draw {
                    red: 5,
                    blue: 4,
                    green: 3
                }
            ))
        )
    }

    #[test]
    fn test_parse_line() {
        assert_eq!(
            parse_line("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"),
            Ok((
                "",
                Game {
                    id: 1,
                    draws: vec![
                        Draw {
                            blue: 3,
                            red: 4,
                            green: 0
                        },
                        Draw {
                            red: 1,
                            green: 2,
                            blue: 6
                        },
                        Draw {
                            green: 2,
                            red: 0,
                            blue: 0
                        }
                    ]
                }
            ))
        )
    }

    #[test]
    fn test_process() {
        assert_eq!(
            process(
                "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"
            ),
            2286
        );
    }
}
