use std::error::Error;
use std::io::BufRead;
use crate::common::file_utils::get_file_reader;

const FORWARD: &str = "forward";
const UP: &str = "up";
const DOWN: &str = "down";

pub fn run(test_run: bool) -> Result<(), Box<dyn Error>> {
    let reader = get_file_reader("2021", "02", test_run)
        .unwrap_or_else(|err| {panic!("{}", err)});

    let mut position = 0;
    let mut depth = 0;
    let mut aim = 0;

    for line_result in reader.lines() {
        let line = line_result.unwrap_or_else(|err| { panic!("{}", err) });

        let parts = line.split_whitespace().collect::<Vec<&str>>();

        let operation = parts[0];
        let value = parts[1]
            .parse::<i32>()
            .unwrap_or_else(|err| { panic!("{}", err) });

        match operation {
            FORWARD => {
                position += value;
                depth += aim * value;
            },
            UP => aim -= value,
            DOWN => aim += value,
            _ => println!("skipping")
        }
    }

    println!("{}", position * depth);

    Ok(())
}
