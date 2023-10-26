// features/users/users.js
import { createSlice } from "@reduxjs/toolkit";

export const usersSlice = createSlice({
  name: "users",
  initialState: {
    user: {
      id: null,
      username: null,
    },
  },
  reducers: {
    setUsers: (state, action) => {
      state.user = { ...action.payload };
    },
  },
});

export const { setUsers } = usersSlice.actions;
