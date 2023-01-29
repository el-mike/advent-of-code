use std::error::Error;
use std::io::BufRead;
use crate::common::file_utils::get_file_reader;
use crate::common::number_utils::binary_string_to_decimal;

#[derive(PartialEq, Eq)]
enum BitCriteria {
    MostCommon,
    LeastCommon,
}

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
            oxygen_candidates = filter_candidates(&oxygen_candidates, i, BitCriteria::MostCommon);
        }

        if scrubber_candidates.len() > 1 {
            scrubber_candidates = filter_candidates(&scrubber_candidates, i, BitCriteria::LeastCommon);
        }

        i += 1;
    }

    let oxygen = binary_string_to_decimal(oxygen_candidates[0].as_str());
    let scrubber = binary_string_to_decimal(scrubber_candidates[0].as_str());

    println!("{}", oxygen * scrubber);

    Ok(())
}

fn filter_candidates(
    candidates: &Vec<String>,
    current_position: usize,
    criteria: BitCriteria,
) -> Vec<String> {
    let mut counter: [i32; 2] = [0, 0];

    for candidate in candidates {
        if candidate
            .chars()
            .nth(current_position)
            .unwrap_or_else(|| { panic!("Couldn't read char") }) == '0' {
            counter[0] += 1;
        } else {
            counter[1] += 1;
        }
    }

    let bit_value: char;

    if criteria == BitCriteria::MostCommon {
        bit_value = if counter[0] > counter[1] { '0' } else { '1' };
    } else {
        bit_value = if counter[0] <= counter[1] { '0' } else { '1' };
    }

    candidates
        .iter()
        .cloned()
        .filter(|item| {
            item
                .chars()
                .nth(current_position)
                .unwrap_or_else(|| { panic!("Couldn't read char") }) == bit_value
        })
        .collect()
}
