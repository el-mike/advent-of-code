use std::error::Error;
use std::io::BufRead;
use regex::Regex;

use crate::common::file_utils::get_file_reader;

const RED_CUBE: &str = "red";
const GREEN_CUBE: &str = "green";
const BLUE_CUBE: &str = "blue";

pub fn run(test_run: bool) -> Result<(), Box<dyn Error>> {
    let reader = get_file_reader("2023", "02", test_run)
        .unwrap_or_else(|err| { panic!("{}", err) });

    let mut sum: u32 = 0;

    for(_, line_result) in (0_i32..).zip(reader.lines()) {
        let line = line_result.expect("Error reading line");

        let id_re = Regex::new(r"\d+").unwrap();
        let cubes_re = Regex::new(r"(\d+\s(red|green|blue))").unwrap();

        let id: u8 = id_re.find(line.as_str()).unwrap().as_str().parse::<u8>().unwrap();
        let cubes: Vec<&str> = cubes_re.find_iter(line.as_str()).map(|m| m.as_str()).collect();

        let mut min_red: u32 = 0;
        let mut min_green: u32 = 0;
        let mut min_blue: u32 = 0;

        let cube_tuples: Vec<(u32, &str)> = cubes.iter().map(|cube| {
            let parts: Vec<&str> = cube.split_whitespace().collect();
            let count: u32 = parts[0].parse::<u32>().expect("Cannot parse cube count");
            let color: &str = parts[1];

            return (count, color)
        }).collect();

        for (count, color) in cube_tuples {
            if color == RED_CUBE {
                if count > min_red {
                    min_red = count;
                }
            } else if color == GREEN_CUBE {
                if count > min_green {
                    min_green = count;
                }
            } else if color == BLUE_CUBE {
                if count > min_blue {
                    min_blue = count;
                }
            }
        }

        sum += (min_red * min_green * min_blue) as u32;
    }

    println!("\n\n{}", sum);

    Ok(())
}

//204 - too low