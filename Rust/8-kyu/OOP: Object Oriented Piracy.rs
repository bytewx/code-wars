struct Ship {
  draft: u32,
  crew: u32,
}

impl Ship {
    fn is_worth_it(&self) -> bool {
        let crew_float = self.crew as f32;
        let draft_float = self.draft as f32;
        
        if draft_float - crew_float * 1.5 > 20.0 {
            true
        } else {
            false
        }
    }
}
