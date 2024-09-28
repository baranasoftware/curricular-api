# Design and Implementation of a REST API for Curricular data in Higher Education

The design and implementation of a REST API for student and course data for Higher Ed. Also includes how to implement a
data-pipeline for mostly static data. This shows how to start with users' needs (user stories) and use that to design
the API specification, and finally the implementation. This API design focus on some of the key resources such as `/students`, `/teachers`, `/classes` 
and can be extended to include other resources such as `/universities` (if the institution consists of multiple universities such as [Universities of Wisconsin](https://www.wisconsin.edu/), that has [13 universities](https://www.wisconsin.edu/campuses/)), `/enrollments`, `/schools`, `/courses` etc..

Feel free to reach us at contact@baranasoftware.com to see how we can collaborate in your API design and implementation
effort.

## System Design

[Design of a REST API for Curricular data](https://github.com/baranasoftware/system-design/blob/main/edu-api.md)

## Running Locally

TODO

## API Design

API design was done
using [Align-Define-Design Process](https://blog.stoplight.io/aligning-on-your-api-design-using-jobs-to-be-done).

### User stories

| Story ID | When... (Triggering Situation)              | I want to...    (Digital Capability)                   | So I can...    (Outcome)                                    |
|----------|---------------------------------------------|--------------------------------------------------------|-------------------------------------------------------------|
| 1        | I want to find students                     | Search students by student ID, first name and lastname | Confirm their details and set up an appointment             |
| 2        | I want to find teachers                     | Search teachers by emplID, first name and lastname     | Confirm their details                                       |
| 3        | I want to find students for a teacher       | View number of students for a teacher                  | Determine if the class size is too big                      |
| 5        | I want to find more details about the class | Book an appointment                                    | So I can set up an appointment to discuss course assignment |
| 4        | I want to enroll in a class                 | Search for a class                                     | Confirm that's the class I need to enroll in                |

### Actives

| Digital Capability                                     | Activity                    | Participants        | Description                                            |
|--------------------------------------------------------|-----------------------------|---------------------|--------------------------------------------------------|
| Search Students by student ID, first name and lastname | Search Students             | Teacher, Admin User | Search for students by student Id, firstname, lastname |
| Search teachers by emplID, first name and lastname     | Search Teachers             | Teacher, Admin User | Search for teachers by empl Id, firstname, lastname    |
| View number of students for a teacher                  | View Students for Teacher   | Teacher, Admin User | View students for the teacher                          |
| Book an appointment                                    | View Classes                | Student             | Search classes by class number, name                   |                        |
| Book an appointment                                    | View Teachers for the Class | Student             | Search for a teacher by class                          |

### Activity Steps

| Digital Capability                                     | Activity                    | Activity Step              | Participants        | Description                                            |
|--------------------------------------------------------|-----------------------------|----------------------------|---------------------|--------------------------------------------------------|
| Search Students by student ID, first name and lastname | Search Students             | Search Students            | Teacher, Admin User | Search for students by student Id, firstname, lastname |
| Search teachers by emplID, first name and lastname     | Search Teachers             | Search Teachers            | Teacher, Admin User | Search for teachers by empl Id, firstname, lastname    |
| View number of students for a teacher                  | View Students for Teacher   | View Teachers              | Teacher, Admin User | Search for teachers by empl Id, firstname, lastname    |
| View number of students for a teacher                  | Search Students for Teacher | View Students for Teacher  | Teacher, Admin User | View students for the teacher                          |
| Book an appointment                                    | View Classes                | View Classes               | Student             | Search classes by class number, name                   |
| Book an appointment                                    | View Teachers for the Class | View teacher for the class | Student             | Search class by teacher                                |

### API Resources and Models

Provide access to students, teachers, classes, courses and appointment data

#### API Resources

| Operation Name          | Description                                            | Participants        | Resource(s) | Emitted Events    | Operation Details                                               | Traits               |
|-------------------------|--------------------------------------------------------|---------------------|-------------|-------------------|-----------------------------------------------------------------|----------------------|
| searchStudents()        | Search Students by student ID, first name and lastname | Teacher, Admin User | Student     | Students.Searched | __Request Parameters:__ searchQuery    __Returns:__   Student[] | safe   / synchronous |
| viewTeachers()          | View available teachers                                | Teacher, Admin User | Student     | Teacher.Viewed    | __Request Parameters:__ teacherId    __Returns:__   Teacher     | safe   / synchronous |
| searchTeachers()        | Search Teachers by empl ID, first name and lastname    | Teacher, Admin User | Teacher     | Teachers.Searched | __Request Parameters:__ searchQuery    __Returns:__   Teacher[] | safe   / synchronous |
| viewClasses()           | View Classes by class number, name                     | Student             | Class       | Classes.Searched  | __Request Parameters:__ searchQuery    __Returns:__   Claas[]   | safe   / synchronous |
| viewTeacherForClass()   | View Class by teacher                                  | Student             | Teacher     | Teacher.Viewed    | __Request Parameters:__ searchQuery    __Returns:__   Claas[]   |                      |
| viewStudentForTeacher() | View Students for teacher                              |                     |             |                   |                                                                 |                      |

#### Modeled Resources

##### Student

| Property Name | Description                   |
|---------------|-------------------------------|
| title         | The book title                |
| isbn          | The unique ISBN of the book   |
| authors       | List of Book Author resources |

#### Curricular API

### Curricular API Design

| Resource Path                      | Operation Name       | HTTP Method | Description                               | Request Details          | Response Details | Response Code(s) |
|------------------------------------|----------------------|-------------|-------------------------------------------|--------------------------|------------------|------------------|
| /books                             | listBooks()          | GET         | List books by category or release date    | categoryId   releaseDate | Books[]          | 200              |
| /books/search                      | searchBooks()        | POST        | Search for books by author, title         | searchQuery              | Books[]          | 200              |
| /carts/{cartId}                    | viewCart()           | GET         | View the current cart and total           | cartId                   | Cart             | 200, 404         |
| /carts/{cartId}                    | clearCart()          | DELETE      | Remove all books from the customer's cart | cartId                   | Cart             | 204, 404         |
| /carts/{cartId}/items              | addItemToCart()      | POST        | Add a book to the customer's cart         | cartId                   | Cart             | 201, 400         |
| /carts/{cartId}/items/{cartItemId} | removeItemFromCart() | DELETE      | Remove a book from the customer's cart    | cartId   cartItemId      | Cart             | 204, 404         |
| /authors                           | getAuthorDetails()   | GET         | Retrieve the details of an author         | authorId                 | BookAuthor       | 200, 404         |

## Roadmap

- [ ] Complete API design
- [ ] Complete system design
    - Include steps to build SQLite data for static data and using
- [ ] Add ability turn locally. Include instructions to test the flow
- [ ] Add Terraform for AWS deployment
