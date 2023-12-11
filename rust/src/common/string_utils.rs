use regex::Regex;

fn get_group_from_sequences(sequences: Vec<&str>) -> String {
    let group = sequences.into_iter().fold(String::new(), |a, b| {
        if a.is_empty() {
            return b.to_string();
        }

        return format!("{}|{}", a, b);
    });

    return format!("({})", group);
}

pub fn get_occurrences(hay: &str, sequences: Vec<&str>) -> Vec<String> {
    let groups = get_group_from_sequences(sequences);

    let re_string = format!(r"{}", groups);
    let re = Regex::new(re_string.as_str()).unwrap();

    let captures: Vec<String> = re.find_iter(hay).map(|m| String::from(m.as_str())).collect();

    return captures;
}