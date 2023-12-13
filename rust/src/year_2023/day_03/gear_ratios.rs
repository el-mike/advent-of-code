use std::collections::HashMap;
use std::error::Error;
use std::io::BufRead;
use crate::common::file_utils::get_file_reader;
use crate::year_2023::day_03::coord::Coord;
use crate::year_2023::day_03::part_number::PartNumber;

const GEAR_CHAR: char = '*';

pub fn run(test_run: bool) -> Result<(), Box<dyn Error>> {
    let reader = get_file_reader("2023", "03", test_run)
        .unwrap_or_else(|err| { panic!("{}", err) });

    let mut gear_symbols_coords: Vec<Coord> = Vec::new();
    let mut part_numbers_map: HashMap<i32, Vec<PartNumber>> = HashMap::new();

    let mut max_y: i32 = 0;

    for (y, line_result) in (0_i32..).zip(reader.lines()) {
        let line = line_result.expect("Error reading line");

        let mut number_string = String::new();
        let mut length: u8 = 0;
        let mut x_start: i32 = -1;

        part_numbers_map.insert(y, Vec::new());

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

                if char == GEAR_CHAR  {
                    gear_symbols_coords.push(Coord::new(x as i32, y));
                }
            }

            if break_number && !number_string.is_empty() {
                let value = number_string.parse::<u32>().expect("Cannot parse number string");


                part_numbers_map.get_mut(&y).unwrap().push(PartNumber::new(Coord::new(x_start, y), length, value));

                number_string = String::new();
                x_start = -1;
                length = 0;
            }
        }

        max_y = y;
    }

    let mut sum: u32 = 0;

    for coord in gear_symbols_coords {
        let y = coord.y;

        let mut adjacent_values: Vec<u32> = Vec::new();


        if y > 0 {
            let mut values = check_row(&part_numbers_map, y - 1, &coord);

            if values.len() > 2 {
                continue;
            }

            adjacent_values.append(&mut values);
        }

        if y < max_y {
            let mut values = check_row(&part_numbers_map, y + 1, &coord);

            if values.len() > 2 {
                continue;
            }

            adjacent_values.append(&mut values);
        }

        let mut values = check_row(&part_numbers_map, y, &coord);

        if values.len() > 2 {
            continue;
        }

        adjacent_values.append(&mut values);

        if adjacent_values.len() == 2 {
            sum += adjacent_values[0] * adjacent_values[1];
        }
    }

    println!("{}", sum);

    Ok(())
}

// Returns part number values that are adjacent to the gear symbol in given row.
fn check_row(part_numbers_map: &HashMap<i32, Vec<PartNumber>>, y: i32, gear_coord: &Coord) -> Vec<u32> {
    let part_numbers = part_numbers_map.get(&y).unwrap();

    let mut values: Vec<u32> = Vec::new();

    for part_number in part_numbers {
        let x_start = part_number.coord.x;
        let x_end = (part_number.coord.x + part_number.length as i32) - 1;

        if gear_coord.x >= (x_start - 1) && gear_coord.x <= (x_end + 1) {
            values.push(part_number.value);
        }
    }

    return values;
}