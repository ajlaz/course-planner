import {configureStore} from '@reduxjs/toolkit';
import {coursesSlice} from './features/courses/courses';
import {usersSlice} from './features/users/users';
import { appStateSlice } from './features/appState/appState';
import { hubsSlice } from './features/hubs/hubs';
import { suggestionSlice } from './features/courses/suggestions';

export default configureStore({
    reducer: {
        courses: coursesSlice.reducer,
        users: usersSlice.reducer,
        appState: appStateSlice.reducer,
        hubs: hubsSlice.reducer,
        suggestions: suggestionSlice.reducer,
    },
});