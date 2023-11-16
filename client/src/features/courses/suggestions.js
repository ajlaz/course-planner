import { createSlice } from "@reduxjs/toolkit";

export const suggestionSlice = createSlice({
    name: "suggestions",
    initialState: {
        courses: [],
    },
    reducers: {
        setSuggestions: (state, action) => {
            state.courses = [...action.payload];
            state.courses = state.courses.filter((course) => course.courseCode !== "");
        },
    }
});

export const { setSuggestions } = suggestionSlice.actions;