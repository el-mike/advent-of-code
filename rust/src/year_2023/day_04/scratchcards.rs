use std::error::Error;
use std::io::BufRead;
use regex::Regex;
use crate::common::file_utils::get_file_reader;
use crate::year_2023::day_04::row::Row;

pub fn run(test_run: bool) -> Result<(), Box<dyn Error>> {
    let (reader, lines_count) = get_file_reader("2023", "04", test_run)
        .unwrap_or_else(|err| { panic!("{}", err) });

    let mut sum: u32 = 0;
    let mut rows: Vec<Row> = Vec::new();

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

        rows.push(Row::new(y as u32, winning_numbers, owned_numbers));
    }

    let mut current_rows: Vec<&Row> = rows.iter().map(|row| row).collect();
    let mut copied_rows: Vec<&Row> = Vec::new();

    loop {
        sum += current_rows.len() as u32;

        for row in current_rows.clone() {
            let i = row.card_number;
            let matches = row.get_number_of_matches();

            for j in 1..=matches {
                copied_rows.push(&rows[(i + j) as usize]);
            }
        }

        if copied_rows.is_empty() {
            break;
        }

        current_rows.clear();
        current_rows.append(&mut copied_rows.clone());

        copied_rows.clear();
    }

    println!("{}", sum);

    Ok(())
}