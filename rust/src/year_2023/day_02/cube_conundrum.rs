use std::collections::HashMap;
use std::error::Error;
use std::io::BufRead;
use regex::Regex;

use crate::common::file_utils::get_file_reader;

const RED_CUBE: &str = "red";
const GREEN_CUBE: &str = "green";
const BLUE_CUBE: &str = "blue";

const RED_CUBE_LIMIT: u8 = 12;
const GREEN_CUBE_LIMIT: u8 = 13;
const BLUE_CUBE_LIMIT: u8 = 14;

pub fn run(test_run: bool) -> Result<(), Box<dyn Error>> {
    let reader = get_file_reader("2023", "02", test_run)
        .unwrap_or_else(|err| { panic!("{}", err) });

    let limits_map: HashMap<&str, u8> = HashMap::from([
        (RED_CUBE, RED_CUBE_LIMIT),
        (GREEN_CUBE, GREEN_CUBE_LIMIT),
        (BLUE_CUBE, BLUE_CUBE_LIMIT),
    ]);

    let mut sum: u32 = 0;

    'outer: for(_, line_result) in (0_i32..).zip(reader.lines()) {
        let line = line_result.expect("Error reading line");

        let id_re = Regex::new(r"\d+").unwrap();
        let cubes_re = Regex::new(r"(\d+\s(red|green|blue))").unwrap();

        let id: u8 = id_re.find(line.as_str()).unwrap().as_str().parse::<u8>().unwrap();
        let cubes: Vec<&str> = cubes_re.find_iter(line.as_str()).map(|m| m.as_str()).collect();

        let cube_tuples: Vec<(u8, &str)> = cubes.iter().map(|cube| {
            let parts: Vec<&str> = cube.split_whitespace().collect();
            let count: u8 = parts[0].parse::<u8>().expect("Cannot parse cube count");
            let color: &str = parts[1];

            return (count, color)
        }).collect();

        for (count, color) in cube_tuples {
            let limit = *limits_map.get(color).expect("Cannot find limit");
            if count > limit {
                continue 'outer;
            }
        }

        sum += id as u32;
    }

    println!("\n\n{}", sum);

    Ok(())
}

//204 - too low