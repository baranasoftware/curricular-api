# Design and Implementation of a REST API for Curricular data in Higher Education

The design and implementation of a REST API for student and course data for a Higher Education institution. Also includes how to implement a
data-pipeline for mostly static data. This shows how to start with users' needs (user stories) and use that to design
the API specification, and finally the implementation. This API design focus on some of the key resources such as
`/students`, `/teachers`, `/classes`
and can be extended to include other resources such as `/universities` (if the institution consists of multiple
universities such as [Universities of Wisconsin](https://www.wisconsin.edu/), that
has [13 universities](https://www.wisconsin.edu/campuses/) or [University of California](https://www.universityofcalifornia.edu/) 
that consist of [10 campuses](https://www.universityofcalifornia.edu/campuses-locations)), `/enrollments`, `/schools`, `/courses` etc..

Feel free to reach us at contact@baranasoftware.com to see how we can collaborate in your API design and implementation
effort.

## Tech Stack 

Technologies used in the implementation. `Go` language was used considering its support for building low cost and maintainable
Cloud Native apps in AWS. `SQLite` was used because this is a high transaction system (500K requests/second) 
with mostly read only data (write-once a day and read-many times a day).


* ![Docker](https://img.shields.io/badge/-Docker-eee?style=flat-square&logo=Docker&logoColor=2496ED)
* ![Go](https://img.shields.io/badge/-Go-eee?style=flat-square&logo=Go&logoColor=00ADD8)
* ![SQLite](https://img.shields.io/badge/-SQLite-eee?style=flat-square&logo=SQLite&logoColor=003B57)
* ![JSON](https://img.shields.io/badge/-JSON-eee?style=flat-square&logo=JSON&logoColor=000000)
* ![Terraform](https://img.shields.io/badge/-Terraform-eee?style=flat-square&logo=Terraform&logoColor=844FBA)
* ![AWSS3](https://img.shields.io/badge/-AWS%20S3-eee?style=flat-square&logo=amazon-s3&logoColor=569A31)
* ![AWSDynamodDB](https://img.shields.io/badge/-AWS%20DynamoDB-eee?style=flat-square&logo=amazon-dynamodb&logoColor=4053D6)
* ![AWSECS](https://img.shields.io/badge/-AWS%20ECS-eee?style=flat-square&logo=amazon-ECS&logoColor=FF9900)
* ![AWSCW](https://img.shields.io/badge/-AWS%20CloudWatch-eee?style=flat-square&logo=amazon-CloudWatch&logoColor=FF4F8B)
* ![AWSAPIGW](https://img.shields.io/badge/-AWS%20APIGateway-eee?style=flat-square&logo=amazon-API-Gateway&logoColor=FF4F8B)
* <img src="https://www.vectorlogo.zone/logos/apigee/apigee-ar21.svg" align="center" alt="Apigee" width="100">


## System Design

[Design of a REST API for Curricular data](https://github.com/baranasoftware/system-design/blob/main/edu-api.md)

## Running Locally
```
go run *.go
Curricular API server listing on port: 8080
```

|||


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

### Activities
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
| getTeachers()           | View available teachers                                | Teacher, Admin User | Student     | Teacher.Viewed    | __Request Parameters:__     __Returns:__   Teacher[]            | safe   / synchronous |
| searchTeachers()        | Search Teachers by empl ID, first name and lastname    | Teacher, Admin User | Teacher     | Teachers.Searched | __Request Parameters:__ searchQuery    __Returns:__   Teacher[] | safe   / synchronous |
| getClasses()            | View Classes by class number, name                     | Student             | Class       | Classes.Searched  | __Request Parameters:__     __Returns:__   Claas[]              | safe   / synchronous |
| getTeachersForClass()   | View Classes by teachers                               | Student             | Teacher     | Teacher.Viewed    | __Request Parameters:__ classId    __Returns:__   Teacher[]     | safe   / synchronous |
| getStudentsForTeacher() | View Students for teacher                              | Teacher, Admin User | Student     | Students.Viewed   | __Request Parameters:__ teacherId    __Returns:__   Student[]   | safe   / synchronous |

#### Modeled Resources   

##### Address 
| Property Name  | Description                    |
|----------------|--------------------------------|
| addressType    | If this is the primary address |
| streetAddress1 | Street 1                       |
| streetAddress2 | Street 2                       |
| city           | City                           |
| state          | State                          |
| zipCode        | ZipCode                        |
| country        | County                         |

##### Student
| Property Name | Description                                       |
|---------------|---------------------------------------------------|
| studentId     | Unique identifier identifying the student         |
| firstName     | Student first name                                |
| lastName      | Student last name                                 |
| address[]     | Student addresses (list of `Address`)                                |
| birthDate     | Student birthdate                                 |
| ageInYears    | Student's age in years (directly consumable data) |
| residency     | Student residency status                          |
 
##### Teacher          
| Property Name | Description                                          |
|---------------|------------------------------------------------------|
| emplId        | Unique employment identifier identifying the teacher |
| firstName     | Teacher first name                                   |
| lastName      | Teacher last name                                    |
| address[]     | Teacher addresses (list of `Address`)                |
| birthDate     | Teacher birthdate                                    |
| ageInYears    | Teacher's age in years (directly consumable data)    |

##### Class
| Property Name | Description                           |
|---------------|---------------------------------------|
| classId       | Unique identifier to identify a class |
| className     | Name of the class                     |
| credit        | Credits for this class                |
| location      | Address of the class location         |
| dayAndTime    | Day and the time of the class         |

##### Course
| Property Name | Description                           |
|---------------|---------------------------------------|
| courseId      | Unique identifier to identify a class |
| courseName    | Name of the class                     |
| termCode      | Term code for this course             |
| credit        | Total credits for this course         |
| class[]       | List of courses (of  type `Class`)    |


### Curricular API Design

| Resource Path                  | Operation Name          | HTTP Method | Description                                        | Request Details | Response Details | Response Code(s) |
|--------------------------------|-------------------------|-------------|----------------------------------------------------|-----------------|------------------|------------------|
| /students                      | getStudents()           | GET         | View students                                      |                 | Students[]       | 200              |
| /students/search               | searchStudents()        | POST        | Search for students by student id, first/last name | searchQuery     | Students[]       | 200              |
| /teachers                      | getTeachers()           | GET         | View teachers                                      |                 | Teacher[]        | 200              |
| /teachers/search               | searchTeachers()        | POST        | Search for teachers by empl id, first/last name    | searchQuery     | Teacher[]        | 200              |
| /classes                       | getClasses()            | GET         | View classes                                       |                 | Class[]          | 200              |
| /classes/{classId}/teachers    | getTeachersForClass()   | GET         | View teachers for a class                          | classId         | Teacher[]        | 200              |
| /teachers/{teacherId}/students | getStudentsForTeacher() | GET         | View students for a teacher                        | teacherId       | Students[]       | 200              |

## Roadmap

- [x] Complete API design
- [x] Include tech stack
- [x] Complete system design. Include steps to build SQLite data for static data and build a Docker image with static data
- [x] Add the ability to turn locally. Include instructions to test the flow
- [ ] Add Terraform for AWS deployment
- [ ] Implement pagination
  - https://www.jsonapi.net/usage/reading/pagination.html
- [ ] Implement batch APIs
- [ ] Implement filtering(searching): document filter/search query langauge
  - Do some design around sorting and filtering through body vs query parameters 
  - https://help.smartsuite.com/en/articles/6963760-sorting-and-filtering-records-in-the-rest-api
  - https://www.jsonapi.net/usage/reading/filtering.html
  - Implement the parser for the filtering
- [ ] Implement sparse filed selection 
  - https://www.jsonapi.net/usage/reading/sparse-fieldset-selection.html
- [ ] Implement sorting
  - https://www.jsonapi.net/usage/reading/sorting.html
