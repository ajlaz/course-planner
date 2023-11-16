import React from 'react'
import { useDispatch, useSelector } from 'react-redux';
import { suggestCoursesToUser } from '../api/courses';
import { setSuggestions } from '../features/courses/suggestions';
import { Tooltip } from '@material-ui/core';

export default function SuggestCourses() {
    const user = useSelector(state => state.users.user);
    const suggestions = useSelector(state => state.suggestions.courses);
    const dispatch = useDispatch();
    const getSuggestions = () => {
        suggestCoursesToUser(user.id).then((res) => {
            console.log(res.data.suggestions);
            dispatch(setSuggestions(res.data.suggestions));
        }).catch((err) => {
            console.log(err);
        });
    }
  return (
    <div>
        <button onClick={getSuggestions}>Course Suggestions</button>
        <div id="suggestions">
            {
                suggestions?.map((course) => (
                    <div key={course.id} className='course'>
                        <p>{course.courseCode}</p>
                    </div>
                ))
            }
            </div>
    </div>
  )
}
