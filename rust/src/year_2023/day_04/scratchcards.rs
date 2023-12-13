use std::error::Error;
use std::io::BufRead;
use regex::Regex;
use crate::common::file_utils::get_file_reader;

pub fn run(test_run: bool) -> Result<(), Box<dyn Error>> {
    let reader = get_file_reader("2023", "04", test_run)
        .unwrap_or_else(|err| { panic!("{}", err) });

    let mut sum: u32 = 0;

    for (y, line_result) in (0_i32..).zip(reader.lines()) {
        let line = line_result.expect("Error reading line");

        let label_re = Regex::new(r"Card\s+\d+:").unwrap();
        let parsed_line = label_re.replace_all(line.as_str(), "");

        let parts = parsed_line.split("|").collect::<Vec<&str>>();

        let winning_numbers = parts[0].trim().split_whitespace().map(|number| {
            number.parse::<u32>().expect("Cannot parse winning number")
        }).collect::<Vec<u32>>();

        let owned_numbers = parts[1].trim().split_whitespace().map(|number| {
            number.parse::<u32>().expect("Cannot parse owned number")
        }).collect::<Vec<u32>>();

        let mut value: u32 = 0;

        for number in winning_numbers {
            if owned_numbers.contains(&number) {
                value = if value == 0 { 1 } else { value * 2 };
            }
        }

        sum += value;
    }

    println!("{}", sum);

    Ok(())
}