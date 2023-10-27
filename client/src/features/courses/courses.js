import { createSlice } from "@reduxjs/toolkit";

export const coursesSlice = createSlice({
    name: "courses",
    initialState: {
        courses: [],
    },
    reducers: {
        setCourses: (state, action) => {
            console.log(action)
            state.courses = [...action.payload.courses];
        },
        addCourses : (state, action) => {
            console.log(action.payload.course)
            state.courses = [...state.courses, action.payload.course];
        }
    },
});

export const { setCourses, addCourses } = coursesSlice.actions;