import { createSlice } from "@reduxjs/toolkit";

export const hubsSlice = createSlice({
    name: "hubs",
    initialState: {
       hubs:  {}
    },
    reducers: {
        setHubs: (state, action) => {
            state.hubs = {...action.payload.hubs}

        }
    },
});

export const { setHubs } = hubsSlice.actions;