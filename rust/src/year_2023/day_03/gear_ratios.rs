use std::collections::HashMap;
use std::error::Error;
use std::io::BufRead;
use crate::common::file_utils::get_file_reader;
use crate::year_2023::day_03::coord::Coord;
use crate::year_2023::day_03::part_number::PartNumber;

pub fn run(test_run: bool) -> Result<(), Box<dyn Error>> {
    let reader = get_file_reader("2023", "03", test_run)
        .unwrap_or_else(|err| { panic!("{}", err) });

    let mut numbers: Vec<PartNumber> = Vec::new();
    let mut symbol_coords_map: HashMap<i32, Vec<i32>> = HashMap::new();

    let mut max_y: i32 = 0;

    for (y, line_result) in (0_i32..).zip(reader.lines()) {
        let line = line_result.expect("Error reading line");

        let mut number_string = String::new();
        let mut length: u8 = 0;
        let mut x_start: i32 = -1;

        symbol_coords_map.insert(y, Vec::new());

        for (x, char) in line.chars().enumerate() {
            let mut break_number = false;

            if char.is_digit(10) {
                if x_start == -1 {
                    x_start = x as i32;
                }

                number_string.push(char);
                length += 1;

                // Needed for cases when the line ends with a number.
                break_number = x == line.len() - 1;
            } else {
                break_number = true;

                if char != '.' {
                    symbol_coords_map.get_mut(&(y)).unwrap().push(x as i32);
                }
            }

            if break_number && !number_string.is_empty() {
                let value = number_string.parse::<u32>().expect("Cannot parse number string");

                numbers.push(PartNumber::new(Coord::new(x_start, y), length, value));

                number_string = String::new();
                x_start = -1;
                length = 0;
            }
        }

        max_y = y;
    }

    let mut sum: u32 = 0;

    for number in numbers {
        let y = number.coord.y;
        let x_start = number.coord.x;
        let x_end = (number.coord.x + number.length as i32) - 1;

        if (y > 0 && check_row(&symbol_coords_map, y - 1, x_start, x_end))
            || (y < max_y && check_row(&symbol_coords_map, y + 1, x_start, x_end))
            || check_row(&symbol_coords_map, y, x_start, x_end) {
            sum += number.value;
        }
    }

    println!("{}", sum);

    Ok(())
}

fn check_row(symbol_coords_map: &HashMap<i32, Vec<i32>>, y: i32, x_start: i32, x_end: i32) -> bool {
    let symbol_coords = symbol_coords_map.get(&y).unwrap();

    for symbol_x in symbol_coords {
        if *symbol_x >= (x_start - 1) && *symbol_x <= (x_end + 1) {
            return true;
        }
    }

    false
}