pub fn binary_string_to_decimal(binary_str: &str) -> i32 {
    let mut result: i32 = 0;
    let mut coefficient: i32 = 1;

    for c in binary_str.chars().rev() {
        let digit = c.to_digit(10)
            .unwrap_or_else(|| { panic!("cannot convert char to digit") });

        result += (digit as i32) * coefficient;
        coefficient *= 2;
    }

    result
}

pub fn split_and_parse(numbers_str: &str, split_by: char) -> Vec<i32> {
    numbers_str
        .split(split_by)
        .map(|x| {
            x.parse::<i32>().unwrap_or_else(|err| { panic!("{}", err) })
        })
        .collect()
}

pub fn split_by_whitespace_and_parse(numbers_str: &str) -> Vec<i32> {
    numbers_str
        .split_whitespace()
        .map(|x| {
            x.parse().unwrap_or_else(|err| { panic!("{}", err) })
        })
        .collect()
}
