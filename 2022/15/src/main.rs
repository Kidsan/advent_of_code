use std::collections::HashSet;

use nom::{
    bytes::complete::tag, character::complete::i32, combinator::all_consuming,
    multi::separated_list0, sequence::tuple, Finish, IResult,
};
#[derive(PartialEq, Debug, Copy, Clone)]
struct Sensor {
    x: i32,
    y: i32,
    detected_beacon: (i32, i32),
    steps_to_nearest_beacon: i32,
}

impl Sensor {
    fn within_range_of_point(&self, point: (i32, i32)) -> bool {
        ((self.x - point.0).abs() + (self.y - point.1).abs()).abs() <= self.steps_to_nearest_beacon
    }

    fn can_see_quadrant(&self, min: (i32, i32), max: (i32, i32)) -> bool {
        let corners = [
            (min.0, min.1),
            (min.0, max.1),
            (max.0, min.1),
            (max.0, max.1),
        ];
        let largest_dist = corners
            .iter()
            .map(|corner| (corner.0 - self.x).abs() + (corner.1 - self.y).abs())
            .max()
            .unwrap();
        largest_dist > self.steps_to_nearest_beacon
    }
}

fn main() {
    let (_, sensors): (&str, Vec<Sensor>) =
        all_consuming(parse_all_sensors)(include_str!("../input.txt"))
            .finish()
            .unwrap();

    let mut lowest_x: i32 = 0;
    let mut lowest_y: i32 = 0;
    let mut biggest_x: i32 = 0;
    let mut biggest_y: i32 = 0;

    for sensor in &sensors {
        lowest_x = lowest_x.min(sensor.x - sensor.steps_to_nearest_beacon);
        lowest_y = lowest_y.min(sensor.y - sensor.steps_to_nearest_beacon);
        biggest_x = biggest_x.max(sensor.x + sensor.steps_to_nearest_beacon);
        biggest_y = biggest_y.max(sensor.y + sensor.steps_to_nearest_beacon);
    }

    let occupied_positions: HashSet<(i32, i32)> = sensors
        .iter()
        .flat_map(|pair| [(pair.x, pair.y), pair.detected_beacon])
        .collect();

    let desired_line = 2000000;

    let max_range = sensors
        .iter()
        .map(|sensor| sensor.steps_to_nearest_beacon)
        .max()
        .unwrap();

    let start_x = lowest_x - max_range;
    let end_x = biggest_x + max_range;

    let mut points_in_range = 0;

    for x in start_x..=end_x {
        let position = (x, desired_line);

        if occupied_positions.contains(&position) {
            continue;
        }

        if sensors
            .iter()
            .any(|sensor| sensor.within_range_of_point(position))
        {
            points_in_range += 1;
        }
    }

    println!("part one: {points_in_range}");

    const START: (i32, i32) = (0, 0);
    const END: (i32, i32) = (4000000, 4000000);

    let position = find_non_covered_point(sensors.clone(), START, END);

    println!(
        "part two: {}",
        position.0 as i64 * 4000000 + position.1 as i64
    )
}

fn find_non_covered_point(sensors: Vec<Sensor>, start: (i32, i32), end: (i32, i32)) -> (i32, i32) {
    let mut q = vec![(start, end)];

    while let Some((start, end)) = q.pop() {
        if start == end {
            if sensors
                .iter()
                .all(|sensor| !sensor.within_range_of_point(start))
            {
                return start;
            }
        }

        let middle = ((start.0 + end.0) / 2, (start.1 + end.1) / 2);

        let quardants = [
            (start, middle),
            ((middle.0 + 1, start.1), (end.0, middle.1)),
            ((start.0, middle.1 + 1), (middle.0, end.1)),
            ((middle.0 + 1, middle.1 + 1), end),
        ];

        for quadrant in quardants.iter() {
            if quadrant.0 .0 > quadrant.1 .0 || quadrant.0 .1 > quadrant.1 .1 {
                continue;
            }

            if sensors
                .iter()
                .all(|sensor| sensor.can_see_quadrant(quadrant.0, quadrant.1))
            {
                q.push(*quadrant);
            }
        }
    }

    (1, 1)
}

fn parse_sensor(input: &str) -> IResult<&str, Sensor> {
    let (i, (_, x, _, y, _, nx, _, ny)): (&str, (&str, i32, &str, i32, &str, i32, &str, i32)) =
        tuple((
            tag("Sensor at x="),
            i32::<&str, nom::error::Error<&str>>,
            tag(", y="),
            i32,
            tag(": closest beacon is at x="),
            i32,
            tag(", y="),
            i32,
        ))(input)
        .unwrap();

    Ok((
        i,
        Sensor {
            x,
            y,
            steps_to_nearest_beacon: (x - nx).abs() + (y - ny).abs(),
            detected_beacon: (nx, ny),
        },
    ))
}

fn parse_all_sensors(i: &str) -> IResult<&str, Vec<Sensor>> {
    separated_list0(tag("\n"), parse_sensor)(i)
}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn test_parse_sensor() {
        assert_eq!(
            (
                "",
                Sensor {
                    x: 2,
                    y: 18,
                    detected_beacon: (-2, 15),
                    steps_to_nearest_beacon: 7
                }
            ),
            parse_sensor("Sensor at x=2, y=18: closest beacon is at x=-2, y=15").unwrap()
        );
        assert_eq!(
            (
                "",
                Sensor {
                    x: 8,
                    y: 7,
                    detected_beacon: (2, 10),
                    steps_to_nearest_beacon: 9
                }
            ),
            parse_sensor("Sensor at x=8, y=7: closest beacon is at x=2, y=10").unwrap()
        );
    }
}
