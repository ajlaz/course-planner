# BU Course Planner

[demo]() (Coming soon)

## About

This is a course planner for Boston University students. It allows students to plan their courses for the next semester and see what their hubs look like, and suggests courses to help them fill their hubs.

## Usage

Courses should be inputted in the form `XXX XX 123`

Ex: `CAS CS 112`, `CAS MA 225`, `CAS WR 120` are all valid inputs.

## Local run

1. Clone the repo
2. make 2 terminals
3. In the first run

```
cd client && npm i && npm run start
```

4. in the second run

```
go run ./...
```

This should work as long as you have a basic postgres db running on your computer with port 5432.

## References

[BU Course Search](https://www.bu.edu/phpbin/course-search/)
