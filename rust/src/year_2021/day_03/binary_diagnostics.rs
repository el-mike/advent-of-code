use std::error::Error;
use std::io::BufRead;
use crate::common::file_utils::get_file_reader;
use crate::common::number_utils::binary_string_to_decimal;

pub fn run(test_run: bool) -> Result<(), Box<dyn Error>> {
    let reader = get_file_reader("03", test_run)
        .unwrap_or_else(|err| { panic!("{}", err) });

    let mut position_counters: Vec<[i32; 2]> = vec![];

    for line_result in reader.lines() {
        let line = line_result.unwrap_or_else(|err| { panic!("{}", err) });

        for (i, c) in line.chars().enumerate() {
            if position_counters.get(i).is_none() {
                position_counters.push([0, 0]);
            }

            if c == '0' { position_counters[i][0] += 1 } else { position_counters[i][1] += 1 }
        }
    }

    // If we want to concatenate strings, we need to use String::new() -
    // using just "" (&str, borrowed type) would not allow us to allocate more memory
    // for growing string.
    let mut gamma_bin = String::new();
    let mut epsilon_bin = String::new();

    for counter in position_counters {
        if counter[0] > counter[1] {
            gamma_bin.push('0');
            epsilon_bin.push('1');
        } else {
            gamma_bin.push('1');
            epsilon_bin.push('0');
        }
    }

    let gamma = binary_string_to_decimal(gamma_bin.as_str());
    let epsilon = binary_string_to_decimal(epsilon_bin.as_str());

    println!("{}", gamma * epsilon);

    Ok(())
}
