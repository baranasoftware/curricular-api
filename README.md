The design and implementation of a REST API for student and course data for Higher Ed. Also includes how to implement a
data-pipeline for
mostly static data.

## System Design

## Running Locally

[Design of a REST API for Curricular data](https://github.com/baranasoftware/system-design/blob/main/edu-api.md)

## API Design

API design was done
using [Align-Define-Design Process](https://blog.stoplight.io/aligning-on-your-api-design-using-jobs-to-be-done).

### User stories

| Story ID | When... (Triggering Situation)        | I want to...    (Digital Capability)                   | So I can...    (Outcome)                                    |
|----------|---------------------------------------|--------------------------------------------------------|-------------------------------------------------------------|
| 1        | I want to find students               | Search students by student ID, first name and lastname | Confirm their details                                       |
| 2        | I want to find teachers               | Search teachers by emplID, first name and lastname     | Confirm their details                                       |
| 3        | I want to find students for a teacher | View number of students for a teacher                  | Determine if the class size is too big                      |
| 4        | I want to find classes                | Search for a class                                     | Confirm that's the class I need to enroll in                |
| 5        | I want to find instructor for a class | Search for the instructor for the class                | So I can set up an appointment to discuss course assignment |

### Actives

### Activity Steps

### API Resources and Profiles

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

- [ ] Complete system design
    - Include steps to build SQLite data for static data and using
- [ ] Add Terraform for AWS deployment
- [ ] Add ability turn locally 
