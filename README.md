The design and implementation of a REST API for student and course data for Higher Ed. Also includes how to implement a data-pipeline for 
mostly static data.

## System Design

## Running Locally
[Design of a REST API for Curricular data](https://github.com/baranasoftware/system-design/blob/main/edu-api.md) 
            
## API Design
API design was done using [Align-Define-Design Process](https://blog.stoplight.io/aligning-on-your-api-design-using-jobs-to-be-done).

### User stories
| Story ID | When... (Triggering Situation)            | I want to...    (Digital Capability)                   | So I can...    (Outcome)                   |
|----------|-------------------------------------------|--------------------------------------------------------|--------------------------------------------|
| 1        | I want to find students                   | Search students by student ID, first name and lastname | Confirm their details                      |
| 2        | I want to find instructors                | Search instructors emplID, first name and lastname     | Confirm their details                      |
| 3        | I want to find students for an instructor | View a book’s details and reviews                      | Determine if the book is of interest to me |


## Roadmap
- [ ] Complete system design
  - Include steps to build SQLite data for static data and using 
- [ ] Add Terraform for AWS deployment
- [ ] Add ability turn locally 
