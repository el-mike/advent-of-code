#[derive(Debug, Clone)]
pub struct Row {
    pub card_number: u32,
    pub winning_numbers: Vec<u32>,
    pub owned_numbers: Vec<u32>,
}

impl Row {
    pub fn new(card_number: u32, winning_numbers: Vec<u32>, owned_numbers: Vec<u32>) -> Row {
        Row {
            card_number,
            winning_numbers,
            owned_numbers,
        }
    }

    pub fn get_number_of_matches(&self) -> u32 {
        let mut matches: u32 = 0;

        for number in self.winning_numbers.iter() {
            if self.owned_numbers.contains(number) {
                matches += 1;
            }
        }

        return matches;
    }
}