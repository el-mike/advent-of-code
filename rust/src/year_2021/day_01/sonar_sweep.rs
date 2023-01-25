use std::error::Error;
use std::io::BufRead;

use crate::common::file_utils::get_file_reader;

const WINDOW_SIZE: usize = 3;

pub fn run(test_run: bool) -> Result<(), Box<dyn Error>> {
    let reader = get_file_reader("01", test_run)
        .unwrap_or_else(|err| { panic!("{}", err) });

    let mut count: i32 = 0;

    let mut measurements: [i32; 3] = [0, 0, 0];

    let mut prev_index: usize = 0;
    let mut curr_index: usize = 1;

    // Please note that using reader.lines() (a buffer iterator) will be slower
    // than reading the entire file to memory at once.
    // @TODO: Consider using other ways of reading files.
    for (i, line_result) in (0_i32..).zip(reader.lines()) {
        // Just to demonstrate, here we have two different ways to unwrap
        // a value and handle an error.
        let line = match line_result {
            Ok(line) => line,
            Err(err) => panic!("reading file failed, {}", err),
        };

        let current_value = line
            .parse::<i32>()
            .unwrap_or_else(|err| { panic!("{}", err) });

        let prev_measurement = measurements[prev_index];

        // Takes care of resetting the values.
        measurements[prev_index] = 0;

        measurements[0] += current_value;
        measurements[1] += if i > 0 { current_value } else { 0 };
        measurements[2] += if i > 1 { current_value } else { 0 };

        if i >= WINDOW_SIZE as i32 {
            if measurements[curr_index] > prev_measurement {
                count += 1;
            }

            prev_index = (prev_index + 1) % WINDOW_SIZE;
            curr_index = (curr_index + 1) % WINDOW_SIZE;
        }
    }

    println!("{}", count);

    Ok(())
}
