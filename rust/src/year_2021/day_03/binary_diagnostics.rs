use std::error::Error;
use std::io::BufRead;
use crate::common::file_utils::get_file_reader;
use crate::common::number_utils::binary_string_to_decimal;

pub fn run(test_run: bool) -> Result<(), Box<dyn Error>> {
    let reader = get_file_reader("03", test_run)
        .unwrap_or_else(|err| { panic!("{}", err) });

    // Two different ways of creating Vectors.
    let mut oxygen_candidates: Vec<String> = Vec::new();
    let mut scrubber_candidates: Vec<String> = vec![];

    for line_result in reader.lines() {
        let line = line_result.unwrap_or_else(|err| { panic!("{}", err) });

        oxygen_candidates.push(line.clone());
        scrubber_candidates.push(line.clone());
    }

    let mut i = 0;

    while oxygen_candidates.len() > 1 || scrubber_candidates.len() > 1 {
        if oxygen_candidates.len() > 1 {
            let mut counter: [i32; 2] = [0, 0];

            // As slice is very important here!
            // Without it, this for loop will be the owner of the String values inside
            // the Vector, moving the value from outer scope.
            for candidate in oxygen_candidates.as_slice() {
                if candidate.chars().nth(i).unwrap() == '0'
                    { counter[0] += 1 } else { counter[1] += 1 }
            }

            let bit_value = if counter[0] > counter[1] { '0' } else { '1' };

            // Instead of using .into_iter().filter().collect(), we can use much simpler
            // retain() method. Please note that retain works in-place.
            oxygen_candidates.retain(
                |item| { item.chars().nth(i).unwrap() == bit_value }
            );
        }

        if scrubber_candidates.len() > 1 {
            let mut counter: [i32; 2] = [0, 0];

            for candidate in scrubber_candidates.as_slice() {
                if candidate.chars().nth(i).unwrap() == '0'
                { counter[0] += 1 } else { counter[1] += 1 }
            }

            let bit_value = if counter[0] <= counter[1] { '0' } else { '1' };

            scrubber_candidates.retain(
                |item| { item.chars().nth(i).unwrap() == bit_value }
            );
        }

        i += 1;
    }

    let oxygen = binary_string_to_decimal(oxygen_candidates[0].as_str());
    let scrubber = binary_string_to_decimal(scrubber_candidates[0].as_str());

    println!("{}", oxygen * scrubber);

    Ok(())
}
