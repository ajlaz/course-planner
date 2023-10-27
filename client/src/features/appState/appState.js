import { createSlice } from "@reduxjs/toolkit";

export const appStateSlice = createSlice({
    name: "appState",
    initialState: {
        appState: {
            login: true,
        },
    },
    reducers: {
        toggleLogin: (state) => {
           state.appState.login = !state.appState.login;
        },
    },
});

export const { toggleLogin } = appStateSlice.actions;