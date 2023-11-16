import { createSlice } from "@reduxjs/toolkit";

export const coursesSlice = createSlice({
    name: "courses",
    initialState: {
        courses: [],
    },
    reducers: {
        setCourses: (state, action) => {
            state.courses = [...action.payload.courses];
            state.courses = state.courses.filter((course) => course.courseCode !== "");
        },
        addCourses : (state, action) => {
            state.courses = [...state.courses, action.payload.course];
            state.courses = state.courses.filter((course) => course.courseCode !== "");
        },
        removeCourses : (state, action) => {
            const temp = [...state.courses]
            state.courses = temp.filter((course) => course.courseCode !== action.payload.courseCode);
            state.courses = state.courses.filter((course) => course.courseCode !== "");
        }

    },
});

export const { setCourses, addCourses, removeCourses } = coursesSlice.actions;