use crate::year_2023::day_03::coord::Coord;

pub struct PartNumber {
    pub coord: Coord,
    pub length: u8,
    pub value: u32,
}

impl PartNumber {
    pub fn new(coord: Coord, length: u8, value: u32) -> PartNumber {
        PartNumber { coord, length, value }
    }

    pub fn get_all_coords(&self) -> Vec<Coord> {
        let mut coords: Vec<Coord> = Vec::new();

        for x in self.coord.x..(self.coord.x + self.length as i32) {
            coords.push(Coord::new(x, self.coord.y));
        }

        return coords;
    }
}