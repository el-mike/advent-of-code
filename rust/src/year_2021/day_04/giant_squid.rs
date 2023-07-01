use std::collections::HashSet;
use std::error::Error;
use std::io::BufRead;
use crate::common::file_utils::get_file_reader;
use crate::common::number_utils::{split_and_parse, split_by_whitespace_and_parse};

pub fn run(test_run: bool) -> Result<(), Box<dyn Error>> {
    let reader = get_file_reader("04", test_run)
        .unwrap_or_else(|err| { panic!("{}", err) });

    let mut lines = reader.lines();

    let numbers = split_and_parse(
        lines
            .next()
            .unwrap_or_else(|| panic!("Couldn't advance iterator"))
            .unwrap_or_else(|err| panic!("{}", err)).as_str(),
        ',');

    // We can use HashSet, as even when value exists more than once in given row/column,
    // remove operation would pick all of them, therefore we can reduce the data complexity
    // to hash set.
    let mut boards: Vec<Vec<HashSet<i32>>> = Vec::new();

    loop {
        let rows = lines
            .by_ref()
            // We need to skip 1 to advance over empty lines between boards.
            .skip(1)
            .take(5);

        let board = rows
            .map(|row| {
                split_by_whitespace_and_parse(
                    row.unwrap_or_else(|err| panic!("{}", err)).as_str()
                )
            })
            .collect::<Vec<Vec<i32>>>();

        if board.is_empty() {
            break;
        }

        let mut sets: Vec<HashSet<i32>> = Vec::new();
        let mut cols: Vec<HashSet<i32>> = Vec::from([
            HashSet::new(),
            HashSet::new(),
            HashSet::new(),
            HashSet::new(),
            HashSet::new()
        ]);

        for row in board {
            for x in 0..5 {
                cols[x].insert(row[x]);
            }

            sets.push(HashSet::from_iter(row));
        }

        sets.append(&mut cols);

        boards.push(sets);
    }


    'numbers: for number in &numbers {
        let mut ids_to_remove: Vec<usize> = Vec::new();
        let num_boards = boards.len();

        'boards: for (i, sets) in boards.iter_mut().enumerate() {
            for set in sets.iter_mut() {
                set.remove(number);

                if set.is_empty() {
                    if num_boards == 1 {
                        let score = boards[0]
                            // Since sets contains both rows and columns,
                            // We only  take the first half (columns) to calculate
                            // the remaining sum.
                            .clone()[0..5]
                            .iter()
                            .flatten()
                            .sum::<i32>();

                        println!("{}", score * *number);
                        break 'numbers;
                    } else {
                        ids_to_remove.push(i);
                        continue 'boards;
                    }
                }
            }
        }

        let mut i = 0;
        boards.retain(|_| {
            i += 1;
            !ids_to_remove.contains(&(i - 1))
        });
    }

    Ok(())
}
