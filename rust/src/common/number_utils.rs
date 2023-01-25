

pub fn binary_string_to_decimal(binary_string: &str) -> i32 {
    let mut result: i32 = 0;
    let mut coefficient: i32 = 1;

    for c in binary_string.chars().rev() {
        let digit = c.to_digit(10)
            .unwrap_or_else(|| { panic!("cannot convert char to digit") });

        result += (digit as i32) * coefficient;
        coefficient *= 2;
    }

    result
}
