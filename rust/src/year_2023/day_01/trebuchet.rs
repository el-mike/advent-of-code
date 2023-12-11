use std::error::Error;
use std::io::BufRead;

use crate::common::file_utils::get_file_reader;
use crate::common::number_utils::get_number_from_word_digit;
use crate::common::string_utils::get_occurrences;

pub fn run(test_run: bool) -> Result<(), Box<dyn Error>> {
    let reader = get_file_reader("2023", "01", test_run)
        .unwrap_or_else(|err| { panic!("{}", err) });

    let mut sum: u32 = 0;

    for (i, line_result) in (0_i32..).zip(reader.lines()) {
        let line = line_result.expect("Error reading line");

        let occurrences = get_occurrences(
            line.as_str(),
            vec!["1", "2", "3", "4", "5", "6", "7", "8", "9", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"],
        );

        let first = occurrences.first().expect("First value not found");
        let last = occurrences.last().expect("Last value not found");

        let x: u32 = if first.len() == 1 { first.parse::<u32>().unwrap() } else { get_number_from_word_digit(first.as_str()) as u32 };
        let y: u32 = if last.len() == 1 { last.parse::<u32>().unwrap() } else { get_number_from_word_digit(last.as_str()) as u32 };

        println!("{}, {}", x, y);

        sum += (x * 10) + y;
    }

    println!("{}", sum);

    Ok(())
}