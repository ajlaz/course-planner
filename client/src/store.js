import {configureStore} from '@reduxjs/toolkit';
import {coursesSlice} from './features/courses/courses';
import {usersSlice} from './features/users/users';

export default configureStore({
    reducer: {
        courses: coursesSlice.reducer,
        users: usersSlice.reducer,
    },
});