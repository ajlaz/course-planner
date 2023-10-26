import { createSlice } from "@reduxjs/toolkit";

export const coursesSlice = createSlice({
    name: "courses",
    initialState: {
        courses: [],
    },
    reducers: {
        setCourses: (state, action) => {
            state.courses = action.payload;
        },
        addCourses : (state, action) => {
            state.courses = [...state.courses, action.payload];
        }
    },
});

export const { setCourses, addCourses } = coursesSlice.actions;