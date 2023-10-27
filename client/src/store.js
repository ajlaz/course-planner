import {configureStore} from '@reduxjs/toolkit';
import {coursesSlice} from './features/courses/courses';
import {usersSlice} from './features/users/users';
import { appStateSlice } from './features/appState/appState';

export default configureStore({
    reducer: {
        courses: coursesSlice.reducer,
        users: usersSlice.reducer,
        appState: appStateSlice.reducer,
    },
});