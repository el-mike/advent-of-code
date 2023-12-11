use std::error::Error;
use std::io::BufRead;

use crate::common::file_utils::get_file_reader;
use crate::common::number_utils::get_number_from_word_digit;
use crate::common::string_utils::get_occurrences;

fn parse_word_digit(word: &str, is_first: bool) -> u8 {
    match word {
        "oneight" => if is_first { 1 } else { 8 },
        "twone" => if is_first { 2 } else { 1 },
        "threeight" => if is_first { 3 } else { 8 },
        "fiveight" => if is_first { 5 } else { 8 },
        "sevenine" => if is_first { 7 } else { 9 },
        "eightwo" => if is_first { 8 } else { 2 },
        "nineight" => if is_first { 9 } else { 8 },
        _ => get_number_from_word_digit(word),
    }
}

// 55148 - too low

pub fn run(test_run: bool) -> Result<(), Box<dyn Error>> {
    let reader = get_file_reader("2023", "01", test_run)
        .unwrap_or_else(|err| { panic!("{}", err) });

    let mut sum: u32 = 0;

    for (_, line_result) in (0_i32..).zip(reader.lines()) {
        let line = line_result.expect("Error reading line");

        let occurrences = get_occurrences(
            line.as_str(),
            // Rust regex does not support look-arounds, therefore we need to handle cases like "twone" manually.
            vec!["1", "2", "3", "4", "5", "6", "7", "8", "9", "oneight", "twone", "threeight", "fiveight", "sevenine", "eightwo", "nineight", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"],
        );

        let first = occurrences.first().expect("First value not found");
        let last = occurrences.last().expect("Last value not found");

        let x: u32 = if first.len() == 1 { first.parse::<u32>().unwrap() } else { parse_word_digit(first.as_str(), true) as u32 };
        let y: u32 = if last.len() == 1 { last.parse::<u32>().unwrap() } else { parse_word_digit(last.as_str(), false) as u32 };

        sum += (x * 10) + y;
    }

    println!("{}", sum);

    Ok(())
}